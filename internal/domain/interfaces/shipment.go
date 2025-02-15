package interfaces

import (
	"context"
	customError "github.com/omniful/go_commons/error"

	"github.com/omniful/shipping-service/internal/services/shipment/requests"
	"github.com/omniful/shipping-service/internal/services/shipment/responses"
)

type ShipmentService interface {
	CreateShipment(ctx context.Context, req *requests.CreateForwardShipmentRequest) (resp responses.CreateForwardShipmentResponse, err customError.CustomError)
}
