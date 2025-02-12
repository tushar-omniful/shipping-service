package requests

import (
	"github.com/omniful/shipping-service/internal/services/shipment/requests"
)

func (a *RequestAdapter) FormatTrackShipment(req *requests.TrackShipmentRequest) (interface{}, error) {
	return nil, nil // GET request doesn't need body
}
