package router

import (
	"context"
	"fmt"

	"github.com/omniful/go_commons/http"

	"github.com/newrelic/go-agent/v3/integrations/nrgin"
	"github.com/omniful/go_commons/config"
	"github.com/omniful/go_commons/health"
	"github.com/omniful/go_commons/log"
	"github.com/omniful/go_commons/newrelic"
)

const (
	DefaultPerPageLimit = 100
)

func Initialize(ctx context.Context, s *http.Server) (err error) {
	fmt.Println("Initializing router")
	//Middleware for newrelic tracking
	s.Engine.Use(nrgin.Middleware(newrelic.GetApplication()))

	//Middleware for adding config to ctx
	s.Engine.Use(config.Middleware())

	s.Engine.Use(log.RequestLogMiddleware(log.MiddlewareOptions{
		Format:      config.GetString(ctx, "log.format"),
		Level:       config.GetString(ctx, "log.level"),
		LogRequest:  config.GetBool(ctx, "log.request"),
		LogResponse: config.GetBool(ctx, "log.response"),
	}))

	// health check route
	s.GET("/health", health.HealthcheckHandler())

	err = InternalRoutes(ctx, s)
	if err != nil {
		return
	}

	return
}
