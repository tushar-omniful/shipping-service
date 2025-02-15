package request

import (
	"github.com/omniful/go_commons/httpclient/request"
	"github.com/omniful/shipping-service/internal/domain/models"
	"github.com/omniful/shipping-service/pkg/api"

	//"github.com/omniful/shipping-service/internal/services/shipping_commons/api"
	"github.com/omniful/shipping-service/internal/services/shipping_commons/api/auth"
)

func NewShippingRequest(req request.Request, psm models.PartnerShippingMethod) ShippingRequest {
	return ShippingRequest{
		HttpRequest: req,

		OpID:  psm.OrderPartnerID,
		PsmID: psm.ID,
	}
}

type ShippingRequest struct {
	HttpRequest request.Request

	OpID  uint64
	PsmID uint64

	RetryStrategy api.DefaultRetryStrategy

	IsUnauthorizedFunc  auth.IsUnauthorizedFunc
	AuthRefresherFunc   auth.AuthRefresherFunc
	AuthProviderFactory auth.AuthProviderFactory
}

func (r ShippingRequest) GetIsUnauthorizedFunc() auth.IsUnauthorizedFunc {
	if r.IsUnauthorizedFunc != nil {
		return r.IsUnauthorizedFunc
	}
	return auth.DefaultIsUnauthorizedFunc
}

func (r ShippingRequest) GetAuthRefresherFunc() auth.AuthRefresherFunc {
	if r.AuthRefresherFunc != nil {
		return r.AuthRefresherFunc
	}
	return auth.DefaultAuthRefresherFunc
}

func (r ShippingRequest) GetAuthProviderFactory() auth.AuthProviderFactory {
	if r.AuthProviderFactory != nil {
		return r.AuthProviderFactory
	}
	return auth.DefaultAuthProviderFactory
}
