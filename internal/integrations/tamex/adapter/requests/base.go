package requests

import (
	"github.com/omniful/shipping-service/internal/domain/interfaces"
)

type RequestAdapter struct {
	spRepo  interfaces.ShippingPartnerRepository
	cmRepo  interfaces.CityMappingRepository
	tcmRepo interfaces.TenantCityMappingRepository
}

func NewRequestAdapter(spRepo interfaces.ShippingPartnerRepository, cmRepo interfaces.CityMappingRepository, tcmRepo interfaces.TenantCityMappingRepository) *RequestAdapter {
	return &RequestAdapter{
		spRepo:  spRepo,
		cmRepo:  cmRepo,
		tcmRepo: tcmRepo,
	}
}
