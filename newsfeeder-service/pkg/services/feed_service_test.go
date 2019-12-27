package services

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/joho/godotenv"
	"github.com/mmcdole/gofeed"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"

	"github.com/michaljirman/newsapp/newsfeeder-service/pkg/configs"
	"github.com/michaljirman/newsapp/newsfeeder-service/pkg/mock_repositories"
	"github.com/michaljirman/newsapp/newsfeeder-service/pkg/models"
)

func setup(t *testing.T) (*configs.Config, *logrus.Logger, context.Context) {
	godotenv.Load("../.env")

	//todo setup method
	cfg, err := configs.Get()
	if err != nil {
		t.Fatalf("failed to load config: %+v", err)
	}

	// setup global logger settings
	ll, err := logrus.ParseLevel("debug")
	if err != nil {
		t.Fatalf("invalid value for loglevel (could not be parsed): %+v", err)
	}
	logger := logrus.New()
	logger.SetReportCaller(true)
	logger.SetLevel(ll)
	return cfg, logger, context.Background()
}

func TestGetFeeds(t *testing.T) {
	cfg, logger, ctx := setup(t)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockFeedRepo := mock_repositories.NewMockFeedRepository(ctrl)

	timeNow := time.Now()
	// Expected results
	var expectedFeeds []models.Feed
	expectedFeed1 := models.Feed{
		ID:       1,
		Created:  timeNow,
		Updated:  timeNow,
		Category: "UK",
		Provider: "BBC",
		URL:      "http://feeds.bbci.co.uk/news/uk/rss.xml",
	}

	expectedFeed2 := models.Feed{
		ID:       2,
		Created:  timeNow,
		Updated:  timeNow,
		Category: "Technology",
		Provider: "BBC",
		URL:      "http://feeds.bbci.co.uk/news/technology/rss.xml",
	}

	expectedFeed3 := models.Feed{
		ID:       3,
		Created:  timeNow,
		Updated:  timeNow,
		Category: "UK",
		Provider: "Reuters",
		URL:      "http://feeds.reuters.com/reuters/UKdomesticNews?format=xml",
	}

	expectedFeed4 := models.Feed{
		ID:       4,
		Created:  timeNow,
		Updated:  timeNow,
		Category: "Technology",
		Provider: "Reuters",
		URL:      "http://feeds.reuters.com/reuters/UKdomesticNews?format=xml",
	}
	expectedFeeds = append(expectedFeeds, expectedFeed1)
	expectedFeeds = append(expectedFeeds, expectedFeed2)
	expectedFeeds = append(expectedFeeds, expectedFeed3)
	expectedFeeds = append(expectedFeeds, expectedFeed4)

	feedSvc := NewFeedService(logger, cfg, mockFeedRepo, gofeed.NewParser())

	mockFeedRepo.EXPECT().GetFeeds(ctx, gomock.Eq(""), gomock.Eq("")).Return(expectedFeeds, error(nil))
	feeds, err := feedSvc.GetFeeds(ctx, "", "")
	if err != nil {
		t.Fatal("failed to get feeds")
	}
	assert.Equal(t, 4, len(feeds))

	mockFeedRepo.EXPECT().GetFeeds(ctx, gomock.Eq("UK"), gomock.Eq("")).Return([]models.Feed{expectedFeed1, expectedFeed3}, error(nil))
	feeds, err = feedSvc.GetFeeds(ctx, "UK", "")
	if err != nil {
		t.Fatal("failed to get feeds")
	}
	assert.Equal(t, 2, len(feeds))
}

func TestGetArticles(t *testing.T) {
	cfg, logger, ctx := setup(t)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockFeedRepo := mock_repositories.NewMockFeedRepository(ctrl)

	var feedTests = []struct {
		file       string
		feedType   string
		feedTitle  string
		hasError   bool
		itemsCount int
	}{
		{"BBCNewsUKXMLFeed.xml", "rss", "BBC News - UK", false, 32},
		{"ReutersUKPoweredByFeedBurner.xml", "rss", "ReutersUKPoweredByFeedBurner", false, 20},
	}

	for _, test := range feedTests {
		fmt.Printf("Testing %s... ", test.file)

		// Get feed content
		path := fmt.Sprintf("../testdata/%s", test.file)
		f, _ := ioutil.ReadFile(path)

		// Get actual value
		server, client := mockServerResponse(200, string(f), 0)
		timeNow := time.Now()
		// Expected results
		expectedFeed := models.Feed{
			ID:       1,
			Created:  timeNow,
			Updated:  timeNow,
			Category: "UK",
			Provider: "BBC",
			URL:      server.URL,
		}

		feedParser := gofeed.NewParser()
		feedParser.Client = client
		feedSvc := NewFeedService(logger, cfg, mockFeedRepo, feedParser)

		feedID := uint64(1)
		mockFeedRepo.EXPECT().GetFeedByID(ctx, gomock.Eq(feedID)).Return(&expectedFeed, error(nil))
		articles, err := feedSvc.GetArticles(ctx, feedID)
		if err != nil {
			t.Fatal("failed to get articles")
		}
		t.Log(articles)
		assert.Equal(t, test.itemsCount, len(articles))
	}
}

// Test Helpers

func mockServerResponse(code int, body string, delay time.Duration) (*httptest.Server, *http.Client) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(code)
		w.Header().Set("Content-Type", "application/xml")
		io.WriteString(w, body)
	}))

	transport := &http.Transport{
		Proxy: func(req *http.Request) (*url.URL, error) {
			return url.Parse(server.URL)
		},
	}

	client := &http.Client{Transport: transport}
	return server, client
}
