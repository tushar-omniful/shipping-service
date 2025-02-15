package shipping_commons

import (
	"context"
	"github.com/omniful/go_commons/log"
	cityMappingRepo "github.com/omniful/shipping-service/internal/repositories/city_mapping"
	orderRepo "github.com/omniful/shipping-service/internal/repositories/order"
	orderPartnerRepo "github.com/omniful/shipping-service/internal/repositories/order_partner"
	partnerShippingMethodRepo "github.com/omniful/shipping-service/internal/repositories/partner_shipping_method"
	shippingPartnerRepo "github.com/omniful/shipping-service/internal/repositories/shipping_partner"
	tenantCityMappingRepo "github.com/omniful/shipping-service/internal/repositories/tenant_city_mapping"
	"github.com/omniful/shipping-service/pkg/db/postgres"
	"sync"

	"github.com/omniful/shipping-service/internal/domain/interfaces"

	// Import the generated wire functions for each integration.
	"github.com/omniful/shipping-service/internal/integrations/ajex"
	"github.com/omniful/shipping-service/internal/integrations/tamex"
)

// CommonShippingService aggregates all shipping service providers keyed by a tag.
type CommonShippingService struct {
	providers map[string]interfaces.ShipmentService
}

var (
	commonShippingInstance *CommonShippingService
	once                   sync.Once
)

func InitializeCommonShippingService(ctx context.Context) {
	once.Do(func() {
		db := postgres.GetCluster().DbCluster
		orderRep := orderRepo.NewRepository(db)
		psmRepo := partnerShippingMethodRepo.NewRepository(db)
		cmRepo := cityMappingRepo.NewRepository(db)
		spRepo := shippingPartnerRepo.NewRepository(db)
		tcmRepo := tenantCityMappingRepo.NewRepository(db)
		opRepo := orderPartnerRepo.NewRepository(db)
		//hmRepo := hub_mapping_repo.NewRepository(db)

		ajexService, err := ajex.AjexService(ctx, psmRepo, spRepo, cmRepo, tcmRepo, orderRep, opRepo)
		if err != nil {
			log.Errorf(err.Error())
		}

		tamexService, err := tamex.TamexService(ctx, psmRepo, spRepo, cmRepo, tcmRepo, orderRep, opRepo)
		if err != nil {
			log.Errorf(err.Error())
		}

		providers := map[string]interfaces.ShipmentService{
			"ajex":  ajexService,
			"tamex": tamexService,
		}
		commonShippingInstance = &CommonShippingService{
			providers: providers,
		}
	})
	log.Infof("Initialized Common Shipping Service Client")
}

// GetShippingService returns the shipping service that exactly matches the provided tag.
// This function panics if the common shipping service is not initialized or the tag is unknown.
func GetShippingService(tag string) interfaces.ShipmentService {
	if commonShippingInstance == nil {
		panic("common shipping service not initialized")
	}
	svc, ok := commonShippingInstance.providers[tag]
	if !ok {
		panic("shipping service with tag \"" + tag + "\" not found")
	}
	return svc
}
