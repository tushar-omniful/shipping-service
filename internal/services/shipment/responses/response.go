package responses

type CreateForwardShipmentResponse struct {
	TrackingNumber string                 `json:"tracking_number"`
	Label          string                 `json:"label_url"`
	Status         string                 `json:"status"`
	Metadata       map[string]interface{} `json:"metadata"`
}

type CancelShipmentResponse struct {
	Message string `json:"message"`
}

type TrackShipmentResponse struct {
	ShippingPartnerStatus string                 `json:"shipping_partner_status"`
	Status                string                 `json:"status"`
	Metadata              map[string]interface{} `json:"metadata"`
}

type WebhookResponse struct {
	Status                string                 `json:"status"`
	ShippingPartnerStatus string                 `json:"shipping_partner_status"`
	Metadata              map[string]interface{} `json:"metadata"`
}
