package models

import (
	"time"

	"gorm.io/gorm"
)

type TenantCityMapping struct {
	CityMappingID      uint64    `json:"city_mapping_id"`
	TenantID           string    `json:"tenant_id" gorm:"type:citext"`
	OmnifulCityID      string    `json:"omniful_city_id" gorm:"type:citext"`
	Status             string    `json:"status" gorm:"type:citext;default:pending"`
	TenantStatus       string    `json:"tenant_status" gorm:"type:citext;default:tenant_approved"`
	CreatedBy          string    `json:"created_by" gorm:"type:citext"`
	UpdatedBy          string    `json:"updated_by" gorm:"type:citext"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
	ShippingPartnerID  int64     `json:"shipping_partner_id"`
	OmnifulCityName    string    `json:"omniful_city_name" gorm:"type:citext;index"`
	OmnifulCountryName string    `json:"omniful_country_name" gorm:"type:citext"`
	OmnifulCountryID   string    `json:"omniful_country_id" gorm:"type:citext"`
	CreatedByUserID    string    `json:"created_by_user_id" gorm:"type:citext"`
	UpdatedByUserID    string    `json:"updated_by_user_id" gorm:"type:citext"`
	CreatedByUserName  string    `json:"created_by_user_name" gorm:"type:citext"`
	UpdatedByUserName  string    `json:"updated_by_user_name" gorm:"type:citext"`

	CityMapping     *CityMapping     `json:"city_mapping,omitempty" gorm:"foreignKey:CityMappingID"`
	ShippingPartner *ShippingPartner `json:"shipping_partner,omitempty" gorm:"foreignKey:ShippingPartnerID"`
}

// Constants for status values
const (
	StatusPending  = "pending"
	StatusApproved = "approved"
	StatusRejected = "rejected"

	TenantStatusApproved = "tenant_approved"
	TenantStatusRejected = "tenant_rejected"
)

// BeforeCreate hook to set default values
func (tcm *TenantCityMapping) BeforeCreate(tx *gorm.DB) error {
	if tcm.Status == "" {
		tcm.Status = StatusPending
	}
	if tcm.TenantStatus == "" {
		tcm.TenantStatus = TenantStatusApproved
	}
	return nil
}

func (TenantCityMapping) TableName() string {
	return "tenant_city_mappings"
}
