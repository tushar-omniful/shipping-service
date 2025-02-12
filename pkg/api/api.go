package api

import (
	"context"
	"fmt"
	"github.com/omniful/go_commons/env"
	commons "github.com/omniful/go_commons/http"
	"github.com/omniful/go_commons/log"
	"net/http"
	"time"
)

// RESTClient handles REST API calls
type RESTClient struct {
	client  *commons.Client
	baseURL string
}

type Config struct {
	BaseURL   string
	Timeout   time.Duration
	Transport *http.Transport
}

// NewRESTClient creates a new REST client instance
func NewRESTClient(serviceName string, config Config) (*RESTClient, error) {
	if config.Timeout == 0 {
		config.Timeout = 60 * time.Second
	}

	config.Transport = &http.Transport{
		MaxIdleConns:        100,
		MaxIdleConnsPerHost: 100,
		IdleConnTimeout:     config.Timeout,
		DisableCompression:  true,
	}

	client, err := commons.NewHTTPClient(
		serviceName,
		config.BaseURL,
		config.Transport,
		commons.WithTimeout(config.Timeout),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP client: %w", err)
	}

	return &RESTClient{
		client:  client,
		baseURL: config.BaseURL,
	}, nil
}

type Request struct {
	Method      string
	Path        string
	Headers     map[string][]string
	Body        interface{}
	QueryParams map[string][]string
}

type Response struct {
	StatusCode int
	Body       interface{}
	IsSuccess  bool
}

func (c *RESTClient) Send(ctx context.Context, req *Request) (response *Response, err error) {
	logTag := fmt.Sprintf("RequestID: %s", env.GetRequestID(ctx))

	httpReq := &commons.Request{
		Url:         c.baseURL + req.Path,
		Headers:     req.Headers,
		QueryParams: req.QueryParams,
		Body:        req.Body,
	}

	method := commons.APIMethod(req.Method)
	resp, err := c.client.Execute(ctx, method, httpReq, nil)

	if err != nil {
		log.Errorf("%s API request failed: %v", logTag, err)
		return
	}

	log.Infof("%s API headers: %v url: %s body: %v response: %s",
		logTag, httpReq.Headers, httpReq.Url, httpReq.Body, resp.Body())

	response = &Response{
		IsSuccess:  resp.IsSuccess(),
		Body:       resp.Body(),
		StatusCode: resp.StatusCode(),
	}

	return
}
