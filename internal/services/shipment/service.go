package shipment

import (
	"context"
	"fmt"
	"sync"

	"github.com/omniful/shipping-service/internal/domain/interfaces"
	"github.com/omniful/shipping-service/internal/services/shipment/requests"
	"github.com/omniful/shipping-service/internal/services/shipment/responses"
)

type Service struct {
	mu                        sync.RWMutex
	shippingServices          map[string]ShippingService // Cache of service instances by tag
	OrderRepo                 interfaces.OrderRepository
	OrderPartnerRepo          interfaces.OrderPartnerRepository
	PartnerShippingMethodRepo interfaces.PartnerShippingMethodRepository
	ShippingPartnerRepo       interfaces.ShippingPartnerRepository
	CityMappingRepo           interfaces.CityMappingRepository
	TenantCityMappingRepo     interfaces.TenantCityMappingRepository
	HubMappingRepo            interfaces.HubMappingRepository
}

var (
	instance *Service
	once     sync.Once
)

// getService returns a cached service instance or creates a new one
func (s *Service) getShippingService(ctx context.Context, tag string) (ShippingService, error) {
	s.mu.RLock()
	if shippingService, exists := s.shippingServices[tag]; exists {
		s.mu.RUnlock()
		return shippingService, nil
	}
	s.mu.RUnlock()

	// Create new service if not in cache
	s.mu.Lock()
	defer s.mu.Unlock()

	// Double-check after acquiring write lock
	if shippingService, exists := s.shippingServices[tag]; exists {
		return shippingService, nil
	}

	shippingService, err := NewShippingService(ctx, tag, s.PartnerShippingMethodRepo, s.ShippingPartnerRepo, s.CityMappingRepo, s.TenantCityMappingRepo, s.HubMappingRepo)
	if err != nil {
		return nil, fmt.Errorf("failed to create service: %w", err)
	}

	s.shippingServices[tag] = shippingService
	return shippingService, nil
}

func NewService(
	orderRepo interfaces.OrderRepository,
	orderPartnerRepo interfaces.OrderPartnerRepository,
	partnerShippingMethodRepo interfaces.PartnerShippingMethodRepository,
	cityMappingRepo interfaces.CityMappingRepository,
	shippingPartnerRepo interfaces.ShippingPartnerRepository,
	tenantCityMappingRepo interfaces.TenantCityMappingRepository,
	hubMappingRepo interfaces.HubMappingRepository,
) *Service {
	once.Do(func() {
		instance = &Service{
			shippingServices:          make(map[string]ShippingService),
			OrderRepo:                 orderRepo,
			ShippingPartnerRepo:       shippingPartnerRepo,
			OrderPartnerRepo:          orderPartnerRepo,
			PartnerShippingMethodRepo: partnerShippingMethodRepo,
			CityMappingRepo:           cityMappingRepo,
			TenantCityMappingRepo:     tenantCityMappingRepo,
			HubMappingRepo:            hubMappingRepo,
		}
	})

	return instance
}

// CreateShipment handles the creation of a new shipment
func (s *Service) CreateShipment(ctx context.Context, req *requests.CreateForwardShipmentRequest) (*responses.CreateForwardShipmentResponse, error) {
	// Get cached service instance
	shippingService, err := s.getShippingService(ctx, req.ShippingPartner.Tag)
	if err != nil {
		return nil, fmt.Errorf("failed to get shippingService: %w", err)
	}

	// Create shipment using service
	resp, createErr := shippingService.CreateShipment(ctx, req)
	if createErr != nil {
		return nil, fmt.Errorf("failed to create shipment: %w", createErr)
	}
	return resp, nil
}

// CancelShipment handles the cancellation of a shipment
func (s *Service) CancelShipment(ctx context.Context, req *requests.CancelShipmentRequest) (*responses.CancelShipmentResponse, error) {
	// Get cached service instance
	shippingService, err := s.getShippingService(ctx, "ajex")
	if err != nil {
		return nil, fmt.Errorf("failed to get shippingService: %w", err)
	}

	// Cancel shipment using shippingService
	return shippingService.CancelShipment(ctx, req)
}

// TrackShipment handles tracking of a shipment
func (s *Service) TrackShipment(ctx context.Context) (*responses.TrackShipmentResponse, error) {
	// Get cached shippingService instance
	//shippingService, err := s.shippingService(ctx, "ajex")
	//if err != nil {
	//	return nil, fmt.Errorf("failed to get shippingService: %w", err)
	//}

	// Track shipment using shippingService
	return &responses.TrackShipmentResponse{}, nil
	//return shippingService.TrackShipment(ctx, "56")
}
