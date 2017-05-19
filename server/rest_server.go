package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	//"log"
	"bytes"
	"flag"
	"io/ioutil"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/gorilla/mux"
	"net/http"
	//"net/url"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
)

var (
	ErrNotFound        = errors.New("Not found")
	ErrAlreadyExists   = errors.New("News already exists")
	ErrInconsistentIDs = errors.New("Inconsistent IDs")
)

type TweetPublisherService interface {
	PostTweets(ctx context.Context, nws News) error
}

type News struct {
	NewsID     string `json:"newsid"`
	NewsAuthor string `json:"newsauthor,omitempty"`
	NewsDate   string `json:"newsdate,omitempty"`
	NewsText   string `json:"newstext,omitempty"`
}

type tweetPublisherService struct {
	mtx sync.RWMutex
	m   map[string]News
}

func NewTweetPublisherService() TweetPublisherService {
	return &tweetPublisherService{
		m: map[string]News{},
	}
}

func (s *tweetPublisherService) PostTweets(_ context.Context, nws News) error {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	if _, ok := s.m[nws.NewsID]; ok {
		return ErrAlreadyExists
	}
	s.m[nws.NewsID] = nws
	fmt.Println(s.m[nws.NewsID])

	dbObject.populateTable(nws.NewsID, nws.NewsAuthor, nws.NewsDate, nws.NewsText)
	//dbObject.queryTable()
	//to write into db
	return nil
}

//###########################
//for client

type Endpoints struct {
	PostTweetsEndpoint endpoint.Endpoint
}

func (e Endpoints) PostTweets(ctx context.Context, nws News) error {
	request := postNewsRequest{News: nws}
	response, err := e.PostTweetsEndpoint(ctx, request)
	if err != nil {
		return err
	}
	resp := response.(postNewsResponse)
	return resp.Err
}

//func MakeClientEndpoints(nws News)

//###########################

func makeServerEndpoints(svc TweetPublisherService) Endpoints {
	return Endpoints{
		PostTweetsEndpoint: makePostTweetsEndpoint(svc),
	}
}

func makePostTweetsEndpoint(svc TweetPublisherService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(postNewsRequest)
		err := svc.PostTweets(ctx, req.News) //ctx,
		return postNewsResponse{Err: err}, nil
	}
}

type postNewsRequest struct {
	News News
}

type postNewsResponse struct {
	Err error `json:"err,omitempty"`
}

func MakeHTTPHandler(svc TweetPublisherService, logger log.Logger) http.Handler {
	router := mux.NewRouter().StrictSlash(true)
	e := makeServerEndpoints(svc)
	options := []httptransport.ServerOption{
		httptransport.ServerErrorLogger(logger),
		httptransport.ServerErrorEncoder(encodeError),
	}

	router.Methods("POST").Path("/News/").Handler(httptransport.NewServer(
		e.PostTweetsEndpoint, decodePostNewsRequest, encodeResponse, options...,
	))

	return router
}

func encodeRequest(_ context.Context, req *http.Request, request interface{}) error {
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(request)
	if err != nil {
		return err
	}
	req.Body = ioutil.NopCloser(&buf)
	return nil
}

func encodePostNewsRequest(ctx context.Context, req *http.Request, request interface{}) error {
	req.Method, req.URL.Path = "POST", "/News/"
	return encodeRequest(ctx, req, request)
}

func decodePostNewsRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req postNewsRequest
	if err := json.NewDecoder(r.Body).Decode(&req.News); err != nil {
		return nil, err
	}
	return req, nil
}

type errorer interface {
	error() error
}

func decodePostNewsResponse(_ context.Context, resp *http.Response) (interface{}, error) {
	var response postNewsResponse
	err := json.NewDecoder(resp.Body).Decode(&response)
	return response, err
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if err, ok := response.(errorer); ok && err.error() != nil {
		encodeError(ctx, err.error(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func encodeError(ctx context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		panic("encodeError with nil error")
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(codeFrom(err))
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

func codeFrom(err error) int {
	switch err {
	case ErrNotFound:
		return http.StatusNotFound
	case ErrAlreadyExists, ErrInconsistentIDs:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}

func restServerMain(dbObj *DB) {
	dbObject = dbObj
	var (
		httpAddr = flag.String("http.addr", ":8082", "HTTP listen address")
	)
	flag.Parse()

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	svc := NewTweetPublisherService()

	var h http.Handler
	{
		h = MakeHTTPHandler(svc, log.With(logger, "component", "HTTP"))
	}

	errs := make(chan error)

	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		logger.Log("transport", "HTTP", "addr", *httpAddr)
		errs <- http.ListenAndServe(*httpAddr, h)
	}()

	logger.Log("exit", <-errs)
}

/*
type apiEndpoint struct {
	PutEndpoint endpoint.Endpoint
}

func MakeServerEndpoints(s Service) apiEndpoint {

}
*/
