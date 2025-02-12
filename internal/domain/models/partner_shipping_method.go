package models

import (
	"time"

	"github.com/lib/pq"
)

type PartnerShippingMethod struct {
	ID                   uint64                 `json:"id"`
	OrderPartnerID       uint64                 `json:"order_partner_id" gorm:"not null"`
	ShippingPartnerID    uint64                 `json:"shipping_partner_id" gorm:"not null"`
	Credentials          map[string]interface{} `json:"credentials" gorm:"serializer:json"`
	Details              map[string]interface{} `json:"details" gorm:"serializer:json"`
	Status               string                 `json:"status" gorm:"type:citext"`
	WebhookToken         string                 `json:"webhook_token" gorm:"type:citext;not null"`
	CredentialKey        string                 `json:"credential_key" gorm:"type:citext"`
	SellerIDs            pq.StringArray         `json:"seller_ids" gorm:"type:citext[]"`
	IsAllSellerMapped    bool                   `json:"is_all_seller_mapped"`
	AccountName          string                 `json:"account_name" gorm:"type:citext"`
	ForwardShipmentOrder int                    `json:"forward_shipment_order"`
	ReverseShipmentOrder int                    `json:"reverse_shipment_order"`
	AWBConfig            map[string]interface{} `json:"awb_config" gorm:"serializer:json"`
	WebhookConfig        map[string]interface{} `json:"webhook_config" gorm:"serializer:json"`
	SLAConfig            interface{}            `json:"sla_config" gorm:"serializer:json"`
	ShippingCostRules    interface{}            `json:"shipping_cost_rules" gorm:"type:jsonb[]"`
	DefaultShippingCost  float64                `json:"default_shipping_cost" gorm:"default:0.0"`
	ShippingCost         map[string]interface{} `json:"shipping_cost" gorm:"serializer:json"`
	SortingCode          string                 `json:"sorting_code"`
	Metadata             map[string]interface{} `json:"metadata" gorm:"serializer:json;default:'{}'"`
	Configurations       map[string]interface{} `json:"configurations" gorm:"serializer:json;default:'{}'"`
	CreatedAt            time.Time              `json:"created_at"`
	UpdatedAt            time.Time              `json:"updated_at"`
}

func (PartnerShippingMethod) TableName() string {
	return "partner_shipping_methods"
}
