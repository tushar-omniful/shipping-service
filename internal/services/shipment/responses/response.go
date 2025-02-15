package responses

type CreateForwardShipmentResponse struct {
	AwbNumber             string         `json:"awb_umber"`
	Label                 string         `json:"label_url"`
	ShippingPartnerStatus string         `json:"shipping_partner_status"`
	Status                string         `json:"status"`
	Metadata              CreateMetadata `json:"metadata"`
}

type CancelShipmentResponse struct {
	Message  string         `json:"message"`
	Metadata CancelMetadata `json:"metadata"`
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

type CreateMetadata struct {
	OmnifulAwbLabel    string `json:"awb_label"`
	ShippingAwbLabel   string `json:"shipping_awb_label"`
	TrackingUrl        string `json:"tracking_url"`
	CreatedByReference bool   `json:"created_by_reference,omitempty"`
}

type CancelMetadata struct {
	StatusUpdatedAt string `json:"status_updated_at"`
}
