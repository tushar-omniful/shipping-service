package requests

import (
	"github.com/omniful/shipping-service/internal/services/shipment/requests"
)

func (a *RequestAdapter) FormatCancelShipment(req *requests.CancelShipmentRequest) (interface{}, error) {
	requestData := map[string]string{
		"waybillNumber": req.OrderID,
	}

	return requestData, nil
}
