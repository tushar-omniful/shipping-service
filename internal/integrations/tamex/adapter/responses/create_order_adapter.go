package responses

import (
	customError "github.com/omniful/go_commons/error"
	"github.com/omniful/shipping-service/internal/services/shipment/responses"
)

type CreateResponse struct {
	CommonResponse
	TmxAWB string `json:"tmxAWB"`
}

func (a *ResponseAdapter) ParseCreateShipment(resp CreateResponse) (response responses.CreateForwardShipmentResponse, err customError.CustomError) {
	if resp.Code != 0 { // 0 means success for Tamex
		return response, raiseBadRequest(resp.Data)
	}

	response = responses.CreateForwardShipmentResponse{
		AwbNumber: resp.TmxAWB,
		Status:    "created",
		Metadata: responses.CreateMetadata{
			TrackingUrl: resp.Data,
		},
	}

	return
}
