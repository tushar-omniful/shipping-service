package requests

import (
	"time"

	"github.com/omniful/shipping-service/internal/domain/models"
	"github.com/omniful/shipping-service/internal/services/shipment/requests"
)

// CreateForwardOrder represents the request structure for creating a forward order
type CreateForwardOrder struct {
	ShippingPartnerTag   string    `json:"shipping_partner_tag,omitempty"`
	TenantID             string    `json:"tenant_id" validate:"required"`
	SellerID             string    `json:"seller_id" validate:"required"`
	AccountID            string    `json:"account_id" validate:"required"`
	OmnifulRuleID        string    `json:"omniful_rule_id,omitempty"`
	SellerSalesChannelID string    `json:"seller_sales_channel_id" validate:"required"`
	Data                 OrderData `json:"data"`
}

// CancelOrderRequest represents the request structure for canceling an order
type CancelOrderRequest struct {
	OrderID            string `json:"order_id" validate:"required"`
	CancellationReason string `json:"cancellation_reason" validate:"required"`
	CancelledBy        string `json:"cancelled_by" validate:"required"`
	Notes              string `json:"notes,omitempty"`
}

// OrderData contains the main order information
type OrderData struct {
	ShipmentType              string                  `json:"shipment_type,omitempty"`
	OrderSource               string                  `json:"order_source,omitempty"`
	OrderAlias                string                  `json:"order_alias,omitempty"`
	OrderPartnerOrderID       string                  `json:"order_partner_order_id" validate:"required"`
	SellerSalesChannelOrderID string                  `json:"seller_sales_channel_order_id" validate:"required"`
	ShippingReference         string                  `json:"shipping_reference,omitempty"`
	OrderTime                 time.Time               `json:"order_time" validate:"required"`
	PickupDetails             *models.PickupDetails   `json:"pickup_details"`
	DropDetails               *models.DropDetails     `json:"drop_details"`
	TaxDetails                *models.TaxDetails      `json:"tax_details,omitempty"`
	Metadata                  *models.OrderMetadata   `json:"metadata,omitempty"`
	ShipmentDetails           *models.ShipmentDetails `json:"shipment_details"`
	PreShipmentDetails        map[string]interface{}  `json:"pre_shipment_details,omitempty"`
}

func (cfo *CreateForwardOrder) TransformToShipmentRequest(sp *models.ShippingPartner, psm *models.PartnerShippingMethod) (*requests.CreateForwardShipmentRequest, error) {
	return &requests.CreateForwardShipmentRequest{
		ShippingPartner:       sp,
		PartnerShippingMethod: psm,
		ShippingPartnerTag:    cfo.ShippingPartnerTag,
		TenantID:              cfo.TenantID,
		SellerID:              cfo.SellerID,
		AccountID:             cfo.AccountID,
		OmnifulRuleID:         cfo.OmnifulRuleID,
		SellerSalesChannelID:  cfo.SellerSalesChannelID,
		Data: requests.OrderData{
			ShipmentType:              cfo.Data.ShipmentType,
			OrderSource:               cfo.Data.OrderSource,
			OrderAlias:                cfo.Data.OrderAlias,
			OrderPartnerOrderID:       cfo.Data.OrderPartnerOrderID,
			SellerSalesChannelOrderID: cfo.Data.SellerSalesChannelOrderID,
			ShippingReference:         cfo.Data.ShippingReference,
			OrderTime:                 cfo.Data.OrderTime,
			PickupDetailsAttr: requests.PickupDetailsAttrs{
				HubName:       cfo.Data.PickupDetails.Name,
				PickupDetails: cfo.Data.PickupDetails,
			},
			DropDetailsAttr: requests.DropDetailsAttrs{
				DropDetails: cfo.Data.DropDetails,
			},
			ShipmentDetailsAttr: requests.ShipmentDetailsAttr{
				ShipmentDetails: cfo.Data.ShipmentDetails,
			},
			TaxDetails:         cfo.Data.TaxDetails,
			Metadata:           cfo.Data.Metadata,
			PreShipmentDetails: cfo.Data.PreShipmentDetails,
		},
	}, nil
}
