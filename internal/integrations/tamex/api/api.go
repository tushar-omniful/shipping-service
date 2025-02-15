package api

import (
	"context"
	"encoding/json"
	"github.com/omniful/go_commons/config"
	logsUtils "github.com/omniful/shipping-service/utils/logs"
	"net/url"
	"sync"
	"time"

	commonError "github.com/omniful/go_commons/error"
	"github.com/omniful/go_commons/httpclient"
	"github.com/omniful/go_commons/httpclient/request"
	"github.com/omniful/go_commons/log"

	"github.com/omniful/shipping-service/constants"
	"github.com/omniful/shipping-service/internal/domain/interfaces"
	"github.com/omniful/shipping-service/internal/domain/models"

	respAdapter "github.com/omniful/shipping-service/internal/integrations/tamex/adapter/responses"
	"github.com/omniful/shipping-service/internal/services/shipping_commons/api"
	apirequest "github.com/omniful/shipping-service/internal/services/shipping_commons/api/request"
	shsError "github.com/omniful/shipping-service/pkg/error"
)

var client *Client
var once sync.Once

func NewClient(ctx context.Context, psmRepo interfaces.PartnerShippingMethodRepository) *Client {
	once.Do(func() {
		c := api.NewShippingClient(config.GetString(ctx, "tamex.host"), psmRepo,
			httpclient.WithContentType(constants.ApplicationJSON),
			httpclient.WithDeadline(time.Second*30),
		)
		client = &Client{c}
	})
	return client
}

func GetClient() *Client {
	return client
}

type Client struct {
	client *api.ShippingClient
}

func (c *Client) Create(ctx context.Context, request map[string]interface{}, psm models.PartnerShippingMethod) (result respAdapter.CreateResponse, cusErr commonError.CustomError) {
	logTag := logsUtils.GetLogTag(ctx, "CreateShipment")

	req, err := c.prepareRequestBuilder().
		SetHeaders(url.Values{
			constants.ContentType: []string{constants.ApplicationJSON},
			constants.Accept:      []string{constants.ApplicationJSON}}).
		SetUri("/api/v2/create").
		SetBody(request).
		Build()

	if err != nil {
		log.Errorf("%s - Failed to build request: %v", logTag, err)
		cusErr = raiseInternalError(err)
		return
	}

	resp, err := c.client.Post(ctx, c.buildShippingRequest(req, psm))
	if err != nil {
		log.Errorf("%s - Failed to send request: %v", logTag, err)
		cusErr = raiseInternalError(err)
		return
	}

	//Tamex is sending text/html content type
	err = json.Unmarshal(resp.Body(), &result)
	if err != nil {
		log.Errorf("%s - Failed to unmarshal response: %v", logTag, err)
		cusErr = raiseInternalError(err)
		return
	}

	return
}

func (c *Client) Cancel(ctx context.Context, request map[string]interface{}, psm models.PartnerShippingMethod) (result respAdapter.CancelResponse, cusErr commonError.CustomError) {
	logTag := logsUtils.GetLogTag(ctx, "CancelShipment")

	req, err := c.prepareRequestBuilder().
		SetHeaders(url.Values{constants.ContentType: []string{constants.ApplicationJSON}}).
		SetUri("/api/v2/cancel").
		SetBody(request).
		Build()

	if err != nil {
		log.Errorf("%s - Failed to build request: %v", logTag, err)
		cusErr = raiseInternalError(err)
		return
	}

	resp, err := c.client.Post(ctx, c.buildShippingRequest(req, psm))
	if err != nil {
		log.Errorf("%s - Failed to send request: %v", logTag, err)
		cusErr = raiseInternalError(err)
		return
	}

	if err = resp.UnmarshalBody(&result); err != nil {
		log.Errorf("%s - Failed to unmarshal response: %v", logTag, err)
		cusErr = raiseInternalError(err)
		return
	}

	return
}

func (c *Client) Track(ctx context.Context, request map[string]interface{}, psm models.PartnerShippingMethod) (result respAdapter.TrackResponse, cusErr commonError.CustomError) {
	logTag := logsUtils.GetLogTag(ctx, "TrackShipment")

	req, err := c.prepareRequestBuilder().
		SetHeaders(url.Values{constants.ContentType: []string{constants.ApplicationJSON}}).
		SetUri("/api/v2/status").
		SetBody(request).
		Build()

	if err != nil {
		log.Errorf("%s - Failed to build request: %v", logTag, err)
		cusErr = raiseInternalError(err)
		return
	}

	resp, err := c.client.Post(ctx, c.buildShippingRequest(req, psm))
	if err != nil {
		log.Errorf("%s - Failed to send request: %v", logTag, err)
		cusErr = raiseInternalError(err)
		return
	}

	if err = resp.UnmarshalBody(&result); err != nil {
		log.Errorf("%s - Failed to unmarshal response: %v", logTag, err)
		cusErr = raiseInternalError(err)
		return
	}

	return
}

func (c *Client) buildShippingRequest(req request.Request, psm models.PartnerShippingMethod) apirequest.ShippingRequest {
	return apirequest.NewShippingRequest(req, psm)
}

func (c *Client) prepareRequestBuilder() request.Builder {
	return request.NewBuilder()
}

func raiseInternalError(err error) commonError.CustomError {
	return commonError.NewCustomError(shsError.ShippingServiceInternalError, err.Error())
}
