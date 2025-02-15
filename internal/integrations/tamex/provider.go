package tamex

import (
	"github.com/google/wire"
	"github.com/omniful/shipping-service/internal/domain/interfaces"

	// Aramex-specific adapters and API client.
	reqAdapter "github.com/omniful/shipping-service/internal/integrations/tamex/adapter/requests"
	respAdapter "github.com/omniful/shipping-service/internal/integrations/tamex/adapter/responses"
	"github.com/omniful/shipping-service/internal/integrations/tamex/api"
)

func ProvideRequestAdapter(spRepo interfaces.ShippingPartnerRepository, cmRepo interfaces.CityMappingRepository, tcmRepo interfaces.TenantCityMappingRepository) *reqAdapter.RequestAdapter {
	return reqAdapter.NewRequestAdapter(spRepo, cmRepo, tcmRepo)
}

func ProvideResponseAdapter(opRepo interfaces.OrderPartnerRepository, orderRepo interfaces.OrderRepository) *respAdapter.ResponseAdapter {
	return respAdapter.NewResponseAdapter(opRepo, orderRepo)
}

var ProviderSet = wire.NewSet(
	api.NewClient,
	ProvideRequestAdapter,
	ProvideResponseAdapter,
	NewService,
	wire.Bind(new(interfaces.ShipmentService), new(*Service)),
)
