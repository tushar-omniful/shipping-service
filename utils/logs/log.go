package logs

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	commonConstants "github.com/omniful/go_commons/constants"
	"github.com/omniful/go_commons/env"
	"github.com/omniful/go_commons/log"
	"github.com/omniful/go_commons/newrelic"
)

func SetRandomRequestID(ctx context.Context, systemID string) context.Context {
	return context.WithValue(ctx, commonConstants.HeaderXOmnifulRequestID, fmt.Sprintf("%s ,  %s",
		uuid.New().String(), systemID))
}

func GetRequestID(ctx context.Context, systemID string) (string, context.Context) {
	requestID, ok := ctx.Value(commonConstants.HeaderXOmnifulRequestID).(string)
	if !ok {
		ctx = SetRandomRequestID(ctx, systemID)

		return GetRequestID(ctx, systemID)
	}

	return requestID, ctx
}

func GetLogTag(ctx context.Context, funcName string) string {
	return fmt.Sprintf(" Request ID: %s Function: %s ", env.GetRequestID(ctx), funcName)
}

func NewrelicLogError(ctx context.Context, template string) {
	log.Error(template)
	newrelic.NoticeError(ctx, fmt.Errorf(template))
}
