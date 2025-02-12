package responses

import (
	"fmt"
	"github.com/omniful/shipping-service/pkg/api"
	"time"

	"github.com/omniful/shipping-service/internal/services/shipment/requests"
	"github.com/omniful/shipping-service/internal/services/shipment/responses"
)

type orderDetailsResponse struct {
	OrderID         string       `json:"orderId"`
	WaybillNumber   string       `json:"waybillNumber"`
	WaybillFileURL  string       `json:"waybillFileUrl"`
	OrderTime       string       `json:"orderTime"`
	CustomerAccount string       `json:"customerAccount"`
	CodAmount       float64      `json:"codAmount"`
	TotalDeclared   float64      `json:"totalDeclaredValue"`
	ReceiverInfo    receiverInfo `json:"receiverInfo"`
}

type receiverInfo struct {
	Name            string `json:"name"`
	Phone           string `json:"phone"`
	City            string `json:"city"`
	DetailedAddress string `json:"detailedAddress"`
	Country         string `json:"country"`
}

func (a *ResponseAdapter) ParseCreateByReference(respData *api.Response, req *requests.CreateForwardShipmentRequest) (*responses.CreateForwardShipmentResponse, error) {
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
	var details orderDetailsResponse
	//if err := json.Unmarshal(resp.Data, &details); err != nil {
	//	return nil, fmt.Errorf("unmarshal order details: %w", err)
	//}
	//
	//// Validate the reference matches
	//if !isValidReference(details, req) {
	//	return nil, fmt.Errorf("reference validation failed")
	//}

	return &responses.CreateForwardShipmentResponse{
		TrackingNumber: details.WaybillNumber,
		Label:          details.WaybillFileURL,
		Status:         "created",
		Metadata: map[string]interface{}{
			"shipping_awb_label":   details.WaybillFileURL,
			"tracking_url":         getTrackingURL(details.WaybillNumber),
			"created_by_reference": true,
		},
	}, nil
}

func isValidReference(details orderDetailsResponse, req *requests.CreateForwardShipmentRequest) bool {
	// Validate customer details match
	if !matchesCustomerDetails(details, req) {
		return false
	}

	// Validate shipment details match
	if !matchesShipmentDetails(details, req) {
		return false
	}

	// Validate creation time
	orderTime, err := time.Parse(time.RFC3339, details.OrderTime)
	if err != nil {
		return false
	}

	// Order should be created within last 3 days
	return time.Since(orderTime) <= 72*time.Hour
}

func matchesShipmentDetails(details orderDetailsResponse, req *requests.CreateForwardShipmentRequest) bool {
	shipmentDetails := req.Data.ShipmentDetailsAttr.ShipmentDetails

	// Validate declared value
	if details.TotalDeclared != shipmentDetails.InvoiceValue {
		return false
	}

	// Validate COD amount if applicable
	if shipmentDetails.TotalDue > 0 && details.CodAmount != shipmentDetails.TotalDue {
		return false
	}

	return true
}

func matchesCustomerDetails(details orderDetailsResponse, req *requests.CreateForwardShipmentRequest) bool {
	dropDetails := req.Data.DropDetailsAttr.DropDetails
	receiverInfo := details.ReceiverInfo

	// Basic validation of receiver details
	if receiverInfo.Name != dropDetails.Name {
		return false
	}
	if receiverInfo.Phone != dropDetails.Phone {
		return false
	}
	if !contains(receiverInfo.DetailedAddress, dropDetails.Address) {
		return false
	}
	if !contains(receiverInfo.City, dropDetails.City) {
		return false
	}
	if !contains(receiverInfo.Country, dropDetails.Country) {
		return false
	}

	return true
}

func contains(str, substr string) bool {
	return len(str) > 0 && len(substr) > 0 && str == substr
}

func getTrackingURL(trackingNumber string) string {
	// TODO: Make this configurable based on environment
	return fmt.Sprintf("https://aj-ex.com/tracking?tracking_number=%s", trackingNumber)
}
