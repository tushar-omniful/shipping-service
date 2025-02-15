package shipment

import (
	"context"
	"sync"

	"github.com/omniful/shipping-service/pkg/lock"

	customError "github.com/omniful/go_commons/error"
	"github.com/omniful/shipping-service/internal/domain/interfaces"
	"github.com/omniful/shipping-service/internal/services/shipment/requests"
	"github.com/omniful/shipping-service/internal/services/shipment/responses"
	"github.com/omniful/shipping-service/internal/services/shipping_commons"
)

type Service struct {
	mu                        sync.RWMutex
	OrderRepo                 interfaces.OrderRepository
	OrderPartnerRepo          interfaces.OrderPartnerRepository
	PartnerShippingMethodRepo interfaces.PartnerShippingMethodRepository
	ShippingPartnerRepo       interfaces.ShippingPartnerRepository
	CityMappingRepo           interfaces.CityMappingRepository
	TenantCityMappingRepo     interfaces.TenantCityMappingRepository
	HubMappingRepo            interfaces.HubMappingRepository
	locker                    lock.Locker
}

var (
	instance *Service
	once     sync.Once
)

func (s *Service) getShipmentService(tag string) (interfaces.ShipmentService, customError.CustomError) {
	return shipping_commons.GetShippingService(tag), customError.CustomError{}
}

func NewService(
	orderRepo interfaces.OrderRepository,
	orderPartnerRepo interfaces.OrderPartnerRepository,
	partnerShippingMethodRepo interfaces.PartnerShippingMethodRepository,
	shippingPartnerRepo interfaces.ShippingPartnerRepository,
	locker lock.Locker,
) *Service {
	once.Do(func() {
		instance = &Service{
			OrderRepo:                 orderRepo,
			ShippingPartnerRepo:       shippingPartnerRepo,
			OrderPartnerRepo:          orderPartnerRepo,
			PartnerShippingMethodRepo: partnerShippingMethodRepo,
			locker:                    locker,
		}
	})

	return instance
}

// CreateShipment handles the creation of a new shipment
func (s *Service) CreateShipment(ctx context.Context, req *requests.CreateForwardShipmentRequest) (resp responses.CreateForwardShipmentResponse, err customError.CustomError) {
	// Get cached service instance
	shippingService, err := s.getShipmentService(req.ShippingPartner.Tag)
	if err.Exists() {
		return resp, err
	}

	resp, err = shippingService.CreateShipment(ctx, req)

	if err.Exists() {
		return resp, err
	}
	return resp, err
}

// CancelShipment handles the cancellation of a shipment
//func (s *Service) CancelShipment(ctx context.Context, req *requests.CancelShipmentRequest) (*responses.CancelShipmentResponse, error) {
//	// Get cached service instance
//	shippingService, err := s.getShippingService(ctx, "ajex")
//	if err != nil {
//		return nil, fmt.Errorf("failed to get shippingService: %w", err)
//	}
//
//	// Cancel shipment using shippingService
//	return shippingService.CancelShipment(ctx, req)
//}

// TrackShipment handles tracking of a shipment
//func (s *Service) TrackShipment(ctx context.Context) (*responses.TrackShipmentResponse, error) {
//	// Get cached shippingService instance
//	//shippingService, err := s.shippingService(ctx, "ajex")
//	//if err != nil {
//	//	return nil, fmt.Errorf("failed to get shippingService: %w", err)
//	//}
//
//	// Track shipment using shippingService
//	return &responses.TrackShipmentResponse{}, nil
//	//return shippingService.TrackShipment(ctx, "56")
//}
