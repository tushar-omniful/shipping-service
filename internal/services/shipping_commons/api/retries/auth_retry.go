package retries

import (
	"fmt"
	"github.com/omniful/go_commons/config"
	"github.com/omniful/shipping-service/internal/domain/interfaces"
	"strconv"
	"time"

	"github.com/omniful/go_commons/dmutex"
	"github.com/omniful/go_commons/httpclient"
	"github.com/omniful/go_commons/httpclient/request"
	"github.com/omniful/go_commons/httpclient/response"
	"github.com/omniful/shipping-service/internal/services/shipping_commons/api/auth"
	"github.com/omniful/shipping-service/pkg/redis"
)

func NewAuthRetry(psmID uint64, isUnauthorizedFunc auth.IsUnauthorizedFunc, authRefresherFunc auth.AuthRefresherFunc, authProviderFactory auth.AuthProviderFactory, psm interfaces.PartnerShippingMethodRepository) httpclient.Retry {
	return &AuthRetry{
		psmID:               psmID,
		isUnauthorizedFunc:  isUnauthorizedFunc,
		authRefresherFunc:   authRefresherFunc,
		authProviderFactory: authProviderFactory,
		psm:                 psm,
	}
}

type AuthRetry struct {
	psmID               uint64
	isUnauthorizedFunc  auth.IsUnauthorizedFunc
	authRefresherFunc   auth.AuthRefresherFunc
	authProviderFactory auth.AuthProviderFactory
	psm                 interfaces.PartnerShippingMethodRepository
}

func (a *AuthRetry) ShouldRetry(ctx *httpclient.Context, req request.Request, resp response.Response) bool {
	if ctx.AttemptCount() > 1 {
		return false
	}
	return a.isUnauthorizedFunc(ctx, req, resp)
}

func (a *AuthRetry) NextAttemptIn(ctx *httpclient.Context, req request.Request, resp response.Response) time.Duration {
	return 0
}

func (a *AuthRetry) PrepareRequest(ctx *httpclient.Context, req request.Request, resp response.Response) (request.Request, error) {
	// Try to acquire lock for refreshing auth
	ttl := time.Duration(config.GetInt(ctx, "salesChannelHttpClient.authRefresh.lockTTL")) * time.Second
	dmx := dmutex.New(_getMutexKey(a.psmID), ttl, redis.GetClient().Client)
	acquired, err := dmx.TryLock(ctx)
	if err != nil {
		return req, err
	}

	if acquired {
		// We got the lock, refresh auth and unlock when done
		defer dmx.Unlock(ctx)
		return a.authRefresherFunc(ctx, req, a.psmID)
	}

	// Wait for lock to be released
	isTTlExpired, err := dmx.WaitUntilUnlocked(ctx)
	if err != nil {
		return req, err
	}
	if isTTlExpired {
		return a.authRefresherFunc(ctx, req, a.psmID)
	}

	// refresh from db
	psm, cusErr := a.psm.GetPartnerShippingMethodByID(ctx, strconv.FormatUint(a.psmID, 10))
	if cusErr.Exists() {
		return req, cusErr
	}
	ap, err := a.authProviderFactory(ctx, req, psm)
	if err != nil {
		return req, err
	}
	return ap.Apply(ctx, req)
}

func _getMutexKey(psmID uint64) string {
	return fmt.Sprintf("shs-auth_retry-%d", psmID)
}
