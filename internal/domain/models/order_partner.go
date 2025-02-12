package models

import (
	"time"
)

type OrderPartner struct {
	ID        uint64                 `json:"id" gorm:"primaryKey"`
	Name      string                 `json:"name" gorm:"type:citext;not null"`
	TenantID  string                 `json:"tenant_id" gorm:"type:citext;not null"`
	Details   map[string]interface{} `json:"details" gorm:"serializer:json"`
	Status    string                 `json:"status" gorm:"type:citext"`
	Metadata  map[string]interface{} `json:"metadata" gorm:"serializer:json"`
	SellerID  string                 `json:"seller_id" gorm:"type:citext"`
	CreatedAt time.Time              `json:"created_at"`
	UpdatedAt time.Time              `json:"updated_at"`
}

func (OrderPartner) TableName() string {
	return "order_partners"
}
