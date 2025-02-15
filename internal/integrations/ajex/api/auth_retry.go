package api

import (
	"github.com/omniful/go_commons/httpclient"
	"github.com/omniful/go_commons/httpclient/request"
	"github.com/omniful/go_commons/httpclient/response"
	"github.com/omniful/shipping-service/constants"
	"github.com/omniful/shipping-service/internal/domain/interfaces"
	"github.com/omniful/shipping-service/internal/domain/models"
	"github.com/omniful/shipping-service/internal/integrations/ajex/adapter/responses"
	"github.com/omniful/shipping-service/internal/services/shipping_commons/api/auth"
)

var isUnauthorizedFunc auth.IsUnauthorizedFunc = func(ctx *httpclient.Context, req request.Request, resp response.Response) bool {
	var r responses.CommonResponse
	if err := resp.UnmarshalBody(&r); err != nil {
		return false
	}
	return r.IsUnauthorized()
}

func buildAuthRefresherFunc(psmRepo interfaces.PartnerShippingMethodRepository, psm models.PartnerShippingMethod) auth.AuthRefresherFunc {
	return func(ctx *httpclient.Context, req request.Request, psmID uint64) (request.Request, error) {
		logResp, cusErr := GetClient().generateAccessToken(ctx, psm)
		if cusErr.Exists() {
			return req, cusErr
		}
		psm.Credentials["action_token"] = logResp.AccessToken
		psm.Credentials["token_type"] = logResp.TokenType

		cusErr = psmRepo.UpdatePartnerShippingMethod(ctx, map[string]any{
			constants.ID: psm.ID,
		}, psm)

		if cusErr.Exists() {
			return req, cusErr
		}

		return NewAuthProvider(psm).Apply(ctx, req)
	}
}

func buildAuthProviderFactory(psm models.PartnerShippingMethod) auth.AuthProviderFactory {
	return func(ctx *httpclient.Context, req request.Request, psm models.PartnerShippingMethod) (httpclient.AuthProvider, error) {
		return NewAuthProvider(psm), nil
	}
}
