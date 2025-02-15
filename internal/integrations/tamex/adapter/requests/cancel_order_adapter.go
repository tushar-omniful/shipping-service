package requests

import (
	customError "github.com/omniful/go_commons/error"
	"github.com/omniful/shipping-service/internal/services/shipment/requests"
)

func (a *RequestAdapter) FormatCancelShipment(req *requests.CancelShipmentRequest) (requestData map[string]interface{}, err customError.CustomError) {
	credentials := req.PartnerShippingMethod.Credentials
	apiKey, ok := credentials["api_key"].(string)
	if !ok {
		return nil, customError.NewCustomError("invalid_credentials", "Missing API key in credentials")
	}

	requestData = map[string]interface{}{
		"apikey":   apiKey,
		"pack_awb": req.Order.AWBNumber,
	}

	return
}
