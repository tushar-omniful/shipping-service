package models

import (
	"encoding/json"
	"time"
)

type CityMapping struct {
	ID                 uint64           `json:"id"`
	OmnifulCityID      string          `json:"omniful_city_id" gorm:"type:citext"`
	Name               string          `json:"name" gorm:"type:citext"`
	Code               string          `json:"code" gorm:"type:citext"`
	ShippingPartnerID  int64           `json:"shipping_partner_id" gorm:"not null"`
	Metadata           json.RawMessage `json:"metadata" gorm:"serializer:json"`
	GooglePlaceID      string          `json:"google_place_id"`
	GooglePlaceName    string          `json:"google_place_name"`
	OmnifulCountryID   string          `json:"omniful_country_id" gorm:"type:citext"`
	OmnifulCityName    string          `json:"omniful_city_name" gorm:"type:citext"`
	OmnifulCountryName string          `json:"omniful_country_name" gorm:"type:citext"`
	CreatedAt          time.Time       `json:"created_at"`
	UpdatedAt          time.Time       `json:"updated_at"`
}

func (CityMapping) TableName() string {
	return "city_mappings"
}
