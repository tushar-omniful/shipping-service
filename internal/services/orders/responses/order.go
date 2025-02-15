package responses

import (
	"github.com/omniful/shipping-service/internal/domain/models"
	"time"
)

type GetOrderResponse struct {
	ID                    uint64    `json:"id"`
	OrderPartnerOrderID   string    `json:"order_partner_order_id"`
	Status                string    `json:"status"`
	ShippingPartnerStatus string    `json:"shipping_partner_status"`
	AWBNumber             string    `json:"awb_number"`
	ShippingLabel         string    `json:"shipping_label"`
	TrackingURL           string    `json:"tracking_url"`
	CreatedAt             time.Time `json:"created_at"`
	UpdatedAt             time.Time `json:"updated_at"`
}

func ConvertOrderModelToCreateResponse(order *models.Order) *GetOrderResponse {
	return &GetOrderResponse{
		ID:                    order.ID,
		OrderPartnerOrderID:   order.OrderPartnerOrderID,
		Status:                string(order.Status),
		ShippingPartnerStatus: order.ShippingPartnerStatus,
		AWBNumber:             order.AWBNumber,
		ShippingLabel:         order.ShippingLabel,
		TrackingURL:           order.GetTrackingURL(),
		CreatedAt:             order.CreatedAt,
		UpdatedAt:             order.UpdatedAt,
	}
}
