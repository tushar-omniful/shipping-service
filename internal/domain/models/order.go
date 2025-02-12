package models

import (
	"time"
)

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

type OrderSource string

const (
	SourceShippingAggregator OrderSource = "shipping_aggregator"
	SourceOMS                OrderSource = "oms"
)

type Order struct {
	ID                      uint64           `json:"id"`
	OrderPartnerOrderID     string           `json:"order_partner_order_id" gorm:"type:citext;not null"`
	ShippingPartnerOrderID  string           `json:"shipping_partner_order_id" gorm:"type:citext"`
	OrderPartnerID          int64            `json:"order_partner_id" gorm:"not null"`
	ShippingPartnerID       int64            `json:"shipping_partner_id" gorm:"not null"`
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
	PartnerShippingMethodID int64            `json:"partner_shipping_method_id"`
	Source                  OrderSource      `json:"source" gorm:"type:citext"`
	CreatedAt               time.Time        `json:"created_at"`
	UpdatedAt               time.Time        `json:"updated_at"`

	// Relationships
	OrderPartner          *OrderPartner          `json:"-" gorm:"foreignKey:OrderPartnerID"`
	ShippingPartner       *ShippingPartner       `json:"-" gorm:"foreignKey:ShippingPartnerID"`
	PartnerShippingMethod *PartnerShippingMethod `json:"-" gorm:"foreignKey:PartnerShippingMethodID"`
	ActionLogs            []ActionLog            `json:"-" gorm:"foreignKey:ActionableID;polymorphic:Actionable"`
	//StatusHistory         *StatusHistory         `json:"-" gorm:"foreignKey:EntityID;polymorphic:Entity"`
}

func (Order) TableName() string {
	return "orders"
}

// Helper methods
func (o *Order) IsInternational() bool {
	if o.IsReverseOrder() {
		// return o.PickupDetails.CountryID != o.DropDetails.OmnifulCountryID
	}
	// return o.PickupDetails.CountryID != o.DropDetails.OmnifulCountryID
	return false
}

func (o *Order) IsReverseOrder() bool {

	return o.ShipmentDetails.IsReturnOrder
}

type PickupDetails struct {
	HubID              string          `json:"hub_id,omitempty"`
	HubCode            string          `json:"hub_code,omitempty"`
	Name               string          `json:"name,omitempty"`
	Phone              string          `json:"phone,omitempty"`
	CountryCallingCode string          `json:"country_calling_code,omitempty"`
	Email              string          `json:"email,omitempty"`
	Timezone           string          `json:"timezone,omitempty"`
	CountryID          string          `json:"country_id,omitempty"`
	StateID            string          `json:"state_id,omitempty"`
	CityID             string          `json:"city_id,omitempty"`
	City               string          `json:"city,omitempty"`
	District           string          `json:"district,omitempty"`
	State              string          `json:"state,omitempty"`
	Country            string          `json:"country,omitempty"`
	Pincode            string          `json:"pincode" validate:"required"`
	CountryCode        string          `json:"country_code,omitempty"`
	Latitude           float64         `json:"latitude,omitempty"`
	Longitude          float64         `json:"longitude,omitempty"`
	Address            string          `json:"address,omitempty"`
	StateCode          string          `json:"state_code,omitempty"`
	SellerDetails      *SellerDetails  `json:"seller_details,omitempty"`
	VillageDetails     *VillageDetails `json:"village_details,omitempty"`
}

// DropDetails contains information about the drop (delivery) location.
type DropDetails struct {
	Address             string          `json:"address,omitempty"`
	Phone               string          `json:"phone,omitempty"`
	MobileNumber        *MobileNumber   `json:"mobile_number,omitempty"`
	Country             string          `json:"country,omitempty"`
	CountryCode         string          `json:"country_code" validate:"required"`
	CountryCurrencyCode string          `json:"country_currency_code,omitempty"`
	State               string          `json:"state,omitempty"`
	City                string          `json:"city,omitempty"`
	CityCode            string          `json:"city_code,omitempty"`
	PostalCode          string          `json:"postal_code,omitempty"`
	Pincode             string          `json:"pincode,omitempty"`
	Name                string          `json:"name,omitempty"`
	HubCode             string          `json:"hub_code,omitempty"`
	HubID               string          `json:"hub_id,omitempty"`
	Email               string          `json:"email,omitempty"`
	Latitude            float64         `json:"latitude,omitempty"`
	Longitude           float64         `json:"longitude,omitempty"`
	StateCode           string          `json:"state_code,omitempty"`
	Area                string          `json:"area,omitempty"`
	District            string          `json:"district,omitempty"`
	OmnifulCountry      string          `json:"omniful_country,omitempty"`
	OmnifulCountryID    string          `json:"omniful_country_id,omitempty"`
	OmnifulState        string          `json:"omniful_state,omitempty"`
	OmnifulStateID      string          `json:"omniful_state_id,omitempty"`
	OmnifulCity         string          `json:"omniful_city,omitempty"`
	OmnifulCityID       string          `json:"omniful_city_id,omitempty"`
	VillageDetails      *VillageDetails `json:"village_details,omitempty"`
	SalesChannelCity    string          `json:"sales_channel_city,omitempty"`
}

// ShipmentDetails contains detailed information about the shipment.
type ShipmentDetails struct {
	CourierPartnerID     int                  `json:"courier_partner_id,omitempty"`
	TotalSKUCount        int                  `json:"total_sku_count" validate:"required"`
	TotalItemCount       int                  `json:"total_item_count" validate:"required"`
	CourierPartner       *OrderCourierPartner `json:"courier_partner,omitempty"`
	Slot                 *DeliverySlot        `json:"slot,omitempty"`
	Tags                 []Tag                `json:"tags,omitempty"`
	OrderType            string               `json:"order_type,omitempty"`
	Remarks              string               `json:"remarks,omitempty"`
	NumberOfBoxes        int                  `json:"number_of_boxes" validate:"required,gt=0"`
	Height               int                  `json:"height,omitempty"`
	Length               int                  `json:"length,omitempty"`
	Breadth              int                  `json:"breadth,omitempty"`
	Weight               float64              `json:"weight" validate:"required"`
	CodValue             float64              `json:"cod_value,omitempty"`
	Count                int                  `json:"count,omitempty"`
	Currency             string               `json:"currency,omitempty"`
	PaymentType          string               `json:"payment_type,omitempty"`
	InvoiceValue         float64              `json:"invoice_value" validate:"required"`
	OrderValue           float64              `json:"order_value,omitempty"`
	OrderNote            string               `json:"order_note,omitempty"`
	OrderCreatedAt       string               `json:"order_created_at,omitempty"`
	TotalPaid            float64              `json:"total_paid" validate:"required"`
	TotalDue             float64              `json:"total_due" validate:"required"`
	IsReturnOrder        bool                 `json:"is_return_order"`
	TotalDuePriceSet     *PriceSet            `json:"total_due_price_set,omitempty"`
	TotalPaidPriceSet    *PriceSet            `json:"total_paid_price_set,omitempty"`
	InvoiceValuePriceSet *PriceSet            `json:"invoice_value_price_set,omitempty"`
	OrderValuePriceSet   *PriceSet            `json:"order_value_price_set,omitempty"`
	ExchangeRate         *ExchangeRate        `json:"exchange_rate,omitempty"`
	InvoiceNumber        string               `json:"invoice_number,omitempty"`
	InvoiceDate          string               `json:"invoice_date,omitempty"`
	PackedQuantity       int                  `json:"packed_quantity,omitempty"`
	ContainsFrozenItems  bool                 `json:"contains_frozen_items,omitempty"`
	Description          string               `json:"description,omitempty"`
	ServiceType          string               `json:"service_type,omitempty"`
	ShippingMethod       string               `json:"shipping_method,omitempty"`
	Items                []Item               `json:"items" validate:"required"`
	PackageDetails       []PackageDetail      `json:"package_details,omitempty"`
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
	Name               string         `json:"name" validate:"required"`
	Phone              string         `json:"phone" validate:"required"`
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
	SKUName         string          `json:"sku_name" validate:"required"`
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
	AwbLabel       string             `json:"awb_label,omitempty"`
	AwbNumber      string             `json:"awb_number,omitempty"`
	DeliveryType   string             `json:"delivery_type,omitempty"`
	TaxNumber      string             `json:"tax_number,omitempty"`
	EnableWhatsapp bool               `json:"enable_whatsapp,omitempty"`
	IsFragile      bool               `json:"is_fragile,omitempty"`
	IsDangerous    bool               `json:"is_dangerous,omitempty"`
	Label          bool               `json:"label,omitempty"`
	ReturnInfo     *OrderReturnInfo   `json:"return_info,omitempty"`
	ResellerInfo   *OrderResellerInfo `json:"reseller_info,omitempty"`
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
