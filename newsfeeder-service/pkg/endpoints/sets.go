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

type Endpoints struct {
	CreateFeedEndpoint  endpoint.Endpoint
	GetFeedsEndpoint    endpoint.Endpoint
	GetArticlesEndpoint endpoint.Endpoint
}

func CreateEndpoints(svc services.FeedService, logger *logrus.Logger) Endpoints {
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

	return Endpoints{
		CreateFeedEndpoint:  createFeedEndpoint,
		GetFeedsEndpoint:    getFeedsEndpoint,
		GetArticlesEndpoint: getArticlesEndpoint,
	}
}

func MakeCreateFeedEndpoint(svc services.FeedService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateFeedRequest)
		feedID, err := svc.CreateFeed(ctx, req.Category, req.Provider, req.URL)
		return CreateFeedResponse{FeedID: feedID, ResourceURI: fmt.Sprintf("%s/%d", req.RequestURI, feedID), Err: err}, nil
	}
}

func MakeGetFeedsEndpoint(svc services.FeedService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(GetFeedsRequest)
		feeds, err := svc.GetFeeds(ctx, req.Category, req.Provider)
		return GetFeedsResponse{Feeds: feeds, Err: err}, nil
	}
}

func MakeGetArticlesEndpoint(svc services.FeedService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(GetArticlesRequest)
		articles, err := svc.GetArticles(ctx, req.FeedID)
		return GetArticlesResponse{Articles: articles, Err: err}, nil
	}
}

type CreateFeedRequest struct {
	Category   string `json:"category"`
	Provider   string `json:"provider"`
	URL        string `json:"url"`
	RequestURI string
}

type CreateFeedResponse struct {
	FeedID      uint64
	ResourceURI string `json:"-"`
	Err         error  `json:"err,omitempty"`
}

func (resp CreateFeedResponse) Failed() error { return resp.Err }

type GetArticlesRequest struct {
	FeedID uint64
}

type GetArticlesResponse struct {
	Articles []models.Article `json:"articles,omitempty"`
	Err      error            `json:"err,omitempty"`
}

func (r GetArticlesResponse) error() error { return r.Err }

type GetFeedsRequest struct {
	Category string
	Provider string
}

type GetFeedsResponse struct {
	Feeds []models.Feed `json:"feeds,omitempty"`
	Err   error         `json:"err,omitempty"`
}

func (r GetFeedsResponse) error() error { return r.Err }
