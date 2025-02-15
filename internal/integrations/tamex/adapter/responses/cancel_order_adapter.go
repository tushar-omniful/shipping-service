package responses

import (
	customError "github.com/omniful/go_commons/error"
	"github.com/omniful/shipping-service/internal/services/shipment/responses"
	"time"
)

type CancelResponse struct {
	CommonResponse
}

func (a *ResponseAdapter) ParseCancelShipment(respData CancelResponse) (resp responses.CancelShipmentResponse, err customError.CustomError) {
	if respData.IsError() {
		return resp, raiseBadRequest(respData.ErrorMsg())
	}

	resp = responses.CancelShipmentResponse{
		Message: "Shipment Cancelled Successfully",
		Metadata: responses.CancelMetadata{
			StatusUpdatedAt: time.Now().String(),
		},
	}

	return
}
