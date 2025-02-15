//go:build wireinject
// +build wireinject

package ajex

import (
	"context"
	"github.com/google/wire"
	"github.com/omniful/shipping-service/internal/domain/interfaces"
)

func AjexService(
	ctx context.Context,
	psmRepo interfaces.PartnerShippingMethodRepository,
	spRepo interfaces.ShippingPartnerRepository,
	cmRepo interfaces.CityMappingRepository,
	tcmRepo interfaces.TenantCityMappingRepository,
	orderRepo interfaces.OrderRepository,
	opRepo interfaces.OrderPartnerRepository,
) (*Service, error) {
	wire.Build(ProviderSet)
	return nil, nil
}
