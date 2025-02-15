package responses

import (
	"fmt"
	customError "github.com/omniful/go_commons/error"
	"time"

	"github.com/omniful/shipping-service/internal/services/shipment/requests"
	"github.com/omniful/shipping-service/internal/services/shipment/responses"
)

type TrackByReferenceResponse struct {
	CommonResponse
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

func (a *ResponseAdapter) ParseCreateByReference(respData TrackByReferenceResponse, req *requests.CreateForwardShipmentRequest) (resp responses.CreateForwardShipmentResponse, err customError.CustomError) {
	if respData.IsError() {
		return resp, raiseBadRequest(respData.ErrorMsg())
	}

	if !isValidReference(respData, req) {
		return resp, raiseBadRequest(respData.ResponseMsg)
	}

	resp = responses.CreateForwardShipmentResponse{
		AwbNumber: respData.WaybillNumber,
		Label:     respData.WaybillFileURL,
		Status:    "created",
		Metadata: responses.CreateMetadata{
			ShippingAwbLabel:   respData.WaybillFileURL,
			TrackingUrl:        getTrackingURL(respData.WaybillNumber),
			CreatedByReference: true,
		},
	}
	return
}

func isValidReference(details TrackByReferenceResponse, req *requests.CreateForwardShipmentRequest) bool {
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

func matchesShipmentDetails(details TrackByReferenceResponse, req *requests.CreateForwardShipmentRequest) bool {
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

func matchesCustomerDetails(details TrackByReferenceResponse, req *requests.CreateForwardShipmentRequest) bool {
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
