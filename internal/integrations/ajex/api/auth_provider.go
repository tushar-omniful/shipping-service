package api

import (
	"fmt"

	"github.com/omniful/go_commons/httpclient"
	"github.com/omniful/go_commons/httpclient/request"
	"github.com/omniful/shipping-service/constants"
	"github.com/omniful/shipping-service/internal/domain/models"
)

func NewAuthProvider(psm models.PartnerShippingMethod) httpclient.AuthProvider {
	return &AuthProvider{
		psm: psm,
	}
}

type AuthProvider struct {
	psm models.PartnerShippingMethod
}

func (a AuthProvider) Apply(ctx *httpclient.Context, req request.Request) (request.Request, error) {
	req.GetHeaders().Set(constants.Authorization, fmt.Sprintf("%v %v", a.psm.Credentials["token_type"], a.psm.Credentials["access_token"]))
	return req, nil
}
