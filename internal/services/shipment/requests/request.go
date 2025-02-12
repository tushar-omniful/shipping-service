package requests

import (
	"time"

	"github.com/omniful/shipping-service/internal/domain/models"
)

// CreateForwardShipmentRequest represents the request structure for creating a forward order
type CreateForwardShipmentRequest struct {
	ShippingPartner       *models.ShippingPartner       `json:"shipping_partner,omitempty"`
	PartnerShippingMethod *models.PartnerShippingMethod `json:"partner_shipping_method,omitempty"`
	ShippingPartnerTag    string                        `json:"shipping_partner_tag,omitempty"`
	TenantID              string                        `json:"tenant_id" validate:"required"`
	SellerID              string                        `json:"seller_id" validate:"required"`
	AccountID             string                        `json:"account_id" validate:"required"`
	OmnifulRuleID         string                        `json:"omniful_rule_id,omitempty"`
	SellerSalesChannelID  string                        `json:"seller_sales_channel_id" validate:"required"`
	Data                  OrderData                     `json:"data"`
}

// CancelShipmentRequest represents the request structure for canceling an order
type CancelShipmentRequest struct {
	OrderID            string `json:"order_id" validate:"required"`
	CancellationReason string `json:"cancellation_reason" validate:"required"`
	CancelledBy        string `json:"cancelled_by" validate:"required"`
	Notes              string `json:"notes,omitempty"`
}

type TrackShipmentRequest struct {
	OrderID string `json:"order_id" validate:"required"`
}

// OrderData contains the main order information
type OrderData struct {
	ShipmentType              string                 `json:"shipment_type,omitempty"`
	OrderSource               string                 `json:"order_source,omitempty"`
	OrderAlias                string                 `json:"order_alias,omitempty"`
	OrderPartnerOrderID       string                 `json:"order_partner_order_id" validate:"required"`
	SellerSalesChannelOrderID string                 `json:"seller_sales_channel_order_id" validate:"required"`
	ShippingReference         string                 `json:"shipping_reference,omitempty"`
	OrderTime                 time.Time              `json:"order_time" validate:"required"`
	PickupDetailsAttr         PickupDetailsAttrs     `json:"pickup_details_attributes"`
	DropDetailsAttr           DropDetailsAttrs       `json:"drop_details_attributes"`
	ShipmentDetailsAttr       ShipmentDetailsAttr    `json:"shipment_details_attributes"`
	TaxDetails                *models.TaxDetails     `json:"tax_details,omitempty"`
	Metadata                  *models.OrderMetadata  `json:"metadata,omitempty"`
	PreShipmentDetails        map[string]interface{} `json:"pre_shipment_details,omitempty"`
}

type PickupDetailsAttrs struct {
	HubName       string                `json:"hub_name"`
	PickupDetails *models.PickupDetails `json:"pickup_details"`
}

type DropDetailsAttrs struct {
	DropDetails *models.DropDetails `json:"drop_details"`
}

type ShipmentDetailsAttr struct {
	ShipmentDetails *models.ShipmentDetails `json:"shipment_details"`
}
