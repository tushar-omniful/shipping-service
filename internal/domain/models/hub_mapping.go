package models

import (
	"encoding/json"
	"time"
)

type HubMapping struct {
	ID                      int64           `json:"id"`
	ShippingHubID           string          `json:"shipping_hub_id" gorm:"type:citext;not null"`
	OmnifulHubID            string          `json:"omniful_hub_id" gorm:"type:citext;not null"`
	TenantID                string          `json:"tenant_id" gorm:"type:citext;not null"`
	ShippingPartnerID       int64           `json:"shipping_partner_id" gorm:"not null"`
	Details                 json.RawMessage `json:"details" gorm:"serializer:json"`
	Status                  string          `json:"status" gorm:"type:citext;not null"`
	PartnerShippingMethodID int64           `json:"partner_shipping_method_id"`
	CreatedAt               time.Time       `json:"created_at"`
	UpdatedAt               time.Time       `json:"updated_at"`
}

func (HubMapping) TableName() string {
	return "hub_mappings"
}
