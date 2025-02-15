package util

import (
	"github.com/omniful/go_commons/httpclient/response"
	"github.com/omniful/shipping-service/pkg/api"
	"golang.org/x/exp/slices"
)

func IsServerDown(rs api.DefaultRetryStrategy, resp response.Response) bool {
	if resp == nil {
		return false
	}

	return slices.Contains(rs.ServerDownStatusCodes, resp.StatusCode())
}
