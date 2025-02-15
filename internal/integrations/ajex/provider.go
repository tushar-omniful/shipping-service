package ajex

import (
	"github.com/google/wire"
	"github.com/omniful/shipping-service/internal/domain/interfaces"

	// AJEX-specific adapters and API client.
	reqAdapter "github.com/omniful/shipping-service/internal/integrations/ajex/adapter/requests"
	respAdapter "github.com/omniful/shipping-service/internal/integrations/ajex/adapter/responses"
	"github.com/omniful/shipping-service/internal/integrations/ajex/api"
)

func ProvideRequestAdapter(spRepo interfaces.ShippingPartnerRepository, cmRepo interfaces.CityMappingRepository, tcmRepo interfaces.TenantCityMappingRepository) *reqAdapter.RequestAdapter {
	return reqAdapter.NewRequestAdapter(spRepo, cmRepo, tcmRepo)
}

func ProvideResponseAdapter(opRepo interfaces.OrderPartnerRepository, orderRepo interfaces.OrderRepository) *respAdapter.ResponseAdapter {
	return respAdapter.NewResponseAdapter(opRepo, orderRepo)
}

// ProviderSet provides all the dependencies for the AJEX integration.
//
// Note: The required repository parameters (such as PartnerShippingMethodRepository,
// ShippingPartnerRepository, CityMappingRepository, TenantCityMappingRepository,
// OrderRepository and OrderPartnerRepository) are expected to be provided
// by a higher-level injector.
var ProviderSet = wire.NewSet(
	// Create the API client (requires PartnerShippingMethodRepository).
	api.NewClient,
	// Create the Request Adapter (requires ShippingPartnerRepository, CityMappingRepository, TenantCityMappingRepository).
	ProvideRequestAdapter,
	// Create the Response Adapter (requires OrderPartnerRepository, OrderRepository).
	ProvideResponseAdapter,
	// Build the AJEX shipment service.
	NewService,
	// Bind the concrete type (*Service) to the common ShipmentService interface.
	wire.Bind(new(interfaces.ShipmentService), new(*Service)),
)
