package api

import (
	"context"
	"fmt"
	"github.com/omniful/go_commons/config"
	"net/url"
	"sync"
	"time"

	logsUtils "github.com/omniful/shipping-service/utils/logs"

	commonError "github.com/omniful/go_commons/error"
	"github.com/omniful/go_commons/httpclient"
	"github.com/omniful/go_commons/httpclient/request"
	"github.com/omniful/go_commons/log"

	"github.com/omniful/shipping-service/constants"
	"github.com/omniful/shipping-service/internal/domain/interfaces"
	"github.com/omniful/shipping-service/internal/domain/models"

	respAdapter "github.com/omniful/shipping-service/internal/integrations/ajex/adapter/responses"
	"github.com/omniful/shipping-service/internal/services/shipping_commons/api"
	apirequest "github.com/omniful/shipping-service/internal/services/shipping_commons/api/request"
	shsError "github.com/omniful/shipping-service/pkg/error"
)

var client *Client
var once sync.Once

func NewClient(ctx context.Context, psmRepo interfaces.PartnerShippingMethodRepository) *Client {
	once.Do(func() {
		c := api.NewShippingClient(config.GetString(ctx, "ajex.host"), psmRepo,
			httpclient.WithContentType(constants.ApplicationJSON),
			httpclient.WithDeadline(time.Second*30),
			//httpclient.WithAfterAttemptCallback(GetAfterAttemptCallback()),
		)
		client = &Client{c, psmRepo}
	})
	return client
}

func GetClient() *Client {
	return client
}

type Client struct {
	client *api.ShippingClient

	psmRepo interfaces.PartnerShippingMethodRepository
}

func (c *Client) Create(ctx context.Context, request map[string]interface{}, psm models.PartnerShippingMethod) (result respAdapter.CreateResponse, cusErr commonError.CustomError) {
	logTag := logsUtils.GetLogTag(ctx, "CreateOrder")

	req, err := c.prepareRequestBuilder().
		SetHeaders(url.Values{constants.ContentType: []string{constants.ApplicationJSON}}).
		SetUri("/order-management/api/v2/order").
		SetBody(request).
		Build()

	if err != nil {
		log.Errorf("%s - Failed to build request: %v", logTag, err)
		cusErr = commonError.NewCustomError(shsError.ShippingServiceInternalError, err.Error())
		return
	}

	opts := c.prepareRequestOptions(psm)

	resp, err := c.client.Post(ctx, c.buildShippingRequest(req, psm), opts...)
	fmt.Println(resp.Body(), string(resp.Body()), resp.StatusCode(), resp.IsSuccess())

	if err != nil {
		log.Errorf("%s - Failed to send request: %v", logTag, err)
		cusErr = commonError.NewCustomError(shsError.ShippingServiceInternalError, err.Error())
		return
	}

	if err = resp.UnmarshalBody(&result); err != nil {
		log.Errorf("%s - Failed to unmarshal fetch products response: %v", logTag, err)
		cusErr = commonError.NewCustomError(shsError.ShippingServiceInternalError, err.Error())
		return
	}

	return
}

func (c *Client) generateAccessToken(ctx context.Context, psm models.PartnerShippingMethod) (loginResponse respAdapter.LoginResponse, cusErr commonError.CustomError) {
	logTag := logsUtils.GetLogTag(ctx, "GenerateAccessToken")

	fd := url.Values{}
	fd.Set(constants.Username, psm.Credentials["username"].(string))
	//fd.Set(constants.Password, psm.Credentials["password"].(string))
	fd.Set(constants.Password, "CQFalB6LJBXqsCtBhL5U")
	req, err := request.NewBuilder().
		SetUri("/authentication-service/api/auth/login").
		SetBody(fd).
		Build()

	if err != nil {
		log.Errorf("%s - Failed to build request: %v", logTag, err)
		cusErr = commonError.NewCustomError(shsError.ShippingServiceInternalError, err.Error())
		return
	}

	resp, err := c.client.Post(ctx, c.buildShippingRequest(req, psm))
	if err != nil {
		log.Errorf("%s - Failed to send request: %v", logTag, err)
		cusErr = commonError.NewCustomError(shsError.ShippingServiceInternalError, err.Error())
		return
	}

	//if resp.IsCode2XX() && len(resp.Body()) < 3 {
	//	log.Errorf("%s - Received unexpected response from Ajex: insufficient data", logTag)
	//	cusErr = commonError.NewCustomError(shsError.ShippingServiceInternalError, "WrongCredentials Ajex token")
	//	return
	//}

	if err = resp.UnmarshalBody(&loginResponse); err != nil {
		log.Errorf("%s - Failed to unmarshal Ajex token response: %v", logTag, err)
		cusErr = commonError.NewCustomError(shsError.ShippingServiceInternalError, err.Error())
		return
	}

	return
}

func (c *Client) buildShippingRequest(req request.Request, psm models.PartnerShippingMethod) apirequest.ShippingRequest {
	shReq := apirequest.NewShippingRequest(req, psm)
	shReq.IsUnauthorizedFunc = isUnauthorizedFunc
	shReq.AuthRefresherFunc = buildAuthRefresherFunc(c.psmRepo, psm)
	shReq.AuthProviderFactory = buildAuthProviderFactory(psm)
	return shReq
}

func (c *Client) prepareRequestBuilder() request.Builder {
	return request.NewBuilder()
}

func (c *Client) prepareRequestOptions(psm models.PartnerShippingMethod) httpclient.Options {
	return httpclient.Options{
		httpclient.WithRequestAuthProvider(NewAuthProvider(psm)),
	}
}
