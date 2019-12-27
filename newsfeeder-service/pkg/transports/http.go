package transports

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"

	kitEndpoint "github.com/go-kit/kit/endpoint"
	kitLog "github.com/go-kit/kit/log/logrus"
	kitHttpTransport "github.com/go-kit/kit/transport/http"

	"github.com/michaljirman/newsapp/newsfeeder-service/pkg/endpoints"
)

var (
	ErrBadRouting = errors.New("inconsistent mapping between route and handler (programmer error)")
)

func NewHTTPHandler(endpoints endpoints.Sets, logger *logrus.Logger) http.Handler {
	options := []kitHttpTransport.ServerOption{
		kitHttpTransport.ServerErrorEncoder(errorEncoder),
		kitHttpTransport.ServerErrorLogger(kitLog.NewLogrusLogger(logger)),
	}

	r := mux.NewRouter()
	r.Methods("POST").Path("/feeds").Handler(kitHttpTransport.NewServer(
		endpoints.CreateFeedEndpoint,
		decodeHTTPPostFeedRequest,
		encodeHTTPPostFeedResponse,
		options...,
	))

	r.Methods("GET").Path("/feeds").Handler(kitHttpTransport.NewServer(
		endpoints.GetFeedsEndpoint,
		decodeHTTPGetFeedsRequest,
		encodeHTTPGenericResponse,
		options...,
	))

	r.Methods("GET").Path("/feeds/{feed_id}/articles").Handler(kitHttpTransport.NewServer(
		endpoints.GetArticlesEndpoint,
		decodeHTTPGetArticlesRequest,
		encodeHTTPGenericResponse,
		options...,
	))

	r.Methods("POST").Path("/articles/searches").Handler(kitHttpTransport.NewServer(
		endpoints.GetArticleEndpoint,
		decodeHTTPGetArticleRequest,
		encodeHTTPGenericResponse,
		options...,
	))

	return r
}

func decodeHTTPPostFeedRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoints.CreateFeedRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	req.RequestURI = r.RequestURI
	return req, err
}

func decodeHTTPGetFeedsRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoints.GetFeedsRequest{}
	if categoryValues, ok := r.URL.Query()["category"]; ok {
		req.Category = categoryValues[0]
	}
	if providerValues, ok := r.URL.Query()["provider"]; ok {
		req.Provider = providerValues[0]
	}
	return req, nil
}

func decodeHTTPGetArticlesRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	feedIDParam, ok := vars["feed_id"]
	if !ok {
		return nil, ErrBadRouting
	}
	feedID, err := strconv.ParseUint(feedIDParam, 10, 64)
	if err != nil {
		return nil, ErrBadRouting
	}
	req := endpoints.GetArticlesRequest{FeedID: feedID}
	return req, nil
}

func decodeHTTPGetArticleRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoints.GetArticleRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

//// encodeHTTPGenericResponse is a transport/http.EncodeResponseFunc that encodes
//// the response as JSON to the response writer. Primarily useful in a server.
func encodeHTTPGenericResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if f, ok := response.(kitEndpoint.Failer); ok && f.Failed() != nil {
		errorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	encoder := json.NewEncoder(w)
	encoder.SetEscapeHTML(false)
	return encoder.Encode(response)
}

func encodeHTTPPostFeedResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if f, ok := response.(kitEndpoint.Failer); ok && f.Failed() != nil {
		errorEncoder(ctx, f.Failed(), w)
		return nil
	}

	if resp, ok := response.(endpoints.CreateFeedResponse); ok {
		w.Header().Set("Location", resp.ResourceURI)
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusCreated)
	return json.NewEncoder(w).Encode(response)
}

func errorEncoder(_ context.Context, err error, w http.ResponseWriter) {
	w.WriteHeader(err2code(err))
	json.NewEncoder(w).Encode(errorWrapper{Error: err.Error()})
}

func err2code(err error) int {
	return http.StatusInternalServerError
}

type errorWrapper struct {
	Error string `json:"error"`
}
