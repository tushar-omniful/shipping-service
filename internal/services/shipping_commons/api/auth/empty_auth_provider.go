package auth

import (
	"github.com/omniful/go_commons/httpclient"
	"github.com/omniful/go_commons/httpclient/request"
)

func NewEmptyAuthProvider() httpclient.AuthProvider {
	return &EmptyAuthProvider{}
}

type EmptyAuthProvider struct{}

func (e EmptyAuthProvider) Apply(ctx *httpclient.Context, req request.Request) (request.Request, error) {
	return req, nil
}
