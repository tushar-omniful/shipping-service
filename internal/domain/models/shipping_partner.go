package models

import "time"

type ShippingPartner struct {
	ID                       uint64                 `json:"id"`
	Name                     string                 `json:"name" gorm:"type:citext"`
	Tag                      string                 `json:"tag" gorm:"type:citext"`
	Details                  map[string]interface{} `json:"details" gorm:"serializer:json"`
	Status                   string                 `json:"status" gorm:"type:citext"`
	Category                 string                 `json:"category" gorm:"type:citext"`
	Logo                     string                 `json:"logo" gorm:"type:citext"`
	WebhookEnabled           string                 `json:"webhook_enabled" gorm:"type:citext"`
	AggregatorTag            string                 `json:"aggregator_tag" gorm:"type:citext"`
	IntegrationStructure     map[string]interface{} `json:"integration_structure" gorm:"serializer:json"`
	TenantIDs                interface{}            `json:"tenant_ids" gorm:"type:citext[]"`
	CityMappingAggregatorTag string                 `json:"city_mapping_aggregator_tag" gorm:"type:citext"`
	CreatedAt                time.Time              `json:"created_at"`
	UpdatedAt                time.Time              `json:"updated_at"`
}

func (ShippingPartner) TableName() string {
	return "shipping_partners"
}
