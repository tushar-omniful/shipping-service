package interfaces

import (
	"context"

	"github.com/omniful/shipping-service/internal/domain/models"

	"github.com/gin-gonic/gin"
	oerror "github.com/omniful/go_commons/error"
	"github.com/omniful/shipping-service/internal/services/orders/requests"
	"github.com/omniful/shipping-service/internal/services/orders/responses"
)

type OrderController interface {
	CreateOrder(ctx *gin.Context)
}

type OrderService interface {
	CreateOrder(ctx context.Context, request *requests.CreateForwardOrder) (*responses.GetOrderResponse, oerror.CustomError)
}

type OrderRepository interface {
	CreateOrder(
		ctx context.Context,
		order *models.Order,
	) (*models.Order, oerror.CustomError)
	CheckExistingOrder(
		ctx context.Context,
		orderPartnerOrderID string,
		orderPartnerID uint64,
	) (bool, oerror.CustomError)
}
