package models

import (
	"encoding/json"
	"time"
)

type ActionLog struct {
	ID             uint64          `json:"id"`
	ActionableType string          `json:"actionable_type" gorm:"type:citext;not null"`
	ActionableID   int64           `json:"actionable_id" gorm:"not null"`
	OmnifulUserID  string          `json:"omniful_user_id" gorm:"type:citext"`
	OmnifulUser    json.RawMessage `json:"omniful_user" gorm:"serializer:json"`
	Action         string          `json:"action" gorm:"type:citext"`
	Metadata       json.RawMessage `json:"metadata" gorm:"serializer:json"`
	CreatedAt      time.Time       `json:"created_at"`
}

func (ActionLog) TableName() string {
	return "action_logs"
}
