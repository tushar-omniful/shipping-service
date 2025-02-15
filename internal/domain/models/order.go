package models

import (
	"time"
)

// OrderStatus represents the status of an order
type OrderStatus string

const (
	NewOrder           OrderStatus = "new_order"
	Created            OrderStatus = "created"
	ToBePicked         OrderStatus = "to_be_picked"
	Picked             OrderStatus = "picked"
	Dispatched         OrderStatus = "dispatched"
	InTransit          OrderStatus = "in_transit"
	ReceivedAtFinalHub OrderStatus = "received_at_final_hub"
	OutForDelivery     OrderStatus = "out_for_delivery"
	Reattempt          OrderStatus = "reattempt"
	UnableToDeliver    OrderStatus = "unable_to_deliver"
	Lost               OrderStatus = "lost"
	PartiallyDelivered OrderStatus = "partially_delivered"
	Delivered          OrderStatus = "delivered"
	Cancelled          OrderStatus = "cancelled"
	ReturnInProgress   OrderStatus = "return_in_progress"
	ReturnToOrigin     OrderStatus = "return_to_origin"
)

// OrderSource represents the source of an order
type OrderSource string

const (
	SourceShippingAggregator OrderSource = "shipping_aggregator"
	SourceOMS                OrderSource = "oms"
)

// Order represents the order model
type Order struct {
	ID                      uint64           `json:"id"`
	OrderPartnerOrderID     string           `json:"order_partner_order_id" gorm:"type:citext;not null"`
	ShippingPartnerOrderID  string           `json:"shipping_partner_order_id" gorm:"type:citext"`
	OrderPartnerID          uint64           `json:"order_partner_id" gorm:"not null"`
	ShippingPartnerID       uint64           `json:"shipping_partner_id" gorm:"not null"`
	Status                  OrderStatus      `json:"status" gorm:"type:citext;not null;default:'new_order'"`
	ShipmentType            string           `json:"shipment_type" gorm:"type:citext;not null"`
	ShippingPartnerStatus   string           `json:"shipping_partner_status" gorm:"type:citext"`
	PickupDetails           *PickupDetails   `json:"pickup_details" gorm:"serializer:json"`
	DropDetails             *DropDetails     `json:"drop_details" gorm:"serializer:json;not null"`
	ShipmentDetails         *ShipmentDetails `json:"shipment_details" gorm:"serializer:json;not null"`
	TaxDetails              *TaxDetails      `json:"tax_details" gorm:"serializer:json"`
	Metadata                *OrderMetadata   `json:"metadata" gorm:"serializer:json"`
	ShippingLabel           string           `json:"shipping_label" gorm:"type:citext"`
	AWBNumber               string           `json:"awb_number" gorm:"type:citext"`
	SellerID                string           `json:"seller_id" gorm:"type:citext"`
	PartnerShippingMethodID uint64           `json:"partner_shipping_method_id"`
	Source                  OrderSource      `json:"source" gorm:"type:citext"`
	CreatedAt               time.Time        `json:"created_at"`
	UpdatedAt               time.Time        `json:"updated_at"`

	// Relationships
	OrderPartner          *OrderPartner          `json:"-" gorm:"foreignKey:OrderPartnerID"`
	ShippingPartner       *ShippingPartner       `json:"-" gorm:"foreignKey:ShippingPartnerID"`
	PartnerShippingMethod *PartnerShippingMethod `json:"-" gorm:"foreignKey:PartnerShippingMethodID"`
}

// Helper methods
func (o *Order) IsInternational() bool {
	if o.IsReverseOrder() {
		return o.PickupDetails.CountryID != o.DropDetails.OmnifulCountryID
	}
	return false
}

func (o *Order) IsReverseOrder() bool {
	return o.ShipmentDetails.IsReturnOrder
}

func (o *Order) IsDispatched() bool {
	return !o.IsBeforeDispatchStatus()
}

func (o *Order) IsBeforeDispatchStatus() bool {
	return o.Status == NewOrder || o.Status == Created || o.Status == ToBePicked
}

func (o *Order) IsTerminatingStatus() bool {
	return o.Status == Cancelled || o.Status == Delivered || o.Status == ReturnToOrigin
}

// GetNumberOfRetries returns the number of retries from metadata
func (o *Order) GetNumberOfRetries() int {
	if o.Metadata == nil {
		return 0
	}
	return o.Metadata.NumberOfRetries
}

// GetTrackingURL returns the tracking URL from metadata
func (o *Order) GetTrackingURL() string {
	if o.Metadata == nil {
		return ""
	}
	return o.Metadata.TrackingURL
}

// GetAwbLabel returns the AWB label from metadata
func (o *Order) GetAwbLabel() string {
	if o.Metadata == nil {
		return ""
	}
	return o.Metadata.OmnifulAwbLabel
}

type PickupDetails struct {
	HubID              string          `json:"hub_id"`
	HubCode            string          `json:"hub_code"`
	Name               string          `json:"name"`
	Phone              string          `json:"phone"`
	CountryCallingCode string          `json:"country_calling_code"`
	Email              string          `json:"email"`
	Timezone           string          `json:"timezone"`
	CountryID          string          `json:"country_id"`
	StateID            string          `json:"state_id"`
	CityID             string          `json:"city_id"`
	City               string          `json:"city"`
	District           string          `json:"district"`
	State              string          `json:"state"`
	Country            string          `json:"country"`
	Pincode            string          `json:"pincode"`
	CountryCode        string          `json:"country_code"`
	Latitude           float64         `json:"latitude"`
	Longitude          float64         `json:"longitude"`
	Address            string          `json:"address"`
	StateCode          string          `json:"state_code"`
	SellerDetails      *SellerDetails  `json:"seller_details"`
	VillageDetails     *VillageDetails `json:"village_details"`
}

// DropDetails contains information about the drop (delivery) location.
type DropDetails struct {
	Address             string          `json:"address"`
	Phone               string          `json:"phone"`
	MobileNumber        *MobileNumber   `json:"mobile_number"`
	Country             string          `json:"country"`
	CountryCode         string          `json:"country_code"`
	CountryCurrencyCode string          `json:"country_currency_code"`
	State               string          `json:"state"`
	City                string          `json:"city"`
	CityCode            string          `json:"city_code"`
	PostalCode          string          `json:"postal_code"`
	Pincode             string          `json:"pincode"`
	Name                string          `json:"name"`
	HubCode             string          `json:"hub_code"`
	HubID               string          `json:"hub_id"`
	Email               string          `json:"email"`
	Latitude            float64         `json:"latitude"`
	Longitude           float64         `json:"longitude"`
	StateCode           string          `json:"state_code"`
	Area                string          `json:"area"`
	District            string          `json:"district"`
	OmnifulCountry      string          `json:"omniful_country"`
	OmnifulCountryID    string          `json:"omniful_country_id"`
	OmnifulState        string          `json:"omniful_state"`
	OmnifulStateID      string          `json:"omniful_state_id"`
	OmnifulCity         string          `json:"omniful_city"`
	OmnifulCityID       string          `json:"omniful_city_id"`
	VillageDetails      *VillageDetails `json:"village_details"`
	SalesChannelCity    string          `json:"sales_channel_city"`
}

// ShipmentDetails contains detailed information about the shipment.
type ShipmentDetails struct {
	CourierPartnerID     int                  `json:"courier_partner_id,omitempty"`
	TotalSKUCount        int                  `json:"total_sku_count"`
	TotalItemCount       int                  `json:"total_item_count"`
	CourierPartner       *OrderCourierPartner `json:"courier_partner,omitempty"`
	Slot                 *DeliverySlot        `json:"slot,omitempty"`
	Tags                 []Tag                `json:"tags,omitempty"`
	OrderType            string               `json:"order_type,omitempty"`
	Remarks              string               `json:"remarks,omitempty"`
	NumberOfBoxes        int                  `json:"number_of_boxes" validate:"required,gt=0"`
	Height               int                  `json:"height"`
	Length               int                  `json:"length"`
	Breadth              int                  `json:"breadth"`
	Weight               float64              `json:"weight"`
	CodValue             float64              `json:"cod_value"`
	Count                int                  `json:"count"`
	Currency             string               `json:"currency"`
	PaymentType          string               `json:"payment_type"`
	InvoiceValue         float64              `json:"invoice_value"`
	OrderValue           float64              `json:"order_value"`
	OrderNote            string               `json:"order_note"`
	OrderCreatedAt       string               `json:"order_created_at"`
	TotalPaid            float64              `json:"total_paid"`
	TotalDue             float64              `json:"total_due"`
	IsReturnOrder        bool                 `json:"is_return_order"`
	TotalDuePriceSet     *PriceSet            `json:"total_due_price_set"`
	TotalPaidPriceSet    *PriceSet            `json:"total_paid_price_set"`
	InvoiceValuePriceSet *PriceSet            `json:"invoice_value_price_set"`
	OrderValuePriceSet   *PriceSet            `json:"order_value_price_set"`
	ExchangeRate         *ExchangeRate        `json:"exchange_rate"`
	InvoiceNumber        string               `json:"invoice_number"`
	InvoiceDate          string               `json:"invoice_date"`
	PackedQuantity       int                  `json:"packed_quantity"`
	ContainsFrozenItems  bool                 `json:"contains_frozen_items"`
	Description          string               `json:"description"`
	ServiceType          string               `json:"service_type"`
	ShippingMethod       string               `json:"shipping_method"`
	Items                []Item               `json:"items"`
	PackageDetails       []PackageDetail      `json:"package_details"`
}

type MobileNumber struct {
	Number             string `json:"number,omitempty"`
	CountryCallingCode string `json:"country_calling_code,omitempty"`
	CountryCode        string `json:"country_code,omitempty"`
}

// VillageDetails contains village information
type VillageDetails struct {
	ID         int    `json:"id,omitempty"`
	NameEn     string `json:"name_en,omitempty"`
	CityID     int    `json:"city_id,omitempty"`
	CityName   string `json:"city_name,omitempty"`
	RegionID   int    `json:"region_id,omitempty"`
	RegionName string `json:"region_name,omitempty"`
}

// SellerDetails contains seller information
type SellerDetails struct {
	ID                 string         `json:"id,omitempty"`
	Name               string         `json:"name"`
	Phone              string         `json:"phone"`
	CountryCode        string         `json:"country_code,omitempty"`
	CountryCallingCode string         `json:"country_calling_code,omitempty"`
	Address            *SellerAddress `json:"address,omitempty"`
}

// SellerAddress contains seller address information
type SellerAddress struct {
	AddressLine1 string                `json:"address_line1,omitempty"`
	AddressLine2 string                `json:"address_line2,omitempty"`
	Country      *SellerAddressCountry `json:"country,omitempty"`
	State        *SellerAddressState   `json:"state,omitempty"`
	City         *SellerAddressCity    `json:"city,omitempty"`
	Pincode      string                `json:"pincode,omitempty"`
}

type SellerAddressCountry struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Code string `json:"code,omitempty"`
}

type SellerAddressState struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type SellerAddressCity struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type TaxDetails struct {
	TaxableValue         int    `json:"taxable_value,omitempty"`
	EwaybillSerialNumber string `json:"ewaybill_serial_number,omitempty"`
	PlaceOfSupply        string `json:"place_of_supply,omitempty"`
	HSNCode              string `json:"hsn_code,omitempty"`
	InvoiceReference     string `json:"invoice_reference,omitempty"`
}

// DeliverySlot contains delivery time slot information
type DeliverySlot struct {
	DeliveryDate string `json:"delivery_date,omitempty"`
	StartTime    int    `json:"start_time,omitempty"`
	EndTime      int    `json:"end_time,omitempty"`
}

// Tag represents a shipment tag
type Tag struct {
	ID    string `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Color string `json:"color,omitempty"`
}

// PriceSet contains currency information for different price types
type PriceSet struct {
	OrderCurrency Currency `json:"order_currency"`
	StoreCurrency Currency `json:"store_currency"`
}

// Currency represents monetary value with its currency code
type Currency struct {
	Amount       float64 `json:"amount,omitempty"`
	CurrencyCode string  `json:"currency_code,omitempty"`
}

// ExchangeRate represents exchange rate information
type ExchangeRate struct {
	OrderCurrency string  `json:"order_currency,omitempty"`
	StoreCurrency string  `json:"store_currency,omitempty"`
	Rate          float64 `json:"rate" validate:"required"`
}

// Item represents an individual item in the package
type Item struct {
	ProductURL      string          `json:"product_url,omitempty"`
	Price           int             `json:"price"`
	Description     string          `json:"description,omitempty"`
	Quantity        int             `json:"quantity"`
	PackedQuantity  int             `json:"packed_quantity,omitempty"`
	SKU             string          `json:"sku,omitempty"`
	SKUName         string          `json:"sku_name"`
	CountryOfOrigin string          `json:"country_of_origin,omitempty"`
	Weight          *ItemWeight     `json:"weight,omitempty"`
	Additional      *ItemAdditional `json:"additional,omitempty"`
}

// ItemWeight represents item weight information
type ItemWeight struct {
	Value float64 `json:"value,omitempty"`
	UOM   string  `json:"uom,omitempty"`
}

// ItemAdditional contains additional item information
type ItemAdditional struct {
	Height     int         `json:"height"`
	Length     int         `json:"length"`
	Breadth    int         `json:"breadth"`
	Weight     int         `json:"weight"`
	Images     []ItemImage `json:"images,omitempty"`
	ReturnDays int         `json:"return_days,omitempty"`
}

// ItemImage represents an item image
type ItemImage struct {
	URL         string `json:"url,omitempty"`
	Type        string `json:"type,omitempty"`
	Description string `json:"description,omitempty"`
}

// PackageDetail contains package-specific information
type PackageDetail struct {
	PackageID        string     `json:"package_id,omitempty"`
	Dimensions       Dimensions `json:"dimensions" validate:"required"`
	Weight           Weight     `json:"weight" validate:"required"`
	InvoiceTotal     float64    `json:"invoice_total,omitempty"`
	InvoiceSubTotal  float64    `json:"invoice_sub_total,omitempty"`
	InvoiceTotalPaid float64    `json:"invoice_total_paid,omitempty"`
	InvoiceTotalDue  float64    `json:"invoice_total_due,omitempty"`
	Items            []Item     `json:"items,omitempty"`
}

// Dimensions represents package dimensions
type Dimensions struct {
	Length  int    `json:"length" validate:"required"`
	Breadth int    `json:"breadth" validate:"required"`
	Height  int    `json:"height" validate:"required"`
	UOM     string `json:"uom,omitempty"`
}

// Weight represents package weight
type Weight struct {
	Value float64 `json:"value" validate:"required"`
	UOM   string  `json:"uom,omitempty"`
}

type OrderCourierPartner struct {
	ID   uint64 `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Tag  string `json:"tag,omitempty"`
	Logo string `json:"logo,omitempty"`
}

type OrderMetadata struct {
	OmnifulAwbLabel  string             `json:"omniful_awb_label,omitempty"`
	ShippingAwbLabel string             `json:"shipping_awb_label,omitempty"`
	TrackingURL      string             `json:"tracking_url,omitempty"`
	AwbNumber        string             `json:"awb_number,omitempty"`
	DeliveryType     string             `json:"delivery_type,omitempty"`
	TaxNumber        string             `json:"tax_number,omitempty"`
	NumberOfRetries  int                `json:"number_of_retries,omitempty"`
	EnableWhatsapp   bool               `json:"enable_whatsapp,omitempty"`
	IsFragile        bool               `json:"is_fragile,omitempty"`
	IsDangerous      bool               `json:"is_dangerous,omitempty"`
	Label            bool               `json:"label,omitempty"`
	ReturnInfo       *OrderReturnInfo   `json:"return_info,omitempty"`
	ResellerInfo     *OrderResellerInfo `json:"reseller_info,omitempty"`
}

// ReturnInfo contains return shipping information
type OrderReturnInfo struct {
	PostalCode string `json:"postal_code,omitempty"`
	Address    string `json:"address,omitempty"`
	State      string `json:"state,omitempty"`
	Phone      string `json:"phone,omitempty"`
	Name       string `json:"name,omitempty"`
	City       string `json:"city,omitempty"`
	Country    string `json:"country,omitempty"`
}

// ResellerInfo contains reseller information
type OrderResellerInfo struct {
	Name  string `json:"name,omitempty"`
	Phone string `json:"phone,omitempty"`
}
