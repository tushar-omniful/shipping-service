package auth

import (
	"github.com/omniful/go_commons/httpclient"
	"github.com/omniful/go_commons/httpclient/request"
	"github.com/omniful/go_commons/httpclient/response"
	"github.com/omniful/shipping-service/internal/domain/models"
)

type IsUnauthorizedFunc func(*httpclient.Context, request.Request, response.Response) bool
type AuthRefresherFunc func(ctx *httpclient.Context, req request.Request, psmID uint64) (request.Request, error)
type AuthProviderFactory func(ctx *httpclient.Context, req request.Request, psm models.PartnerShippingMethod) (httpclient.AuthProvider, error)

var DefaultIsUnauthorizedFunc IsUnauthorizedFunc = func(*httpclient.Context, request.Request, response.Response) bool { return false }
var DefaultAuthRefresherFunc AuthRefresherFunc = func(ctx *httpclient.Context, req request.Request, psmID uint64) (request.Request, error) {
	return req, nil
}
var DefaultAuthProviderFactory AuthProviderFactory = func(ctx *httpclient.Context, req request.Request, psm models.PartnerShippingMethod) (httpclient.AuthProvider, error) {
	return NewEmptyAuthProvider(), nil
}
