package ajex

import (
	"context"
	"sync"

	customError "github.com/omniful/go_commons/error"
	"github.com/omniful/shipping-service/internal/domain/interfaces"
	requestAdapter "github.com/omniful/shipping-service/internal/integrations/ajex/adapter/requests"
	responseAdapter "github.com/omniful/shipping-service/internal/integrations/ajex/adapter/responses"
	apiClient "github.com/omniful/shipping-service/internal/integrations/ajex/api"
	"github.com/omniful/shipping-service/internal/services/shipment/requests"
	"github.com/omniful/shipping-service/internal/services/shipment/responses"
	shsError "github.com/omniful/shipping-service/pkg/error"
)

var (
	svc     *Service
	svcOnce sync.Once
)

type Service struct {
	*apiClient.Client
	reqAdapter  *requestAdapter.RequestAdapter
	respAdapter *responseAdapter.ResponseAdapter
	retryCount  int
	maxRetries  int
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
			retryCount:  1,
		}
	})

	return svc
}

func (p *Service) trackByReference(ctx context.Context,
	req *requests.CreateForwardShipmentRequest,
) (resp responseAdapter.TrackByReferenceResponse, err customError.CustomError) {
	return resp, customError.CustomError{}
}

func (p *Service) CreateShipment(ctx context.Context, req *requests.CreateForwardShipmentRequest) (resp responses.CreateForwardShipmentResponse, err customError.CustomError) {

	reqBody, err := p.reqAdapter.FormatCreateShipment(req)

	response, err := p.Client.Create(ctx, reqBody, *req.PartnerShippingMethod)

	if err.Exists() {
		return
	}

	// Parse response
	resp, err = p.respAdapter.ParseCreateShipment(response)
	if err.Exists() {
		resp, err = p.handleCreateError(ctx, err, req)
		if err.Exists() {
			return
		}
		return
	}

	return
}

func (p *Service) handleCreateError(ctx context.Context,
	err customError.CustomError, originalReq *requests.CreateForwardShipmentRequest,
) (resp responses.CreateForwardShipmentResponse, error customError.CustomError) {

	if p.retryCount >= p.maxRetries {
		return resp, err
	}

	switch err.ErrorCode() {
	case shsError.OrderReferenceExistError:
		// Returning error from shipment creation not from tracking in case of failure
		trackResp, error2 := p.trackByReference(ctx, originalReq)
		if error2.Exists() {
			return resp, err
		}
		resp, error2 = p.respAdapter.ParseCreateByReference(trackResp, originalReq)
		if error2.Exists() {
			return resp, err
		}
	default:
		return resp, err
	}

	return resp, err
}

//func (p *Service) CancelShipment(ctx context.Context, req *requests.CancelShipmentRequest) (*responses.CancelShipmentResponse, error) {
//	// Get auth token
//	//token, err := p.authAdapter.GetAuthToken(req.PartnerShippingMethod)
//	//if err != nil {
//	//	return nil, fmt.Errorf("get auth token: %w", err)
//	//}
//
//	// Format request
//	reqBody, err := p.ReqAdapter.FormatCancelShipment(req)
//	if err != nil {
//		return nil, fmt.Errorf("format request: %w", err)
//	}
//
//	// Send request with auth token
//	headers := map[string][]string{
//		"Content-Type": {"application/json"},
//		//"Authorization": {token},
//	}
//	respBody, err := p.SendRESTRequest(ctx, "POST", "/shipments/cancel", headers, reqBody, nil)
//	if err != nil {
//		return nil, err
//	}
//
//	// Parse response
//	resp, err := p.RespAdapter.ParseCancelShipment(respBody)
//	//if err != nil {
//	// Handle errors and retry if needed
//	//return p.handleCancelError(ctx, err, req, req.PartnerShippingMethod)
//	//}
//
//	return resp, nil
//}
//
//func (p *AjexService) TrackShipment(ctx context.Context, orderID string) (*responses.TrackShipmentResponse, error) {
//	// Format request
//	reqBody, err := p.ReqAdapter.FormatTrackShipment(&requests.TrackShipmentRequest{
//		OrderID: orderID,
//	})
//	if err != nil {
//		return nil, fmt.Errorf("format request: %w", err)
//	}
//
//	// Send request
//	headers := map[string][]string{
//		"Accept": {"application/json"},
//	}
//	respBody, err := p.SendRESTRequest(ctx, "GET", fmt.Sprintf("/shipments/%s/track", orderID), headers, reqBody, nil)
//	if err != nil {
//		return nil, fmt.Errorf("send request: %w", err)
//	}
//
//	// Parse response
//	return p.RespAdapter.ParseTrackShipment(respBody)
//}

//func (p *AjexService) handleCancelError(ctx context.Context, err error, originalReq *requests.CancelShipmentRequest, psm *models.PartnerShippingMethod) (*responses.CancelShipmentResponse, error) {
//	if p.retryCount >= p.maxRetries {
//		return nil, fmt.Errorf("max retries exceeded: %v", err)
//	}
//
//	switch e := err.(type) {
//	case *errors.TokenError:
//		p.retryCount++
//		// Invalidate token and retry
//		// updatedPSM, refreshErr := p.authAdapter.HandleUnauthorized(ctx, psm)
//		// if refreshErr != nil {
//		// 	return nil, fmt.Errorf("token refresh failed: %w", refreshErr)
//		// }
//		return p.CancelShipment(ctx, originalReq)
//	default:
//		return nil, e
//	}
//}
