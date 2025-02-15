package tamex

import (
	"context"
	"sync"

	customError "github.com/omniful/go_commons/error"
	"github.com/omniful/shipping-service/internal/domain/interfaces"
	requestAdapter "github.com/omniful/shipping-service/internal/integrations/tamex/adapter/requests"
	responseAdapter "github.com/omniful/shipping-service/internal/integrations/tamex/adapter/responses"
	apiClient "github.com/omniful/shipping-service/internal/integrations/tamex/api"
	"github.com/omniful/shipping-service/internal/services/shipment/requests"
	"github.com/omniful/shipping-service/internal/services/shipment/responses"
)

var (
	svc     *Service
	svcOnce sync.Once
)

type Service struct {
	*apiClient.Client
	reqAdapter  *requestAdapter.RequestAdapter
	respAdapter *responseAdapter.ResponseAdapter
}

func NewService(
	ctx context.Context,
	psmRepo interfaces.PartnerShippingMethodRepository,
	requestAdapter *requestAdapter.RequestAdapter,
	responseAdapter *responseAdapter.ResponseAdapter,
) *Service {
	svcOnce.Do(func() {
		svc = &Service{
			Client:      apiClient.NewClient(ctx, psmRepo),
			reqAdapter:  requestAdapter,
			respAdapter: responseAdapter,
		}
	})

	return svc
}

func (p *Service) CreateShipment(ctx context.Context,
	req *requests.CreateForwardShipmentRequest,
) (resp responses.CreateForwardShipmentResponse, err customError.CustomError) {
	// Format request
	reqBody, err := p.reqAdapter.FormatCreateShipment(req)
	if err.Exists() {
		return
	}

	// Send request
	response, err := p.Client.Create(ctx, reqBody, *req.PartnerShippingMethod)
	if err.Exists() {
		return
	}

	// Parse response
	resp, err = p.respAdapter.ParseCreateShipment(response)
	if err.Exists() {
		return
	}

	return
}

func (p *Service) CancelShipment(ctx context.Context,
	req *requests.CancelShipmentRequest,
) (resp responses.CancelShipmentResponse, err customError.CustomError) {
	// Format request
	reqBody, err := p.reqAdapter.FormatCancelShipment(req)
	if err.Exists() {
		return
	}

	// Send request
	response, err := p.Client.Cancel(ctx, reqBody, *req.PartnerShippingMethod)
	if err.Exists() {
		return
	}

	// Parse response
	resp, err = p.respAdapter.ParseCancelShipment(response)
	if err.Exists() {
		return
	}

	return
}

func (p *Service) TrackShipment(ctx context.Context,
	req *requests.TrackShipmentRequest,
) (resp responses.TrackShipmentResponse, err customError.CustomError) {
	// Format request
	reqBody, err := p.reqAdapter.FormatTrackShipment(req)
	if err.Exists() {
		return
	}

	// Send request
	response, err := p.Client.Track(ctx, reqBody, *req.PartnerShippingMethod)
	if err.Exists() {
		return
	}

	// Parse response
	resp, err = p.respAdapter.ParseTrackShipment(response)
	if err.Exists() {
		return
	}

	return
}
