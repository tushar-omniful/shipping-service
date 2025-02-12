package ajex

import (
	"context"
	"fmt"
	"github.com/omniful/shipping-service/pkg/api"
	"time"

	appConfig "github.com/omniful/go_commons/config"
	"github.com/omniful/shipping-service/internal/domain/interfaces"
	"github.com/omniful/shipping-service/internal/domain/models"
	requestAdapter "github.com/omniful/shipping-service/internal/services/shipment/integrations/ajex/adapter/requests"
	responseAdapter "github.com/omniful/shipping-service/internal/services/shipment/integrations/ajex/adapter/responses"
	ajexClient "github.com/omniful/shipping-service/internal/services/shipment/integrations/ajex/client"
	"github.com/omniful/shipping-service/internal/services/shipment/integrations/ajex/errors"
	"github.com/omniful/shipping-service/internal/services/shipment/requests"
	"github.com/omniful/shipping-service/internal/services/shipment/responses"
)

type ShippingService struct {
	*ajexClient.BaseService
	authAdapter *AuthTokenAdapter
	retryCount  int
	maxRetries  int
}

func NewService(ctx context.Context, psmRepo interfaces.PartnerShippingMethodRepository) (*ShippingService, error) {
	// Increase timeout for AJEX API
	config := ajexClient.Config{
		BaseURL: appConfig.GetString(ctx, "ajex.host"),
		Timeout: appConfig.GetDuration(ctx, "ajex.timeout") * time.Second,
	}

	// Initialize adapters
	// reqAdapter := adapter.NewRequestAdapter()
	reqAdapter := requestAdapter.NewRequestAdapter()
	// respAdapter := adapter.NewResponseAdapter()
	respAdapter := responseAdapter.NewResponseAdapter()
	// Create base service
	baseService, err := ajexClient.NewBaseService(config, reqAdapter, respAdapter)

	if err != nil {
		return nil, fmt.Errorf("failed to create base service: %w", err)
	}

	p := &ShippingService{
		BaseService: baseService,
		maxRetries:  1,
	}

	p.authAdapter = NewAuthTokenAdapter(baseService.RestClient, psmRepo)

	return p, nil
}

func (p *ShippingService) trackByReference(ctx context.Context, req *requests.CreateForwardShipmentRequest) (*api.Response, error) {
	token, err := p.authAdapter.GetAuthToken(req.PartnerShippingMethod)
	if err != nil {
		return nil, fmt.Errorf("get auth token: %w", err)
	}

	headers := map[string][]string{
		"Accept":        {"application/json"},
		"Authorization": {token},
	}

	// Build URL with query parameters
	url := fmt.Sprintf("/order-management/api/v2/order-details?orderId=%s", req.Data.OrderPartnerOrderID)

	return p.SendRESTRequest(ctx, "GET", url, headers, nil, nil)
}

func (p *ShippingService) CreateShipment(ctx context.Context, req *requests.CreateForwardShipmentRequest) (*responses.CreateForwardShipmentResponse, error) {
	//token, err := p.authAdapter.GetAuthToken(req.PartnerShippingMethod)
	//if err != nil {
	//	return nil, fmt.Errorf("Error from Ajex: %w", err)
	//}

	// Format request
	reqBody, err := p.ReqAdapter.FormatCreateShipment(req)
	if err != nil {
		return nil, err
	}

	// Send request
	headers := map[string][]string{
		"Content-Type":  {"application/json"},
		"Authorization": {"eyJhbGciOiJIUzUxMiJ9.eyJzdWIiOiJBSlM2MzAxNTEwMDAxNDBfb21uaSIsImlzcyI6IkFqZXggVG9rZW4gQXV0aG9yaXR5Iiwic2NvcGVzIjpbIlJPTEVfQ1JFQVRFX09SREVSIl0sImFjY291bnRzIjpbXSwiaWF0IjoxNzM5MjgyOTY5LCJleHAiOjE3MzkzNjkzNjl9.nkUpsyAD3wlg7hxga9YISDcBTc_v4FCZDhs7UIoWu03k5GhCz6ZH2KRaH593rttTWiRsJFwSg87BOzuBSl1TUg"},
	}
	respBody, apiErr := p.SendRESTRequest(ctx, "POST", "/order-management/api/v2/order", headers, reqBody, nil)
	if apiErr != nil {
		return nil, apiErr
	}

	// Parse response
	formattedResponse, parseError := p.RespAdapter.ParseCreateShipment(respBody)
	if parseError != nil {
		createResponse, createErr := p.handleCreateError(ctx, parseError, req, req.PartnerShippingMethod)
		if createErr != nil {
			return nil, createErr
		}
		return createResponse, nil
	}

	return formattedResponse, nil
}

func (p *ShippingService) handleCreateError(ctx context.Context, err error, originalReq *requests.CreateForwardShipmentRequest, psm *models.PartnerShippingMethod) (*responses.CreateForwardShipmentResponse, error) {
	if p.retryCount >= p.maxRetries {
		return nil, fmt.Errorf("max retries exceeded: %v", err)
	}

	switch e := err.(type) {
	case *errors.TokenError:
		p.retryCount++
		// Invalidate token and retry
		updatedPSM, refreshErr := p.authAdapter.HandleUnauthorized(ctx, psm)
		if refreshErr != nil {
			return nil, fmt.Errorf("token refresh failed: %w", refreshErr)
		}
		originalReq.PartnerShippingMethod = updatedPSM
		return p.CreateShipment(ctx, originalReq)
	case *errors.ReferenceExistsError:
		// Get tracking details for the existing reference
		trackResp, err := p.trackByReference(ctx, originalReq)
		if err != nil {
			return nil, fmt.Errorf("track by reference: %w", err)
		}
		return p.RespAdapter.ParseCreateByReference(trackResp, originalReq)
	default:
		return nil, e
	}
}

func (p *ShippingService) CancelShipment(ctx context.Context, req *requests.CancelShipmentRequest) (*responses.CancelShipmentResponse, error) {
	// Get auth token
	//token, err := p.authAdapter.GetAuthToken(req.PartnerShippingMethod)
	//if err != nil {
	//	return nil, fmt.Errorf("get auth token: %w", err)
	//}

	// Format request
	reqBody, err := p.ReqAdapter.FormatCancelShipment(req)
	if err != nil {
		return nil, fmt.Errorf("format request: %w", err)
	}

	// Send request with auth token
	headers := map[string][]string{
		"Content-Type": {"application/json"},
		//"Authorization": {token},
	}
	respBody, err := p.SendRESTRequest(ctx, "POST", "/shipments/cancel", headers, reqBody, nil)
	if err != nil {
		return nil, err
	}

	// Parse response
	resp, err := p.RespAdapter.ParseCancelShipment(respBody)
	//if err != nil {
	// Handle errors and retry if needed
	//return p.handleCancelError(ctx, err, req, req.PartnerShippingMethod)
	//}

	return resp, nil
}

func (p *ShippingService) handleCancelError(ctx context.Context, err error, originalReq *requests.CancelShipmentRequest, psm *models.PartnerShippingMethod) (*responses.CancelShipmentResponse, error) {
	if p.retryCount >= p.maxRetries {
		return nil, fmt.Errorf("max retries exceeded: %v", err)
	}

	switch e := err.(type) {
	case *errors.TokenError:
		p.retryCount++
		// Invalidate token and retry
		// updatedPSM, refreshErr := p.authAdapter.HandleUnauthorized(ctx, psm)
		// if refreshErr != nil {
		// 	return nil, fmt.Errorf("token refresh failed: %w", refreshErr)
		// }
		return p.CancelShipment(ctx, originalReq)
	default:
		return nil, e
	}
}

func (p *ShippingService) TrackShipment(ctx context.Context, orderID string) (*responses.TrackShipmentResponse, error) {
	// Format request
	reqBody, err := p.ReqAdapter.FormatTrackShipment(&requests.TrackShipmentRequest{
		OrderID: orderID,
	})
	if err != nil {
		return nil, fmt.Errorf("format request: %w", err)
	}

	// Send request
	headers := map[string][]string{
		"Accept": {"application/json"},
	}
	respBody, err := p.SendRESTRequest(ctx, "GET", fmt.Sprintf("/shipments/%s/track", orderID), headers, reqBody, nil)
	if err != nil {
		return nil, fmt.Errorf("send request: %w", err)
	}

	// Parse response
	return p.RespAdapter.ParseTrackShipment(respBody)
}
