package router

import (
	"context"
	"github.com/omniful/go_commons/http"
	"github.com/omniful/go_commons/pagination"
	"github.com/omniful/shipping-service/internal/controllers/order"
	"github.com/omniful/shipping-service/pkg/db/postgres"
	"github.com/omniful/shipping-service/pkg/redis"
)

func InternalRoutes(ctx context.Context, s *http.Server) (err error) {

	s.Engine.Use(pagination.Middleware())

	shippingRouter := s.Engine.Group("/api/v1/orders")
	{

		orderController, ctrlErr := order_controller.Wire(
			ctx,
			postgres.GetCluster().DbCluster,
			redis.GetClient().Client,
		)
		if ctrlErr != nil {
			err = ctrlErr
			return
		}

		shippingRouter.POST("", orderController.CreateOrder)
	}
	return nil
}
