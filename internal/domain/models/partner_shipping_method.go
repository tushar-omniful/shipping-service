package models

import (
	"fmt"
	"time"

	"gorm.io/gorm"

	"github.com/lib/pq"
)

// Status represents the status of a partner shipping method
type Status string

const (
	StatusActive   Status = "active"
	StatusInactive Status = "inactive"
)

// AWBReferenceConfig defines the types of AWB references
var AWBReferenceConfig = map[string]string{
	"order_id":           "order_id",
	"reference_order_id": "reference_order_id",
	"combined_reference": "combined_reference",
	"system_default":     "system_default",
}

type ShippingCostRules struct {
	PriorityOrder  int    `json:"priority_order"`
	RateCardRuleID string `json:"rate_card_rule_id"`
}

// PartnerShippingMethod represents the shipping method configuration for a partner
type PartnerShippingMethod struct {
	ID                   uint64                 `json:"id"`
	OrderPartnerID       uint64                 `json:"order_partner_id" gorm:"not null"`
	ShippingPartnerID    uint64                 `json:"shipping_partner_id" gorm:"not null"`
	Credentials          map[string]interface{} `json:"credentials" gorm:"serializer:json"`
	Details              map[string]interface{} `json:"details" gorm:"serializer:json"`
	Status               Status                 `json:"status" gorm:"type:citext"`
	WebhookToken         string                 `json:"webhook_token" gorm:"type:citext;not null"`
	CredentialKey        string                 `json:"credential_key" gorm:"type:citext"`
	SellerIDs            pq.StringArray         `json:"seller_ids" gorm:"type:citext[]"`
	IsAllSellerMapped    bool                   `json:"is_all_seller_mapped"`
	AccountName          string                 `json:"account_name" gorm:"type:citext"`
	ForwardShipmentOrder int                    `json:"forward_shipment_order"`
	ReverseShipmentOrder int                    `json:"reverse_shipment_order"`
	AWBConfig            map[string]interface{} `json:"awb_config" gorm:"serializer:json"`
	WebhookConfig        map[string]interface{} `json:"webhook_config" gorm:"serializer:json"`
	SLAConfig            map[string]interface{} `json:"sla_config" gorm:"serializer:json"`
	ShippingCostRules    interface{}            `json:"shipping_cost_rules" gorm:"type:jsonb[]"`
	DefaultShippingCost  float64                `json:"default_shipping_cost" gorm:"default:0.0"`
	ShippingCost         map[string]interface{} `json:"shipping_cost" gorm:"serializer:json"`
	SortingCode          string                 `json:"sorting_code"`
	Metadata             map[string]interface{} `json:"metadata" gorm:"serializer:json;default:'{}'"`
	Configurations       map[string]interface{} `json:"configurations" gorm:"serializer:json;default:'{}'"`
	CreatedAt            time.Time              `json:"created_at"`
	UpdatedAt            time.Time              `json:"updated_at"`

	// Relationships
	OrderPartner    *OrderPartner    `json:"-" gorm:"foreignKey:OrderPartnerID"`
	ShippingPartner *ShippingPartner `json:"-" gorm:"foreignKey:ShippingPartnerID"`
	Orders          []Order          `json:"-" gorm:"foreignKey:PartnerShippingMethodID"`
}

func (psm *PartnerShippingMethod) GetSortingCodeRules() []map[string]interface{} {
	if psm.Metadata == nil {
		return nil
	}
	if rules, ok := psm.Metadata["sorting_code_rules"].([]map[string]interface{}); ok {
		return rules
	}
	return nil
}

func (psm *PartnerShippingMethod) AddSortingCodeRule(ruleID string) {
	if psm.Metadata == nil {
		psm.Metadata = make(map[string]interface{})
	}

	rules := psm.GetSortingCodeRules()
	if rules == nil {
		rules = make([]map[string]interface{}, 0)
	}

	// Check if rule already exists
	for _, rule := range rules {
		if rule["rule_id"] == ruleID {
			return
		}
	}

	rules = append(rules, map[string]interface{}{"rule_id": ruleID})
	psm.Metadata["sorting_code_rules"] = rules
}

func (psm *PartnerShippingMethod) RemoveSortingCodeRule(ruleID string) {
	rules := psm.GetSortingCodeRules()
	if rules == nil {
		return
	}

	newRules := make([]map[string]interface{}, 0)
	for _, rule := range rules {
		if rule["rule_id"] != ruleID {
			newRules = append(newRules, rule)
		}
	}

	psm.Metadata["sorting_code_rules"] = newRules
}

func (psm *PartnerShippingMethod) GetUpdateOrderWebhookURL(publicURL string) string {
	return fmt.Sprintf("%s/v1/webhooks/update_order?webhook_token=%s", publicURL, psm.WebhookToken)
}

// BeforeCreate hook to initialize default values
func (psm *PartnerShippingMethod) BeforeCreate(tx *gorm.DB) error {
	psm.initializeDefaults()
	return nil
}

func (psm *PartnerShippingMethod) initializeDefaults() {
	if psm.Credentials == nil {
		psm.Credentials = make(map[string]interface{})
	}
	if psm.Details == nil {
		psm.Details = make(map[string]interface{})
	}
	if psm.Status == "" {
		psm.Status = StatusActive
	}
	if psm.SellerIDs == nil {
		psm.SellerIDs = make(pq.StringArray, 0)
	}
	if psm.SLAConfig == nil {
		psm.SLAConfig = make(map[string]interface{})
	}
	if psm.Metadata == nil {
		psm.Metadata = make(map[string]interface{})
	}
	if psm.ShippingCostRules == nil {
		psm.ShippingCostRules = make([]interface{}, 0)
	}

	psm.initializeAWBConfig()
	psm.initializeWebhookConfig()
}

func (psm *PartnerShippingMethod) initializeAWBConfig() {
	if psm.AWBConfig != nil {
		return
	}

	psm.AWBConfig = map[string]interface{}{
		"seller": map[string]interface{}{
			"phone_enabled": true,
			"name_enabled":  true,
			"email_enabled": true,
		},
		"hub": map[string]interface{}{
			"phone_enabled": false,
			"name_enabled":  false,
			"email_enabled": false,
		},
		"custom": map[string]interface{}{
			"phone_enabled": false,
			"name_enabled":  false,
			"email_enabled": false,
			"phone":         "",
			"name":          "",
			"email":         "",
		},
		"reference": map[string]interface{}{
			"type": AWBReferenceConfig["system_default"],
		},
	}
}

func (psm *PartnerShippingMethod) initializeWebhookConfig() {
	if psm.WebhookConfig != nil {
		return
	}

	psm.WebhookConfig = map[string]interface{}{
		"is_webhook_message_ignored": false,
	}
}

// HasSeller checks if a given seller ID is mapped to this shipping method
func (psm *PartnerShippingMethod) HasSeller(sellerID string) bool {
	// If all sellers are mapped, return true
	if psm.IsAllSellerMapped {
		return true
	}

	// Check if seller exists in SellerIDs array
	for _, id := range psm.SellerIDs {
		if id == sellerID {
			return true
		}
	}
	return false
}
