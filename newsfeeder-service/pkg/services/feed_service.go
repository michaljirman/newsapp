package services

import (
	"context"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"sync"

	"github.com/mmcdole/gofeed"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"github.com/michaljirman/newsapp/newsfeeder-service/pkg/configs"
	"github.com/michaljirman/newsapp/newsfeeder-service/pkg/models"
	"github.com/michaljirman/newsapp/newsfeeder-service/pkg/repositories"
)

// FeedService describes a service that manages feeds and articles.
type FeedService interface {
	CreateFeed(ctx context.Context, category, provider, url string) (uint64, error)
	GetFeeds(ctx context.Context, category, provider string) ([]models.Feed, error)
	GetArticles(ctx context.Context, feedID uint64) ([]models.Article, error)
	GetArticle(ctx context.Context, feedID uint64, articleGUID string) (models.Article, error)
}

// NewFeedService returns a basic FeedService with all necessary dependencies such as feedParser, repository, logger and config.
func NewFeedService(logger *logrus.Logger, cfg *configs.Config, repo repositories.FeedRepository, feedParser *gofeed.Parser) FeedService {
	var svc FeedService
	{
		svc = newBasicFeedService(logger, cfg, repo, feedParser)
	}
	return svc
}

type basicFeedService struct {
	logger     *logrus.Logger
	cfg        *configs.Config
	repo       repositories.FeedRepository
	feedParser *gofeed.Parser
}

// CreateFeed creates a new feed.
func (s *basicFeedService) CreateFeed(ctx context.Context, category, provider, url string) (uint64, error) {
	feed := models.Feed{Category: category, Provider: provider, URL: url}
	feedID, err := s.repo.CreateFeed(ctx, feed)
	if err != nil {
		return 0, errors.Wrap(err, "failed to create a news feed")
	}
	return feedID, nil
}

// GetFeeds retrieves feeds by category and provider.
func (s *basicFeedService) GetFeeds(ctx context.Context, category, provider string) ([]models.Feed, error) {
	feeds, err := s.repo.GetFeeds(ctx, category, provider)
	if err != nil {
		return nil, errors.Wrap(err, "failed get feeds")
	}
	return feeds, nil
}

// GetArticle retrieves a single article by feed identifier and article guid.
func (s *basicFeedService) GetArticle(ctx context.Context, feedID uint64, articleGUID string) (models.Article, error) {
	feedInfoDB, err := s.repo.GetFeedByID(ctx, feedID)
	if err != nil {
		return models.Article{}, errors.Wrapf(err, "failed to retrieve feed by ID %d", feedID)
	}

	feed, err := s.feedParser.ParseURL(feedInfoDB.URL)
	if err != nil {
		s.logger.Error(errors.Wrap(err, "failed to parse feedID as url"))
	}

	for _, item := range feed.Items {
		if item.GUID == articleGUID {
			newsArticle := models.Article{
				Title:       item.Title,
				Description: item.Description,
				Link:        item.Link,
				GUID:        item.GUID,
			}

			if item.PublishedParsed != nil {
				newsArticle.Published = *item.PublishedParsed
			}

			htmlContent, err := fetchHTMLlContent(item)
			if err != nil {
				s.logger.Warn("failed to fetch html content")
			}
			newsArticle.HTMLContent = base64.StdEncoding.EncodeToString(htmlContent)

			thumbnailFeedImgURL := prepareThumbnailFeedImageURL(item, s.cfg.URLBox)
			if len(thumbnailFeedImgURL) != 0 {
				newsArticle.ThumbnailImageURL = thumbnailFeedImgURL
			}
			return newsArticle, nil
		}
	}
	return models.Article{}, errors.Errorf("failed to fetch article %s", articleGUID)
}

// GetArticles retrieves all articles by feed ID.
func (s *basicFeedService) GetArticles(ctx context.Context, feedID uint64) ([]models.Article, error) {
	var newsArticles []models.Article
	feedInfoDB, err := s.repo.GetFeedByID(ctx, feedID)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to retrieve feed by ID %d", feedID)
	}

	feed, err := s.feedParser.ParseURL(feedInfoDB.URL)
	if err != nil {
		s.logger.Error(errors.Wrap(err, "failed to parse feedID as url"))
	}

	var lock sync.Mutex
	var wg sync.WaitGroup
	for _, item := range feed.Items {
		wg.Add(1)
		go func(item *gofeed.Item) {
			defer wg.Done()

			newsArticle := models.Article{
				Title:       item.Title,
				Description: item.Description,
				Link:        item.Link,
				GUID:        item.GUID,
			}

			if item.PublishedParsed != nil {
				newsArticle.Published = *item.PublishedParsed
			}

			thumbnailFeedImgURL := prepareThumbnailFeedImageURL(item, s.cfg.URLBox)
			if len(thumbnailFeedImgURL) != 0 {
				newsArticle.ThumbnailImageURL = thumbnailFeedImgURL
			}
			lock.Lock()
			newsArticles = append(newsArticles, newsArticle)
			lock.Unlock()

		}(item)
	}
	wg.Wait()

	sort.Slice(newsArticles, func(i, j int) bool {
		return newsArticles[i].Published.Before(newsArticles[j].Published)
	})

	return newsArticles, nil
}

// prepareThumbnailFeedImageURL extracts thumbnail image url from gofeed.Item
func prepareThumbnailFeedImageURL(item *gofeed.Item, cfg configs.URLBoxConfig) string {
	// select the image field if it is set
	if item.Image != nil {
		return item.Image.URL
	}
	return fmt.Sprintf("%s%s/png?url=%s&thumb_width=%d&ttl=%d", cfg.URL, cfg.Token, url.QueryEscape(item.Link), 150, 86400)
}

// fetchThumbnailImageData fetches image data from external provider
// TODO use 3rd party API (api.urlbox.io) to fetch thumbnail data
// https://api.urlbox.io/v1/TFb7yVm5aCzXBXhD/png?url=apple.com&thumb_width=150&ttl=86400
// resp, err := http.Get(fmt.Sprintf("%s/v1/%s/png?url=%s&thumb_width=%d?ttl=%d", s.urlBox.URL, link, 150, 86400))
func fetchThumbnailImageData(item *gofeed.Item, url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer resp.Body.Close()
	thumbnail, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return thumbnail, nil
}

// fetchHTMLContent fetches escaped html string
func fetchHTMLlContent(item *gofeed.Item) ([]byte, error) {
	resp, err := http.Get(item.Link)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer resp.Body.Close()
	pageHTML, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return pageHTML, nil
}

func newBasicFeedService(logger *logrus.Logger, cfg *configs.Config, repo repositories.FeedRepository, feedParser *gofeed.Parser) FeedService {
	return &basicFeedService{
		logger:     logger.WithField("service", "feedservice").Logger,
		cfg:        cfg,
		repo:       repo,
		feedParser: feedParser,
	}
}
