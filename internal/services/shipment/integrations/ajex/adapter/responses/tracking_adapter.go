package responses

import (
	"github.com/omniful/shipping-service/internal/services/shipment/responses"
	"github.com/omniful/shipping-service/pkg/api"
)

func (a *ResponseAdapter) ParseTrackShipment(respData *api.Response) (*responses.TrackShipmentResponse, error) {
	return nil, nil // GET request doesn't need body
}
