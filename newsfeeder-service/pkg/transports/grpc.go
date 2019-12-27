package transports

import (
	"context"
	"errors"
	"time"

	"github.com/go-kit/kit/circuitbreaker"
	"github.com/go-kit/kit/endpoint"
	kitLog "github.com/go-kit/kit/log/logrus"
	"github.com/go-kit/kit/ratelimit"
	kitGRPCTransport "github.com/go-kit/kit/transport/grpc"
	"github.com/golang/protobuf/ptypes"
	"github.com/sirupsen/logrus"
	"github.com/sony/gobreaker"
	"golang.org/x/time/rate"
	"google.golang.org/grpc"

	"github.com/michaljirman/newsapp/newsfeeder-service/pb"
	"github.com/michaljirman/newsapp/newsfeeder-service/pkg/endpoints"
	"github.com/michaljirman/newsapp/newsfeeder-service/pkg/models"
)

type gRPCServer struct {
	createFeed  kitGRPCTransport.Handler
	getArticles kitGRPCTransport.Handler
	getFeeds    kitGRPCTransport.Handler
}

// NewGRPCServer creates a new instance of GRPC server
func NewGRPCServer(endpoints endpoints.Endpoints, logger *logrus.Logger) pb.FeederServer {
	options := []kitGRPCTransport.ServerOption{
		kitGRPCTransport.ServerErrorLogger(kitLog.NewLogrusLogger(logger)),
	}

	return &gRPCServer{
		createFeed: kitGRPCTransport.NewServer(
			endpoints.CreateFeedEndpoint,
			decodeGRPCCreateFeedRequest,
			encodeGRPCCreateFeedResponse,
			options...,
		),
		getFeeds: kitGRPCTransport.NewServer(
			endpoints.GetFeedsEndpoint,
			decodeGRPCGetFeedsRequest,
			encodeGRPCGetFeedsResponse,
			options...,
		),
		getArticles: kitGRPCTransport.NewServer(
			endpoints.GetArticlesEndpoint,
			decodeGRPCGetArticlesRequest,
			encodeGRPCGetArticlesResponse,
			options...,
		),
	}
}

func decodeGRPCCreateFeedRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.CreateFeedRequest)
	return endpoints.CreateFeedRequest{Category: req.Category, Provider: req.Provider, URL: req.Url, RequestURI: req.RequestUri}, nil
}

func encodeGRPCCreateFeedResponse(_ context.Context, r interface{}) (interface{}, error) {
	resp := r.(endpoints.CreateFeedResponse)
	return &pb.CreateFeedReply{FeedId: resp.FeedID, ResourceUri: resp.ResourceURI, Err: err2str(resp.Err)}, nil
}

func decodeGRPCGetArticlesRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.GetArticlesRequest)
	return endpoints.GetArticlesRequest{FeedID: req.FeedId}, nil
}

func convertToPBArticle(articles []models.Article) ([]*pb.Article, error) {
	var articlesPb []*pb.Article
	for _, article := range articles {
		articlePb := &pb.Article{
			Title:             article.Title,
			Description:       article.Description,
			Link:              article.Link,
			Guid:              article.GUID,
			ThumbnailImageUrl: article.ThumbnailImageURL,
			HtmlContent:       article.HTMLContent,
		}
		publishedPb, err := ptypes.TimestampProto(article.Published)
		if err != nil {
			return nil, err
		}
		articlePb.Published = publishedPb
		articlesPb = append(articlesPb, articlePb)
	}
	return articlesPb, nil
}

func encodeGRPCGetArticlesResponse(_ context.Context, r interface{}) (interface{}, error) {
	resp := r.(endpoints.GetArticlesResponse)
	articlesPb, err := convertToPBArticle(resp.Articles)
	if err != nil {
		return nil, err
	}
	return &pb.GetArticlesReply{Articles: articlesPb, Err: err2str(resp.Err)}, nil
}

func decodeGRPCGetFeedsRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.GetFeedsRequest)
	return endpoints.GetFeedsRequest{Category: req.Category, Provider: req.Provider}, nil
}

func convertToPBFeed(feeds []models.Feed) ([]*pb.Feed, error) {
	var feedsPb []*pb.Feed
	for _, feed := range feeds {
		feedPb := &pb.Feed{
			FeedId:   feed.ID,
			Category: feed.Category,
			Provider: feed.Provider,
			Url:      feed.URL,
		}
		createdAtPb, err := ptypes.TimestampProto(feed.Created)
		if err != nil {
			return nil, err
		}
		feedPb.CreatedAt = createdAtPb

		updatedAtPb, err := ptypes.TimestampProto(feed.Updated)
		if err != nil {
			return nil, err
		}
		feedPb.UpdatedAt = updatedAtPb

		feedsPb = append(feedsPb, feedPb)
	}
	return feedsPb, nil
}

func encodeGRPCGetFeedsResponse(_ context.Context, r interface{}) (interface{}, error) {
	resp := r.(endpoints.GetFeedsResponse)
	feedsPb, err := convertToPBFeed(resp.Feeds)
	if err != nil {
		return nil, err
	}
	return &pb.GetFeedsReply{Feeds: feedsPb, Err: err2str(resp.Err)}, nil
}

func (g *gRPCServer) CreateFeed(ctx context.Context, req *pb.CreateFeedRequest) (*pb.CreateFeedReply, error) {
	_, reply, err := g.createFeed.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return reply.(*pb.CreateFeedReply), nil
}

func (g *gRPCServer) GetFeeds(ctx context.Context, req *pb.GetFeedsRequest) (*pb.GetFeedsReply, error) {
	_, reply, err := g.getFeeds.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return reply.(*pb.GetFeedsReply), nil
}

func (g *gRPCServer) GetArticles(ctx context.Context, req *pb.GetArticlesRequest) (*pb.GetArticlesReply, error) {
	_, reply, err := g.getArticles.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return reply.(*pb.GetArticlesReply), nil
}

// Helper functions are required to translate Go error types to
// and from strings, which is the type we use in our IDLs to represent errors.
// There is special casing to treat empty strings as nil errors.
func str2err(s string) error {
	if s == "" {
		return nil
	}
	return errors.New(s)
}

func err2str(err error) string {
	if err == nil {
		return ""
	}
	return err.Error()
}

func NewGRPCClient(conn *grpc.ClientConn, logger *logrus.Logger) endpoints.Endpoints {
	limiter := ratelimit.NewErroringLimiter(rate.NewLimiter(rate.Every(time.Second), 100))

	var options []kitGRPCTransport.ClientOption

	var getArticlesEndpoint endpoint.Endpoint
	{
		getArticlesEndpoint = kitGRPCTransport.NewClient(
			conn,
			"pb.Feeder",
			"GetArticles",
			encodeGRPCGetArticlesRequest,
			decodeGRPCGetArticlesResponse,
			pb.GetArticlesReply{},
			options...,
		).Endpoint()
		getArticlesEndpoint = limiter(getArticlesEndpoint)
		getArticlesEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
			Name:    "getArticles",
			Timeout: 30 * time.Second,
		}))(getArticlesEndpoint)
	}

	var getFeedsEndpoint endpoint.Endpoint
	{
		getFeedsEndpoint = kitGRPCTransport.NewClient(
			conn,
			"pb.Feeder",
			"GetFeeds",
			encodeGRPCGetFeedsRequest,
			decodeGRPCGetFeedsResponse,
			pb.GetFeedsReply{},
			options...,
		).Endpoint()
		getFeedsEndpoint = limiter(getFeedsEndpoint)
		getFeedsEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
			Name:    "getFeeds",
			Timeout: 30 * time.Second,
		}))(getFeedsEndpoint)
	}

	var createFeedEndpoint endpoint.Endpoint
	{
		createFeedEndpoint = kitGRPCTransport.NewClient(
			conn,
			"pb.Feeder",
			"CreateFeed",
			encodeGRPCCreateFeedRequest,
			decodeGRPCCreateFeedResponse,
			pb.CreateFeedReply{},
			options...,
		).Endpoint()
		createFeedEndpoint = limiter(createFeedEndpoint)
		createFeedEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
			Name:    "createFeed",
			Timeout: 30 * time.Second,
		}))(createFeedEndpoint)
	}

	return endpoints.Endpoints{
		CreateFeedEndpoint:  createFeedEndpoint,
		GetFeedsEndpoint:    getFeedsEndpoint,
		GetArticlesEndpoint: getArticlesEndpoint,
	}
}

func encodeGRPCGetArticlesRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoints.GetArticlesRequest)
	return &pb.GetArticlesRequest{FeedId: req.FeedID}, nil
}

func decodeGRPCGetArticlesResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*pb.GetArticlesReply)
	var articles []models.Article
	for _, articlePb := range reply.Articles {
		article := models.Article{
			Title:             articlePb.Title,
			Description:       articlePb.Description,
			Link:              articlePb.Link,
			GUID:              articlePb.Guid,
			ThumbnailImageURL: articlePb.ThumbnailImageUrl,
		}
		published, err := ptypes.Timestamp(articlePb.Published)
		if err != nil {
			return nil, err
		}
		article.Published = published
	}
	return endpoints.GetArticlesResponse{Articles: articles, Err: str2err(reply.Err)}, nil
}

func encodeGRPCGetFeedsRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoints.GetFeedsRequest)
	return &pb.GetFeedsRequest{Category: req.Category, Provider: req.Provider}, nil
}

func convertFromPBFeed(feedPb []*pb.Feed) ([]models.Feed, error) {
	var feeds []models.Feed
	for _, feedPb := range feedPb {
		feed := models.Feed{
			ID:       feedPb.FeedId,
			Category: feedPb.Category,
			Provider: feedPb.Provider,
			URL:      feedPb.Url,
		}
		created, err := ptypes.Timestamp(feedPb.CreatedAt)
		if err != nil {
			return nil, err
		}
		feed.Created = created

		updated, err := ptypes.Timestamp(feedPb.UpdatedAt)
		if err != nil {
			return nil, err
		}
		feed.Updated = updated
	}
	return feeds, nil
}

func decodeGRPCGetFeedsResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*pb.GetFeedsReply)
	feeds, err := convertFromPBFeed(reply.Feeds)
	if err != nil {
		return nil, err
	}
	return endpoints.GetFeedsResponse{Feeds: feeds, Err: str2err(reply.Err)}, nil
}

func encodeGRPCCreateFeedRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoints.CreateFeedRequest)
	return &pb.CreateFeedRequest{Category: req.Category, Provider: req.Provider, Url: req.URL, RequestUri: req.RequestURI}, nil
}

func decodeGRPCCreateFeedResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*pb.CreateFeedReply)
	return endpoints.CreateFeedResponse{FeedID: reply.FeedId, ResourceURI: reply.GetResourceUri(), Err: str2err(reply.Err)}, nil
}
