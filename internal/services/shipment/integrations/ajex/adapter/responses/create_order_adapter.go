package responses

import (
	"encoding/json"
	"fmt"
	"github.com/omniful/shipping-service/pkg/api"
	"strings"

	"github.com/omniful/shipping-service/internal/services/shipment/integrations/ajex/errors"
	"github.com/omniful/shipping-service/internal/services/shipment/responses"
)

// Common response structure
type CommonResponse struct {
	IsSuccess  bool        `json:"is_success"`
	StatusCode int         `json:"status_code"`
	Data       interface{} `json:"data"`
}

type AjexResponse struct {
	Success      bool            `json:"success"`
	ResponseCode string          `json:"responseCode"`
	ResponseMsg  string          `json:"responseMessage"`
	Data         json.RawMessage `json:"data"`
}

func (a *ResponseAdapter) ParseCreateShipment(respData *api.Response) (*responses.CreateForwardShipmentResponse, error) {
	// Handle non-successful responses first
	if !respData.IsSuccess {
		// Handle specific error cases
		return nil, handleError(respData)
	}

	// Handle successful response
	respBytes, ok := respData.Body.([]byte)
	if !ok {
		return nil, fmt.Errorf("invalid success response type: %v", respData.Body)
	}

	var ajexResp AjexResponse
	if err := json.Unmarshal(respBytes, &ajexResp); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}

	// Verify response codes even for successful HTTP response
	if !ajexResp.Success || ajexResp.ResponseCode != "100" || ajexResp.ResponseMsg != "Success" {
		return nil, handleError(respData)
	}

	// Parse shipment data
	var shipmentData struct {
		TrackingNumber string `json:"tracking_number"`
		LabelURL       string `json:"label_url"`
		Status         string `json:"status"`
	}
	if err := json.Unmarshal(ajexResp.Data, &shipmentData); err != nil {
		return nil, fmt.Errorf("unmarshal shipment data: %w", err)
	}

	return &responses.CreateForwardShipmentResponse{
		TrackingNumber: shipmentData.TrackingNumber,
		Label:          shipmentData.LabelURL,
		Status:         shipmentData.Status,
	}, nil
}

func getStatusCode(msg string) int {
	if strings.Contains(strings.ToLower(msg), errors.Unauthorized) {
		return 401
	}
	if strings.Contains(strings.ToLower(msg), errors.NotFound) {
		return 404
	}
	if strings.Contains(strings.ToLower(msg), errors.ServiceTemporarilyUnavailable) {
		return 503
	}
	return 200
}

func handleError(response *api.Response) error {
	// Try to get error message from response body
	bodyStr, ok := response.Body.(string)
	if ok && strings.Contains(strings.ToLower(bodyStr), errors.ServiceTemporarilyUnavailable) {
		return fmt.Errorf("service temporarily unavailable. Please try again later")
	}

	// Handle 401 status code
	if response.StatusCode == 401 {
		return &errors.TokenError{
			Message: "Unauthorized",
			Code:    "UNAUTHORIZED",
		}
	}

	// Try to parse error response
	respBytes, ok := response.Body.([]byte)
	if !ok {
		return fmt.Errorf("invalid error response type: %v", response.Body)
	}

	var ajexResp AjexResponse
	if err := json.Unmarshal(respBytes, &ajexResp); err != nil {
		return fmt.Errorf("failed to parse error response: %w", err)
	}

	errMsg := ajexResp.ResponseMsg
	if errMsg == "" {
		errMsg = "unknown error"
	}

	if strings.Contains(strings.ToLower(errMsg), errors.OrderIdExists) {
		return &errors.ReferenceExistsError{
			Message: errMsg,
			Code:    "ORDER_EXISTS",
		}
	}
	if strings.Contains(strings.ToLower(errMsg), errors.Unauthorized) {
		return &errors.TokenError{
			Message: errMsg,
			Code:    "UNAUTHORIZED",
		}
	}
	return fmt.Errorf("AJEX API error: %s", errMsg)
}
