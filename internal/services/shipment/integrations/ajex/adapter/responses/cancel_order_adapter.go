package responses

import (
	"github.com/omniful/shipping-service/internal/services/shipment/responses"
	"github.com/omniful/shipping-service/pkg/api"
)

func (a *ResponseAdapter) ParseCancelShipment(respData *api.Response) (*responses.CancelShipmentResponse, error) {
	//var resp ajexResponse
	//if err := json.Unmarshal(data.([]byte), &resp); err != nil {
	//	return nil, fmt.Errorf("unmarshal response: %w", err)
	//}
	//
	//if !resp.Success {
	//	// Return error response with code and message
	//	return nil, &provider.ErrorResponse{
	//		Code:    resp.Error.Code,
	//		Message: resp.Error.Message,
	//		Data:    resp.Error.Data,
	//	}
	//}
	//
	//var cancelResp struct {
	//	Message string `json:"message"`
	//}
	//if err := json.Unmarshal(resp.Data, &cancelResp); err != nil {
	//	return nil, fmt.Errorf("unmarshal cancel data: %w", err)
	//}

	return &responses.CancelShipmentResponse{
		//Message: cancelResp.Message,
	}, nil
}
