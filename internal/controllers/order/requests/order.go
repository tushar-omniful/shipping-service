package requests

import (
	"time"

	"github.com/omniful/shipping-service/internal/domain/models"
	svcRequests "github.com/omniful/shipping-service/internal/services/orders/requests"
)

// CreateForwardOrder represents the request structure for creating a forward order
type CreateForwardOrder struct {
	ShippingPartnerTag   string    `json:"shipping_partner_tag,omitempty"`
	TenantID             string    `json:"tenant_id" validate:"required"`
	SellerID             string    `json:"seller_id" validate:"required"`
	AccountID            string    `json:"account_id" validate:"required"`
	OmnifulRuleID        string    `json:"omniful_rule_id,omitempty"`
	SellerSalesChannelID string    `json:"seller_sales_channel_id" validate:"required"`
	Data                 OrderData `json:"data"`
}

// CancelOrderRequest represents the request structure for canceling an order
type CancelOrderRequest struct {
	OrderID            string `json:"order_id" validate:"required"`
	CancellationReason string `json:"cancellation_reason" validate:"required"`
	CancelledBy        string `json:"cancelled_by" validate:"required"`
	Notes              string `json:"notes,omitempty"`
}

// OrderData contains the main order information
type OrderData struct {
	ShipmentType              string                 `json:"shipment_type,omitempty"`
	OrderSource               string                 `json:"order_source,omitempty"`
	OrderAlias                string                 `json:"order_alias,omitempty"`
	OrderPartnerOrderID       string                 `json:"order_partner_order_id" validate:"required"`
	SellerSalesChannelOrderID string                 `json:"seller_sales_channel_order_id" validate:"required"`
	ShippingReference         string                 `json:"shipping_reference,omitempty"`
	OrderTime                 time.Time              `json:"order_time" validate:"required"`
	PickupDetails             PickupDetails          `json:"pickup_details"`
	DropDetails               DropDetails            `json:"drop_details"`
	TaxDetails                *TaxDetails            `json:"tax_details,omitempty"`
	Metadata                  *Metadata              `json:"metadata,omitempty"`
	ShipmentDetails           ShipmentDetails        `json:"shipment_details"`
	PreShipmentDetails        map[string]interface{} `json:"pre_shipment_details,omitempty"`
}

// TaxDetails contains tax and cost information
type TaxDetails struct {
	TaxableValue         int    `json:"taxable_value,omitempty"`
	EwaybillSerialNumber string `json:"ewaybill_serial_number,omitempty"`
	PlaceOfSupply        string `json:"place_of_supply,omitempty"`
	HSNCode              string `json:"hsn_code,omitempty"`
	InvoiceReference     string `json:"invoice_reference,omitempty"`
}

type Metadata struct {
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

func (m *Metadata) ToModelMetadata() *models.OrderMetadata {
	if m == nil {
		return nil
	}
	return &models.OrderMetadata{
		OmnifulAwbLabel: m.AwbLabel,
		AwbNumber:       m.AwbNumber,
		DeliveryType:    m.DeliveryType,
		TaxNumber:       m.TaxNumber,
		EnableWhatsapp:  m.EnableWhatsapp,
		IsFragile:       m.IsFragile,
		IsDangerous:     m.IsDangerous,
		Label:           m.Label,
		ReturnInfo:      m.ReturnInfo.ToModelReturnInfo(),
		ResellerInfo:    m.ResellerInfo.ToModelResellerInfo(),
	}
}

type OrderReturnInfo struct {
	PostalCode string `json:"postal_code,omitempty"`
	Address    string `json:"address,omitempty"`
	State      string `json:"state,omitempty"`
	Phone      string `json:"phone,omitempty"`
	Name       string `json:"name,omitempty"`
	City       string `json:"city,omitempty"`
	Country    string `json:"country,omitempty"`
}

type OrderResellerInfo struct {
	Name  string `json:"name,omitempty"`
	Phone string `json:"phone,omitempty"`
}

func (r *OrderReturnInfo) ToModelReturnInfo() *models.OrderReturnInfo {
	if r == nil {
		return nil
	}
	return &models.OrderReturnInfo{
		PostalCode: r.PostalCode,
		Address:    r.Address,
		State:      r.State,
		Phone:      r.Phone,
		Name:       r.Name,
		City:       r.City,
		Country:    r.Country,
	}
}

func (r *OrderResellerInfo) ToModelResellerInfo() *models.OrderResellerInfo {
	if r == nil {
		return nil
	}
	return &models.OrderResellerInfo{
		Name:  r.Name,
		Phone: r.Phone,
	}
}

type ShipmentDetails struct {
	CourierPartnerID     int             `json:"courier_partner_id,omitempty"`
	TotalSKUCount        int             `json:"total_sku_count" `
	TotalItemCount       int             `json:"total_item_count"`
	CourierPartner       *CourierPartner `json:"courier_partner,omitempty"`
	Slot                 *DeliverySlot   `json:"slot,omitempty"`
	Tags                 []Tag           `json:"tags,omitempty"`
	OrderType            string          `json:"order_type,omitempty"`
	Remarks              string          `json:"remarks,omitempty"`
	NumberOfBoxes        int             `json:"number_of_boxes" validate:"required,gt=0"`
	Height               int             `json:"height,omitempty"`
	Length               int             `json:"length,omitempty"`
	Breadth              int             `json:"breadth,omitempty"`
	Weight               float64         `json:"weight" validate:"required"`
	CodValue             float64         `json:"cod_value,omitempty"`
	Count                int             `json:"count,omitempty"`
	Currency             string          `json:"currency,omitempty"`
	PaymentType          string          `json:"payment_type,omitempty"`
	InvoiceValue         float64         `json:"invoice_value"`
	OrderValue           float64         `json:"order_value,omitempty"`
	OrderNote            string          `json:"order_note,omitempty"`
	OrderCreatedAt       string          `json:"order_created_at,omitempty"`
	TotalPaid            float64         `json:"total_paid"`
	TotalDue             float64         `json:"total_due"`
	TotalDuePriceSet     *PriceSet       `json:"total_due_price_set,omitempty"`
	TotalPaidPriceSet    *PriceSet       `json:"total_paid_price_set,omitempty"`
	InvoiceValuePriceSet *PriceSet       `json:"invoice_value_price_set,omitempty"`
	OrderValuePriceSet   *PriceSet       `json:"order_value_price_set,omitempty"`
	ExchangeRate         *ExchangeRate   `json:"exchange_rate,omitempty"`
	InvoiceNumber        string          `json:"invoice_number,omitempty"`
	InvoiceDate          string          `json:"invoice_date,omitempty"`
	PackedQuantity       int             `json:"packed_quantity,omitempty"`
	ContainsFrozenItems  bool            `json:"contains_frozen_items,omitempty"`
	Description          string          `json:"description,omitempty"`
	ServiceType          string          `json:"service_type,omitempty"`
	ShippingMethod       string          `json:"shipping_method,omitempty"`
	Items                []Item          `json:"items"`
	PackageDetails       []PackageDetail `json:"package_details,omitempty"`
}

func (s *ShipmentDetails) ToModelShipmentDetails() *models.ShipmentDetails {
	return &models.ShipmentDetails{
		CourierPartnerID: s.CourierPartnerID,
		TotalSKUCount:    s.TotalSKUCount,
		TotalItemCount:   s.TotalItemCount,
		CourierPartner:   s.CourierPartner.ToModelCourierPartner(),
		Slot:             s.Slot.ToModelDeliverySlot(),
		Tags:             s.ToModelTags(s.Tags),
		OrderType:        s.OrderType,
		Remarks:          s.Remarks,
		NumberOfBoxes:    s.NumberOfBoxes,
		Height:           s.Height,
		Length:           s.Length,
		Breadth:          s.Breadth,
		Weight:           s.Weight,
		CodValue:         s.CodValue,
		Count:            s.Count,
		Currency:         s.Currency,
		PaymentType:      s.PaymentType,
		InvoiceValue:     s.InvoiceValue,
		OrderValue:       s.OrderValue,
		OrderNote:        s.OrderNote,
		OrderCreatedAt:   s.OrderCreatedAt,
		TotalPaid:        s.TotalPaid,
		TotalDue:         s.TotalDue,
		Items:            s.ToModelItems(s.Items),
		PackageDetails:   s.ToModelPackageDetails(s.PackageDetails),
	}
}

type CourierPartner struct {
	ID   uint64 `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Tag  string `json:"tag,omitempty"`
	Logo string `json:"logo,omitempty"`
}

type DeliverySlot struct {
	DeliveryDate string `json:"delivery_date,omitempty"`
	StartTime    int    `json:"start_time,omitempty"`
	EndTime      int    `json:"end_time,omitempty"`
}

type Tag struct {
	ID    string `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Color string `json:"color,omitempty"`
}

type PriceSet struct {
	OrderCurrency Currency `json:"order_currency"`
	StoreCurrency Currency `json:"store_currency"`
}

type Currency struct {
	Amount       float64 `json:"amount,omitempty"`
	CurrencyCode string  `json:"currency_code,omitempty"`
}

type ExchangeRate struct {
	OrderCurrency string  `json:"order_currency,omitempty"`
	StoreCurrency string  `json:"store_currency,omitempty"`
	Rate          float64 `json:"rate"`
}

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
	Additional      *ItemAdditional `json:"additional"`
}

func (i *Item) ToModelItem() models.Item {
	return models.Item{
		ProductURL:      i.ProductURL,
		Price:           i.Price,
		Description:     i.Description,
		Quantity:        i.Quantity,
		PackedQuantity:  i.PackedQuantity,
		SKU:             i.SKU,
		SKUName:         i.SKUName,
		CountryOfOrigin: i.CountryOfOrigin,
		Weight:          i.Weight.ToModelItemWeight(),
		Additional:      i.Additional.ToModelItemAdditional(),
	}
}

type ItemWeight struct {
	Value float64 `json:"value,omitempty"`
	UOM   string  `json:"uom,omitempty"`
}

type ItemAdditional struct {
	Height     int         `json:"height"`
	Length     int         `json:"length"`
	Breadth    int         `json:"breadth"`
	Weight     int         `json:"weight"`
	Images     []ItemImage `json:"images,omitempty"`
	ReturnDays int         `json:"return_days,omitempty"`
}

type ItemImage struct {
	URL         string `json:"url,omitempty"`
	Type        string `json:"type,omitempty"`
	Description string `json:"description,omitempty"`
}

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

type Dimensions struct {
	Length  int    `json:"length" validate:"required"`
	Breadth int    `json:"breadth" validate:"required"`
	Height  int    `json:"height" validate:"required"`
	UOM     string `json:"uom,omitempty"`
}

type Weight struct {
	Value float64 `json:"value" validate:"required"`
	UOM   string  `json:"uom,omitempty"`
}

type PickupDetails struct {
	HubID              string          `json:"hub_id,omitempty"`
	HubCode            string          `json:"hub_code,omitempty"`
	Name               string          `json:"name,omitempty"`
	Phone              string          `json:"phone,omitempty"`
	CountryCallingCode string          `json:"country_calling_code,omitempty"`
	Email              string          `json:"email,omitempty"`
	Timezone           string          `json:"timezone,omitempty"`
	CountryID          string          `json:"country_id" validate:"required"`
	StateID            string          `json:"state_id,omitempty"`
	CityID             string          `json:"city_id" validate:"required"`
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
	SellerDetails      *SellerDetails  `json:"seller_details" validate:"required"`
	VillageDetails     *VillageDetails `json:"village_details,omitempty"`
}

func (p *PickupDetails) ToModelPickupDetails() *models.PickupDetails {
	return &models.PickupDetails{
		HubID:              p.HubID,
		HubCode:            p.HubCode,
		Name:               p.Name,
		Phone:              p.Phone,
		CountryCallingCode: p.CountryCallingCode,
		Email:              p.Email,
		Timezone:           p.Timezone,
		CountryID:          p.CountryID,
		StateID:            p.StateID,
		CityID:             p.CityID,
		City:               p.City,
		District:           p.District,
		State:              p.State,
		Country:            p.Country,
		Pincode:            p.Pincode,
		CountryCode:        p.CountryCode,
		Latitude:           p.Latitude,
		Longitude:          p.Longitude,
		Address:            p.Address,
		StateCode:          p.StateCode,
		SellerDetails:      p.SellerDetails.ToModelSellerDetails(),
		VillageDetails:     p.VillageDetails.ToModelVillageDetails(),
	}
}

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
	OmnifulCountry      string          `json:"omniful_country" validate:"required"`
	OmnifulCountryID    string          `json:"omniful_country_id" validate:"required"`
	OmnifulState        string          `json:"omniful_state" validate:"required"`
	OmnifulStateID      string          `json:"omniful_state_id" validate:"required"`
	OmnifulCity         string          `json:"omniful_city" validate:"required"`
	OmnifulCityID       string          `json:"omniful_city_id" validate:"required"`
	VillageDetails      *VillageDetails `json:"village_details,omitempty"`
	SalesChannelCity    string          `json:"sales_channel_city,omitempty"`
}

func (d *DropDetails) ToModelDropDetails() *models.DropDetails {
	return &models.DropDetails{
		Address:             d.Address,
		Phone:               d.Phone,
		MobileNumber:        d.MobileNumber.ToModelMobileNumber(),
		Country:             d.Country,
		CountryCode:         d.CountryCode,
		CountryCurrencyCode: d.CountryCurrencyCode,
		State:               d.State,
		City:                d.City,
		CityCode:            d.CityCode,
		PostalCode:          d.PostalCode,
		Pincode:             d.Pincode,
		Name:                d.Name,
		HubCode:             d.HubCode,
		HubID:               d.HubID,
		Email:               d.Email,
		Latitude:            d.Latitude,
		Longitude:           d.Longitude,
		StateCode:           d.StateCode,
		Area:                d.Area,
		District:            d.District,
		OmnifulCountry:      d.OmnifulCountry,
		OmnifulCountryID:    d.OmnifulCountryID,
		OmnifulState:        d.OmnifulState,
		OmnifulStateID:      d.OmnifulStateID,
		OmnifulCity:         d.OmnifulCity,
		OmnifulCityID:       d.OmnifulCityID,
		VillageDetails:      d.VillageDetails.ToModelVillageDetails(),
		SalesChannelCity:    d.SalesChannelCity,
	}
}

// MobileNumber represents a mobile phone number with country information
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

func (t *TaxDetails) ToModelTaxDetails() *models.TaxDetails {
	if t == nil {
		return nil
	}
	return &models.TaxDetails{
		TaxableValue:         t.TaxableValue,
		EwaybillSerialNumber: t.EwaybillSerialNumber,
		PlaceOfSupply:        t.PlaceOfSupply,
		HSNCode:              t.HSNCode,
		InvoiceReference:     t.InvoiceReference,
	}
}

func (s *SellerDetails) ToModelSellerDetails() *models.SellerDetails {
	if s == nil {
		return nil
	}
	return &models.SellerDetails{
		ID:                 s.ID,
		Name:               s.Name,
		Phone:              s.Phone,
		CountryCode:        s.CountryCode,
		CountryCallingCode: s.CountryCallingCode,
		Address:            s.Address.ToModelSellerAddress(),
	}
}

func (a *SellerAddress) ToModelSellerAddress() *models.SellerAddress {
	if a == nil {
		return nil
	}
	return &models.SellerAddress{
		AddressLine1: a.AddressLine1,
		AddressLine2: a.AddressLine2,
		Country:      a.Country.ToModelSellerAddressCountry(),
		State:        a.State.ToModelSellerAddressState(),
		City:         a.City.ToModelSellerAddressCity(),
		Pincode:      a.Pincode,
	}
}

func (c *CourierPartner) ToModelCourierPartner() *models.OrderCourierPartner {
	if c == nil {
		return nil
	}
	return &models.OrderCourierPartner{
		ID:   c.ID,
		Name: c.Name,
		Tag:  c.Tag,
		Logo: c.Logo,
	}
}

func (s *ShipmentDetails) ToModelTags(tags []Tag) []models.Tag {
	if tags == nil {
		return nil
	}
	modelTags := make([]models.Tag, len(tags))
	for i, t := range tags {
		modelTags[i] = models.Tag{
			ID:    t.ID,
			Name:  t.Name,
			Color: t.Color,
		}
	}
	return modelTags
}

func (s *ShipmentDetails) ToModelItems(items []Item) []models.Item {
	if items == nil {
		return nil
	}
	modelItems := make([]models.Item, len(items))
	for i, item := range items {
		modelItems[i] = item.ToModelItem()
	}
	return modelItems
}

func (s *ShipmentDetails) ToModelPackageDetails(details []PackageDetail) []models.PackageDetail {
	if details == nil {
		return nil
	}
	modelDetails := make([]models.PackageDetail, len(details))
	for i, d := range details {
		modelDetails[i] = *d.ToModelPackageDetail()
	}
	return modelDetails
}

func (p *PackageDetail) ToModelPackageDetail() *models.PackageDetail {
	return &models.PackageDetail{
		PackageID:        p.PackageID,
		Dimensions:       p.Dimensions.ToModelDimensions(),
		Weight:           p.Weight.ToModelWeight(),
		InvoiceTotal:     p.InvoiceTotal,
		InvoiceSubTotal:  p.InvoiceSubTotal,
		InvoiceTotalPaid: p.InvoiceTotalPaid,
		InvoiceTotalDue:  p.InvoiceTotalDue,
		Items:            p.ToModelItems(p.Items),
	}
}

func (p *PackageDetail) ToModelItems(items []Item) []models.Item {
	if items == nil {
		return nil
	}
	modelItems := make([]models.Item, len(items))
	for i, item := range items {
		modelItems[i] = item.ToModelItem()
	}
	return modelItems
}

func (d *Dimensions) ToModelDimensions() models.Dimensions {
	return models.Dimensions{
		Length:  d.Length,
		Breadth: d.Breadth,
		Height:  d.Height,
		UOM:     d.UOM,
	}
}

func (w *Weight) ToModelWeight() models.Weight {
	return models.Weight{
		Value: w.Value,
		UOM:   w.UOM,
	}
}

func (w *ItemWeight) ToModelItemWeight() *models.ItemWeight {
	if w == nil {
		return nil
	}
	return &models.ItemWeight{
		Value: w.Value,
		UOM:   w.UOM,
	}
}

func (c *SellerAddressCountry) ToModelSellerAddressCountry() *models.SellerAddressCountry {
	if c == nil {
		return nil
	}
	return &models.SellerAddressCountry{
		ID:   c.ID,
		Name: c.Name,
		Code: c.Code,
	}
}

func (s *SellerAddressState) ToModelSellerAddressState() *models.SellerAddressState {
	if s == nil {
		return nil
	}
	return &models.SellerAddressState{
		ID:   s.ID,
		Name: s.Name,
	}
}

func (c *SellerAddressCity) ToModelSellerAddressCity() *models.SellerAddressCity {
	if c == nil {
		return nil
	}
	return &models.SellerAddressCity{
		ID:   c.ID,
		Name: c.Name,
	}
}

func (d *DeliverySlot) ToModelDeliverySlot() *models.DeliverySlot {
	if d == nil {
		return nil
	}
	return &models.DeliverySlot{
		DeliveryDate: d.DeliveryDate,
		StartTime:    d.StartTime,
		EndTime:      d.EndTime,
	}
}

func (i *ItemAdditional) ToModelItemAdditional() *models.ItemAdditional {
	if i == nil {
		return nil
	}
	return &models.ItemAdditional{
		Height:     i.Height,
		Length:     i.Length,
		Breadth:    i.Breadth,
		Weight:     i.Weight,
		Images:     ToModelItemImages(i.Images),
		ReturnDays: i.ReturnDays,
	}
}

func ToModelItemImages(images []ItemImage) []models.ItemImage {
	if images == nil {
		return nil
	}
	modelImages := make([]models.ItemImage, len(images))
	for i, img := range images {
		modelImages[i] = models.ItemImage{
			URL:         img.URL,
			Type:        img.Type,
			Description: img.Description,
		}
	}
	return modelImages
}

func (m *MobileNumber) ToModelMobileNumber() *models.MobileNumber {
	if m == nil {
		return nil
	}
	return &models.MobileNumber{
		Number:             m.Number,
		CountryCallingCode: m.CountryCallingCode,
		CountryCode:        m.CountryCode,
	}
}

func (v *VillageDetails) ToModelVillageDetails() *models.VillageDetails {
	if v == nil {
		return nil
	}
	return &models.VillageDetails{
		ID:         v.ID,
		NameEn:     v.NameEn,
		CityID:     v.CityID,
		CityName:   v.CityName,
		RegionID:   v.RegionID,
		RegionName: v.RegionName,
	}
}

func (c *CreateForwardOrder) ToServiceRequest() *svcRequests.CreateForwardOrder {
	return &svcRequests.CreateForwardOrder{

		ShippingPartnerTag:   c.ShippingPartnerTag,
		TenantID:             c.TenantID,
		SellerID:             c.SellerID,
		AccountID:            c.AccountID,
		OmnifulRuleID:        c.OmnifulRuleID,
		SellerSalesChannelID: c.SellerSalesChannelID,
		Data: svcRequests.OrderData{
			ShipmentType:              c.Data.ShipmentType,
			OrderSource:               c.Data.OrderSource,
			OrderAlias:                c.Data.OrderAlias,
			OrderPartnerOrderID:       c.Data.OrderPartnerOrderID,
			SellerSalesChannelOrderID: c.Data.SellerSalesChannelOrderID,
			ShippingReference:         c.Data.ShippingReference,
			OrderTime:                 c.Data.OrderTime,
			PickupDetails:             c.Data.PickupDetails.ToModelPickupDetails(),
			DropDetails:               c.Data.DropDetails.ToModelDropDetails(),
			TaxDetails:                c.Data.TaxDetails.ToModelTaxDetails(),
			Metadata:                  c.Data.Metadata.ToModelMetadata(),
			ShipmentDetails:           c.Data.ShipmentDetails.ToModelShipmentDetails(),
			PreShipmentDetails:        c.Data.PreShipmentDetails,
		},
	}
}
