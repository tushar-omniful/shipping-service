package retries

import (
	"github.com/omniful/go_commons/httpclient"
	"github.com/omniful/go_commons/httpclient/request"
	"github.com/omniful/go_commons/httpclient/response"
	"github.com/omniful/shipping-service/internal/services/shipping_commons/api/util"
	"github.com/omniful/shipping-service/pkg/api"
	"golang.org/x/exp/slices"
	"time"
)

func NewDefaultRetry(rs api.DefaultRetryStrategy) httpclient.Retry {
	return &DefaultRetry{rs}
}

type DefaultRetry struct {
	rs api.DefaultRetryStrategy
}

func (r DefaultRetry) ShouldRetry(ctx *httpclient.Context, req request.Request, resp response.Response) bool {
	if ctx.AttemptCount() >= r.rs.MaxRetries {
		return false
	}
	if slices.Contains(r.rs.ServerDownStatusCodes, resp.StatusCode()) {
		return false
	}
	return util.IsServerDown(r.rs, resp)
}

func (r DefaultRetry) NextAttemptIn(ctx *httpclient.Context, req request.Request, resp response.Response) time.Duration {
	return r.rs.BackoffDuration * time.Duration(1<<ctx.AttemptCount())
}

func (r DefaultRetry) PrepareRequest(ctx *httpclient.Context, req request.Request, resp response.Response) (request.Request, error) {
	return req, nil
}
