package endpoints

import (
	"context"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/sony/gobreaker"
	"golang.org/x/time/rate"

	"github.com/go-kit/kit/circuitbreaker"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/ratelimit"

	"github.com/michaljirman/newsapp/newsfeeder-service/pkg/models"
	"github.com/michaljirman/newsapp/newsfeeder-service/pkg/services"
)

// Sets collects all of the endpoints that compose an feed service. It's meant to
// be used as a helper struct, to collect all of the endpoints into a single
// parameter.
type Sets struct {
	CreateFeedEndpoint  endpoint.Endpoint
	GetFeedsEndpoint    endpoint.Endpoint
	GetArticlesEndpoint endpoint.Endpoint
	GetArticleEndpoint  endpoint.Endpoint
}

// CreateEndpoints returns a Set that wraps the provided server, and wires in all of the
// expected endpoint middlewares via the various parameters.
func CreateEndpoints(svc services.FeedService, logger *logrus.Logger) Sets {
	var createFeedEndpoint endpoint.Endpoint
	{
		createFeedEndpoint = MakeCreateFeedEndpoint(svc)
		createFeedEndpoint = ratelimit.NewErroringLimiter(rate.NewLimiter(rate.Every(time.Second), 1))(createFeedEndpoint)
		createFeedEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{}))(createFeedEndpoint)
	}

	var getFeedsEndpoint endpoint.Endpoint
	{
		getFeedsEndpoint = MakeGetFeedsEndpoint(svc)
		getFeedsEndpoint = ratelimit.NewErroringLimiter(rate.NewLimiter(rate.Every(time.Second), 1))(getFeedsEndpoint)
		getFeedsEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{}))(getFeedsEndpoint)
	}

	var getArticlesEndpoint endpoint.Endpoint
	{
		getArticlesEndpoint = MakeGetArticlesEndpoint(svc)
		getArticlesEndpoint = ratelimit.NewErroringLimiter(rate.NewLimiter(rate.Every(time.Second), 1))(getArticlesEndpoint)
		getArticlesEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{}))(getArticlesEndpoint)
	}

	var getArticleEndpoint endpoint.Endpoint
	{
		getArticleEndpoint = MakeGetArticleEndpoint(svc)
		getArticleEndpoint = ratelimit.NewErroringLimiter(rate.NewLimiter(rate.Every(time.Second), 1))(getArticleEndpoint)
		getArticleEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{}))(getArticleEndpoint)
	}

	return Sets{
		CreateFeedEndpoint:  createFeedEndpoint,
		GetFeedsEndpoint:    getFeedsEndpoint,
		GetArticlesEndpoint: getArticlesEndpoint,
		GetArticleEndpoint:  getArticleEndpoint,
	}
}

// MakeCreateFeedEndpoint constructs a CreateFeed endpoint wrapping the service.
func MakeCreateFeedEndpoint(svc services.FeedService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateFeedRequest)
		feedID, err := svc.CreateFeed(ctx, req.Category, req.Provider, req.URL)
		return CreateFeedResponse{FeedID: feedID, ResourceURI: fmt.Sprintf("%s/%d", req.RequestURI, feedID), Err: err}, nil
	}
}

// MakeGetFeedsEndpoint constructs a GetFeeds endpoint wrapping the service.
func MakeGetFeedsEndpoint(svc services.FeedService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(GetFeedsRequest)
		feeds, err := svc.GetFeeds(ctx, req.Category, req.Provider)
		return GetFeedsResponse{Feeds: feeds, Err: err}, nil
	}
}

// MakeGetArticlesEndpoint constructs a GetArticles endpoint wrapping the service.
func MakeGetArticlesEndpoint(svc services.FeedService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(GetArticlesRequest)
		articles, err := svc.GetArticles(ctx, req.FeedID)
		return GetArticlesResponse{Articles: articles, Err: err}, nil
	}
}

// MakeGetArticleEndpoint constructs a GetArticle endpoint wrapping the service.
func MakeGetArticleEndpoint(svc services.FeedService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(GetArticleRequest)
		article, err := svc.GetArticle(ctx, req.FeedID, req.ArticleGUID)
		return GetArticleResponse{Article: article, Err: err}, nil
	}
}

// CreateFeedRequest collects the request parameters for the CreateFeed method.
type CreateFeedRequest struct {
	Category   string `json:"category"`
	Provider   string `json:"provider"`
	URL        string `json:"url"`
	RequestURI string
}

// CreateFeedResponse collects the response parameters for the CreateFeed method.
type CreateFeedResponse struct {
	FeedID      uint64
	ResourceURI string `json:"-"`
	Err         error  `json:"err,omitempty"`
}

// Failed implements endpoint.Failer.
func (resp CreateFeedResponse) Failed() error { return resp.Err }

// GetArticlesRequest collects the request parameters for the GetArticles method.
type GetArticlesRequest struct {
	FeedID uint64
}

// GetArticlesResponse collects the response parameters for the GetArticles method.
type GetArticlesResponse struct {
	Articles []models.Article `json:"articles"`
	Err      error            `json:"err,omitempty"`
}

// Failed implements endpoint.Failer.
func (r GetArticlesResponse) Failed() error { return r.Err }

// GetFeedsRequest collects the request parameters for the GetFeeds method.
type GetFeedsRequest struct {
	Category string `json:"category"`
	Provider string `json:"provider"`
}

// GetFeedsResponse collects the response parameters for the GetFeeds method.
type GetFeedsResponse struct {
	Feeds []models.Feed `json:"feeds"`
	Err   error         `json:"err,omitempty"`
}

// Failed implements endpoint.Failer.
func (r GetFeedsResponse) Failed() error { return r.Err }

// GetArticleRequest collects the request parameters for the GetArticle method.
type GetArticleRequest struct {
	FeedID      uint64 `json:"feed_id"`
	ArticleGUID string `json:"article_guid"`
}

// GetArticleResponse collects the response parameters for the GetArticle method.
type GetArticleResponse struct {
	Article models.Article `json:"article"`
	Err     error          `json:"err,omitempty"`
}

// Failed implements endpoint.Failer.
func (r GetArticleResponse) Failed() error { return r.Err }
