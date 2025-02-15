package api

import (
	"context"
	"fmt"
	"github.com/omniful/go_commons/httpclient"
	"github.com/omniful/go_commons/httpclient/response"
	"github.com/omniful/shipping-service/internal/domain/interfaces"
	"github.com/omniful/shipping-service/internal/services/shipping_commons/api/request"
	"github.com/omniful/shipping-service/internal/services/shipping_commons/api/retries"
	"github.com/omniful/shipping-service/pkg/api"
	"net/http"
)

func NewShippingClient(host string, psmRepo interfaces.PartnerShippingMethodRepository, opts ...httpclient.Option) *ShippingClient {
	c := api.NewClient(host, opts...)
	return &ShippingClient{Client: c, psmRepo: psmRepo}
}

type ShippingClient struct {
	httpclient.Client
	psmRepo interfaces.PartnerShippingMethodRepository
}

func (sc *ShippingClient) Get(ctx context.Context, req request.ShippingRequest, opts ...httpclient.Option) (response.Response, error) {
	req, err := sc.setHttpMethod(req, http.MethodGet)
	if err != nil {
		return nil, err
	}
	return sc.Send(ctx, req, opts...)
}

func (sc *ShippingClient) Post(ctx context.Context, req request.ShippingRequest, opts ...httpclient.Option) (response.Response, error) {
	req, err := sc.setHttpMethod(req, http.MethodPost)
	if err != nil {
		return nil, err
	}
	return sc.Send(ctx, req, opts...)
}

func (sc *ShippingClient) Put(ctx context.Context, req request.ShippingRequest, opts ...httpclient.Option) (response.Response, error) {
	req, err := sc.setHttpMethod(req, http.MethodPut)
	if err != nil {
		return nil, err
	}
	return sc.Send(ctx, req, opts...)
}

func (sc *ShippingClient) Patch(ctx context.Context, req request.ShippingRequest, opts ...httpclient.Option) (response.Response, error) {
	req, err := sc.setHttpMethod(req, http.MethodPatch)
	if err != nil {
		return nil, err
	}
	return sc.Send(ctx, req, opts...)
}

func (sc *ShippingClient) Delete(ctx context.Context, req request.ShippingRequest, opts ...httpclient.Option) (response.Response, error) {
	req, err := sc.setHttpMethod(req, http.MethodDelete)
	if err != nil {
		return nil, err
	}
	return sc.Send(ctx, req, opts...)
}

func (sc *ShippingClient) Send(ctx context.Context, req request.ShippingRequest, opts ...httpclient.Option) (response.Response, error) {
	// Prepend default options
	opts = append(sc.defaultRequestOptions(req), opts...)
	return sc.Client.Send(ctx, req.HttpRequest, opts...)
}

func (sc *ShippingClient) defaultRequestOptions(req request.ShippingRequest) httpclient.Options {
	return httpclient.Options{
		httpclient.WithRetry(retries.NewAuthRetry(req.PsmID, req.GetIsUnauthorizedFunc(), req.GetAuthRefresherFunc(), req.GetAuthProviderFactory(), sc.psmRepo)),
		httpclient.WithRetry(retries.NewDefaultRetry(req.RetryStrategy)),
		//httpclient.WithAfterSendCallback(callbacks.GetAfterSendCallback(req)),
		//httpclient.WithOnErrorCallback(callbacks.GetOnErrorCallback(req)),
	}
}

func (sc *ShippingClient) setHttpMethod(req request.ShippingRequest, method string) (request.ShippingRequest, error) {
	newHttpReq, err := req.HttpRequest.ToBuilder().SetMethod(method).Build()
	if err != nil {
		return req, err
	}
	req.HttpRequest = newHttpReq
	return req, nil
}

func (sc *ShippingClient) PrepareUrl(domain, uri string) string {
	return fmt.Sprintf("https://%s%s", domain, uri)
}
