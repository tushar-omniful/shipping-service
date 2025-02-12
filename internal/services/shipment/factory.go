package shipment

import (
	"context"
	"fmt"
	"github.com/omniful/shipping-service/internal/services/shipment/requests"
	"github.com/omniful/shipping-service/internal/services/shipment/responses"

	"github.com/omniful/shipping-service/internal/domain/interfaces"
	"github.com/omniful/shipping-service/internal/services/shipment/integrations/ajex"
)

type ShippingService interface {
	CreateShipment(ctx context.Context, req *requests.CreateForwardShipmentRequest) (*responses.CreateForwardShipmentResponse, error)
	CancelShipment(ctx context.Context, req *requests.CancelShipmentRequest) (*responses.CancelShipmentResponse, error)
	TrackShipment(ctx context.Context, orderID string) (*responses.TrackShipmentResponse, error)
}

// NewShippingService creates a new shipping provider based on the tag
func NewShippingService(ctx context.Context, tag string, psmRepo interfaces.PartnerShippingMethodRepository, shippingPartnerRepo interfaces.ShippingPartnerRepository, cityMappingRepo interfaces.CityMappingRepository, tenantCityMappingRepo interfaces.TenantCityMappingRepository, hubMappingRepo interfaces.HubMappingRepository) (ShippingService, error) {
	switch tag {
	case "ajex":
		return ajex.NewService(ctx, psmRepo)
	default:
		return nil, fmt.Errorf("unsupported provider: %s", tag)
	}
}
