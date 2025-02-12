package models

import (
	"encoding/json"
	"time"
)

type CourierPartner struct {
	ID        uint64           `json:"id" gorm:"primaryKey"`
	Name      string          `json:"name" gorm:"column:name;type:citext"`
	Tag       string          `json:"tag" gorm:"column:tag;type:citext"`
	Details   json.RawMessage `json:"details" gorm:"serializer:json"`
	Status    string          `json:"status" gorm:"column:status;type:citext"`
	Logo      string          `json:"logo" gorm:"column:logo;type:citext"`
	Synonyms  []string        `json:"synonyms" gorm:"column:synonyms;type:citext[]"`
	TenantIDs []string        `json:"tenant_ids" gorm:"column:tenant_ids;type:citext[]"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
}

func (CourierPartner) TableName() string {
	return "courier_partners"
}
