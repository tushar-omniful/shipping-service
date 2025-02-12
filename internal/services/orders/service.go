package order_service

import (
	"context"
	"fmt"
	"github.com/omniful/shipping-service/internal/services/shipment"
	"strconv"
	"sync"

	"github.com/omniful/go_commons/error"
	"github.com/omniful/shipping-service/internal/domain/interfaces"
	"github.com/omniful/shipping-service/internal/services/orders/requests"
	"github.com/omniful/shipping-service/internal/services/orders/responses"
	ss_error "github.com/omniful/shipping-service/pkg/error"
)

type OrderService struct {
	OrderRepo                 interfaces.OrderRepository
	OrderPartnerRepo          interfaces.OrderPartnerRepository
	PartnerShippingMethodRepo interfaces.PartnerShippingMethodRepository
	ShippingPartnerRepo       interfaces.ShippingPartnerRepository
	CityMappingRepo           interfaces.CityMappingRepository
	TenantCityMappingRepo     interfaces.TenantCityMappingRepository
	HubMappingRepo            interfaces.HubMappingRepository
	ShipmentService           *shipment.Service
}

var (
	instance *OrderService
	once     sync.Once
)

func NewService(
	orderRepo interfaces.OrderRepository,
	orderPartnerRepo interfaces.OrderPartnerRepository,
	partnerShippingMethodRepo interfaces.PartnerShippingMethodRepository,
	cityMappingRepo interfaces.CityMappingRepository,
	shippingPartnerRepo interfaces.ShippingPartnerRepository,
	tenantCityMappingRepo interfaces.TenantCityMappingRepository,
	hubMappingRepo interfaces.HubMappingRepository,
	shipmentService *shipment.Service,
) *OrderService {
	once.Do(func() {
		instance = &OrderService{
			OrderRepo:                 orderRepo,
			OrderPartnerRepo:          orderPartnerRepo,
			ShippingPartnerRepo:       shippingPartnerRepo,
			PartnerShippingMethodRepo: partnerShippingMethodRepo,
			CityMappingRepo:           cityMappingRepo,
			TenantCityMappingRepo:     tenantCityMappingRepo,
			HubMappingRepo:            hubMappingRepo,
			ShipmentService:           shipmentService,
		}
	})
	return instance
}

func (s *OrderService) CreateOrder(ctx context.Context, req *requests.CreateForwardOrder) (*responses.GetOrderResponse, error.CustomError) {
	// Get order partner
	op, err := s.OrderPartnerRepo.GetOrderPartnerByTenantID(ctx, req.TenantID)
	if err.Exists() {
		return nil, error.NewCustomError(ss_error.NotFound, fmt.Sprintf("order partner not found for tenant ID %s: %v", req.TenantID, err))
	}

	// Get partner shipping method
	sp, err := s.ShippingPartnerRepo.GetShippingPartnerByTag(ctx, req.ShippingPartnerTag)
	if err.Exists() {
		return nil, error.NewCustomError(ss_error.NotFound, fmt.Sprintf("shipping partner not found  %s: %v", req.ShippingPartnerTag, err))
	}

	// Get partner shipping method
	psm, err := s.PartnerShippingMethodRepo.GetPartnerShippingMethodByID(ctx, req.AccountID, strconv.FormatUint(op.ID, 10))
	if err.Exists() {
		return nil, error.NewCustomError(ss_error.NotFound, fmt.Sprintf("partner shipping method not found for account ID %s: %v", req.AccountID, err))
	}

	// Transform to shipment request
	shipmentReq, transformErr := req.TransformToShipmentRequest(sp, psm)
	if transformErr != nil {
		return nil, error.NewCustomError(ss_error.SqlCreateError, fmt.Sprintf("failed to transform request: %v", transformErr))
	}

	// Create shipment using shipment service
	_, shipErr := s.ShipmentService.CreateShipment(ctx, shipmentReq)
	if shipErr != nil {
		return nil, error.NewCustomError(ss_error.FailedToCreateShipment, shipErr.Error())
	}

	// Create order record
	//order := &models.Order{
	//	OrderPartnerOrderID:     req.Data.OrderPartnerOrderID,
	//	OrderPartnerID:          int64(op.ID),
	//	PartnerShippingMethodID: int64(psm.ID),
	//	ShippingPartnerID:       int64(psm.ShippingPartnerID),
	//	Status:                  models.NewOrder,
	//	ShipmentType:            req.Data.ShipmentType,
	//	ShippingPartnerStatus:   shipmentResp.Status,
	//	PickupDetails:           req.Data.PickupDetails,
	//	DropDetails:             req.Data.DropDetails,
	//	ShipmentDetails:         req.Data.ShipmentDetails,
	//	TaxDetails:              req.Data.TaxDetails,
	//	Metadata:                &models.OrderMetadata{AwbLabel: shipmentResp.Label, AwbNumber: shipmentResp.TrackingNumber},
	//	ShippingLabel:           shipmentResp.Label,
	//	AWBNumber:               shipmentResp.TrackingNumber,
	//	SellerID:                req.SellerID,
	//	Source:                  models.SourceShippingAggregator,
	//}

	//createdOrder, createErr := s.OrderRepo.CreateOrder(ctx, order)
	//if createErr.Exists() {
	//	return nil, error.NewCustomError(ss_error.SqlCreateError, fmt.Sprintf("failed to create order record: %v", createErr))
	//}

	return &responses.GetOrderResponse{
		//ID: "createdOrder.ID",
	}, error.CustomError{}
}
