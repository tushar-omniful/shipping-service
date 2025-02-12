package order_controller

import (
	"github.com/google/wire"
	"github.com/omniful/shipping-service/internal/domain/interfaces"
	cityMappingRepo "github.com/omniful/shipping-service/internal/repositories/city_mapping"
	hubMappingRepo "github.com/omniful/shipping-service/internal/repositories/hub_mapping"
	orderRepo "github.com/omniful/shipping-service/internal/repositories/order"
	orderPartnerRepo "github.com/omniful/shipping-service/internal/repositories/order_partner"
	partnerShippingMethodRepo "github.com/omniful/shipping-service/internal/repositories/partner_shipping_method"
	shippingPartnerRepo "github.com/omniful/shipping-service/internal/repositories/shipping_partner"
	tenantCityMappingRepo "github.com/omniful/shipping-service/internal/repositories/tenant_city_mapping"
	orderService "github.com/omniful/shipping-service/internal/services/orders"
	shipmentService "github.com/omniful/shipping-service/internal/services/shipment"
)

var ProviderSet wire.ProviderSet = wire.NewSet(
	NewController,
	orderService.NewService,
	shipmentService.NewService,
	orderRepo.NewRepository,
	shippingPartnerRepo.NewRepository,
	partnerShippingMethodRepo.NewRepository,
	cityMappingRepo.NewRepository,
	tenantCityMappingRepo.NewRepository,
	orderPartnerRepo.NewRepository,
	hubMappingRepo.NewRepository,

	wire.Bind(new(interfaces.OrderController), new(*Controller)),
	wire.Bind(new(interfaces.OrderRepository), new(*orderRepo.Repository)),
	wire.Bind(new(interfaces.OrderService), new(*orderService.OrderService)),
	wire.Bind(new(interfaces.OrderPartnerRepository), new(*orderPartnerRepo.Repository)),
	wire.Bind(new(interfaces.ShippingPartnerRepository), new(*shippingPartnerRepo.Repository)),
	wire.Bind(new(interfaces.PartnerShippingMethodRepository), new(*partnerShippingMethodRepo.Repository)),
	wire.Bind(new(interfaces.CityMappingRepository), new(*cityMappingRepo.Repository)),
	wire.Bind(new(interfaces.TenantCityMappingRepository), new(*tenantCityMappingRepo.Repository)),
	wire.Bind(new(interfaces.HubMappingRepository), new(*hubMappingRepo.Repository)),
)
