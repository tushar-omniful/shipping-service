package shipment

import (
	"context"
	"github.com/omniful/shipping-service/internal/services/shipment/integrations/ajex/adapter/requests"
	"github.com/omniful/shipping-service/internal/services/shipment/integrations/ajex/adapter/responses"
	"github.com/omniful/shipping-service/pkg/api"
	"net/http"
	"time"
)

type Config struct {
	BaseURL   string
	Timeout   time.Duration
	Transport *http.Transport
}

// BaseService provides common functionality for all shipping services
type BaseService struct {
	RestClient  *api.RESTClient
	config      Config
	ReqAdapter  *requests.RequestAdapter
	RespAdapter *responses.ResponseAdapter
}

func NewBaseService(config Config, reqAdapter *requests.RequestAdapter, respAdapter *responses.ResponseAdapter) (*BaseService, error) {
	restClient, err := api.NewRESTClient("shipping-service", api.Config{
		BaseURL:   config.BaseURL,
		Timeout:   config.Timeout,
		Transport: config.Transport,
	})
	if err != nil {
		return nil, err
	}

	return &BaseService{
		config:      config,
		RestClient:  restClient,
		ReqAdapter:  reqAdapter,
		RespAdapter: respAdapter,
	}, nil
}

func (s *BaseService) SendRESTRequest(ctx context.Context, method, path string, headers map[string][]string, body interface{}, queryParams map[string][]string) (*api.Response, error) {
	req := &api.Request{
		Method:      method,
		Path:        path,
		Headers:     headers,
		Body:        body,
		QueryParams: queryParams,
	}
	return s.RestClient.Send(ctx, req)
}
