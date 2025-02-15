package order_service

import (
	"context"
	"fmt"
	"github.com/omniful/go_commons/log"
	logsUtils "github.com/omniful/shipping-service/utils/logs"
	"strconv"
	"sync"

	"github.com/omniful/shipping-service/internal/services/shipment"

	"github.com/omniful/go_commons/error"
	"github.com/omniful/shipping-service/internal/domain/interfaces"
	"github.com/omniful/shipping-service/internal/domain/models"
	"github.com/omniful/shipping-service/internal/services/orders/requests"
	"github.com/omniful/shipping-service/internal/services/orders/responses"
	ss_error "github.com/omniful/shipping-service/pkg/error"
	"github.com/omniful/shipping-service/pkg/lock"
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
	locker                    lock.Locker
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
	locker lock.Locker,
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
			locker:                    locker,
		}
	})
	return instance
}

func (s *OrderService) CreateOrder(ctx context.Context, req *requests.CreateForwardOrder) (*responses.GetOrderResponse, error.CustomError) {
	logTag := logsUtils.GetLogTag(ctx, "CreateOrder")

	op, err := s.getOrderPartner(ctx, req.TenantID)
	if err.Exists() {
		return nil, err
	}

	// Get partner shipping method
	sp, err := s.getShippingPartner(ctx, req.ShippingPartnerTag)
	if err.Exists() {
		return nil, err
	}

	// Get partner shipping method
	psm, err := s.PartnerShippingMethodRepo.GetIntegratedPartnerShippingMethod(ctx, req.AccountID, strconv.FormatUint(op.ID, 10), req.SellerID)
	if err.Exists() {
		return nil, error.NewCustomError(ss_error.NotFound, fmt.Sprintf("partner shipping method not found for account ID %s: %v", req.AccountID, err))
	}

	// Transform to shipment request
	shipmentReq, transformErr := req.TransformToShipmentRequest(&sp, &psm)
	if transformErr != nil {
		return nil, error.NewCustomError(ss_error.SqlCreateError, fmt.Sprintf("failed to transform request: %v", transformErr))
	}

	exists, checkErr := s.OrderRepo.CheckExistingOrder(ctx, shipmentReq.Data.OrderPartnerOrderID, op.ID)
	if checkErr.Exists() {
		return nil, checkErr
	}
	if exists {
		return nil, error.NewCustomError(ss_error.BadRequest, fmt.Sprintf("Order with this reference is already present!"))
	}

	// Create shipment using shipment service
	shipmentResp, shipErr := s.ShipmentService.CreateShipment(ctx, shipmentReq)
	if shipErr.Exists() {
		return nil, shipErr
	}

	// Create order record
	order := TransformCreateShipmentResponseToOrderModel(shipmentReq, &shipmentResp, &op, &psm)

	createdOrder, createErr := s.OrderRepo.CreateOrder(ctx, order)
	if createErr.Exists() {
		log.Errorf("%s - Failed to save order: %v", logTag, err)
		return nil, error.NewCustomError(ss_error.ShippingServiceInternalError, fmt.Sprintf("failed to create order record: %v", createErr))
	}

	return responses.ConvertOrderModelToCreateResponse(createdOrder), error.CustomError{}
}

func (s *OrderService) getOrderPartner(ctx context.Context, tenantID string) (op models.OrderPartner, err error.CustomError) {
	op, err = s.OrderPartnerRepo.GetOrderPartnerByTenantID(ctx, tenantID)
	if err.Exists() {
		err = error.NewCustomError(ss_error.BadRequest, fmt.Sprintf("Invalid Tenant ID"))
		return
	}
	return
}

func (s *OrderService) getShippingPartner(ctx context.Context, spTag string) (sp models.ShippingPartner, err error.CustomError) {
	sp, err = s.ShippingPartnerRepo.GetShippingPartnerByTag(ctx, spTag)
	if err.Exists() {
		err = error.NewCustomError(ss_error.BadRequest, fmt.Sprintf("Invalid shipping company selected."))
		return
	}
	return
}

func (s *OrderService) getPartnerShippingMethod(ctx context.Context, accountID string, opID string, sellerID *string,
) (psm models.PartnerShippingMethod, err error.CustomError) {
	psm, err = s.PartnerShippingMethodRepo.GetPartnerShippingMethodByID(ctx, accountID)
	if err.Exists() {
		err = error.NewCustomError(ss_error.BadRequest, fmt.Sprintf("Invalid shipping account selected."))
		return
	}

	if sellerID != nil {
		if psm.IsAllSellerMapped == true {
			return
		}
	}
	return
}

func getPartnerShippingMethod() {}

func getOrderPartner() {}
