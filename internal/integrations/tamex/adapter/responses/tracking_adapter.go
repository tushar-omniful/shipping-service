package responses

import (
	customError "github.com/omniful/go_commons/error"
	"github.com/omniful/shipping-service/internal/services/shipment/responses"
)

type TrackResponse struct {
	CommonResponse
	AWB     string `json:"awb"`
	Status  string `json:"status"`
	Message string `json:"message"`
	History []struct {
		Status    string `json:"status"`
		Timestamp string `json:"timestamp"`
		Location  string `json:"location"`
		Message   string `json:"message"`
	} `json:"history"`
}

func (a *ResponseAdapter) ParseTrackShipment(resp TrackResponse) (response responses.TrackShipmentResponse, err customError.CustomError) {
	if resp.IsError() {
		return response, raiseBadRequest(resp.ErrorMsg())
	}

	response = responses.TrackShipmentResponse{
		ShippingPartnerStatus: resp.Status,
		Status:                resp.Status,
		Metadata: map[string]interface{}{
			"message": resp.Message,
			"history": resp.History,
		},
	}

	return
}
