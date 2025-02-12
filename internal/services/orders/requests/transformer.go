package requests

// import (
// 	"fmt"
// 	"github.com/omniful/shipping-service/internal/domain/models"
// 	shipmentRequest "github.com/omniful/shipping-service/internal/services/shipment/requests"
// )

// func (cfo *CreateForwardOrder) TransformToIntegrationRequest(psm *models.PartnerShippingMethod) (shipmentRequest.CreateForwardOrder, error) {
// 	// Create integration request
// 	integrationRequest := shipmentRequest.CreateForwardOrder{
// 		ShippingPartnerTag:   cfo.ShippingPartnerTag,
// 		TenantID:             cfo.TenantID,
// 		SellerID:             cfo.SellerID,
// 		AccountID:            cfo.AccountID,
// 		OmnifulRuleID:        cfo.OmnifulRuleID,
// 		SellerSalesChannelID: cfo.SellerSalesChannelID,
// 		Data:                 shipmentRequest.OrderData{},
// 	}

// 	// Apply AWB configuration from partner shipping method
// 	if err := cfo.applyAWBConfig(&cfo.Data, psm.AWBConfig); err != nil {
// 		return shipmentRequest.CreateForwardOrder{}, fmt.Errorf("error applying AWB config: %v", err)
// 	}

// 	// Handle currency conversion
// 	if err := cfo.handleCurrencyConversion(&cfo.Data); err != nil {
// 		return shipmentRequest.CreateForwardOrder{}, fmt.Errorf("error handling currency conversion: %v", err)
// 	}

// 	// Copy transformed data
// 	integrationRequest.Data = shipmentRequest.OrderData{
// 		ShipmentType:              cfo.Data.ShipmentType,
// 		OrderSource:               cfo.Data.OrderSource,
// 		OrderAlias:                cfo.Data.OrderAlias,
// 		OrderPartnerOrderID:       cfo.Data.OrderPartnerOrderID,
// 		SellerSalesChannelOrderID: cfo.Data.SellerSalesChannelOrderID,
// 		ShippingReference:         cfo.Data.ShippingReference,
// 		OrderTime:                 cfo.Data.OrderTime,
// 		PickupDetails:             transformPickupDetails(cfo.Data.PickupDetails),
// 		DropDetails:               transformDropDetails(cfo.Data.DropDetails),
// 		TaxDetails:                transformCostDetails(cfo.Data.TaxDetails),
// 		Metadata:                  transformMetadata(cfo.Data.Metadata),
// 		ShipmentDetails:           transformShipmentDetails(cfo.Data.ShipmentDetails),
// 		PreShipmentDetails:        cfo.Data.PreShipmentDetails,
// 	}

// 	return integrationRequest, nil
// }

// func (cfo *CreateForwardOrder) applyAWBConfig(orderData *OrderData, awbConfig map[string]interface{}) error {
// 	if awbConfig == nil {
// 		return nil
// 	}

// 	// Store original hub name
// 	hubName := orderData.PickupDetails.Name
// 	orderData.PickupDetails.HubCode = hubName

// 	// Apply seller configuration
// 	if sellerConfig, ok := awbConfig["seller"].(map[string]interface{}); ok {
// 		applySellerConfig(&orderData.PickupDetails, sellerConfig)
// 	}

// 	// Apply custom configuration
// 	if customConfig, ok := awbConfig["custom"].(map[string]interface{}); ok {
// 		applyCustomConfig(&orderData.PickupDetails, customConfig)
// 	}

// 	return nil
// }

// func applySellerConfig(pickupDetails *PickupDetails, sellerConfig map[string]interface{}) {
// 	if isEnabled(sellerConfig, "name_enabled") && pickupDetails.SellerDetails != nil {
// 		pickupDetails.Name = pickupDetails.SellerDetails.Name
// 	}
// 	if isEnabled(sellerConfig, "phone_enabled") && pickupDetails.SellerDetails != nil {
// 		pickupDetails.Phone = pickupDetails.SellerDetails.Phone
// 	}
// 	if isEnabled(sellerConfig, "address_enabled") && pickupDetails.SellerDetails != nil && pickupDetails.SellerDetails.Address != nil {
// 		applyAddressToPickupDetails(pickupDetails, pickupDetails.SellerDetails.Address)
// 	}
// }

// func applyCustomConfig(pickupDetails *PickupDetails, customConfig map[string]interface{}) {
// 	if isEnabled(customConfig, "name_enabled") {
// 		if name, ok := customConfig["name"].(string); ok {
// 			pickupDetails.Name = name
// 		}
// 	}
// 	if isEnabled(customConfig, "phone_enabled") {
// 		if phone, ok := customConfig["phone"].(string); ok {
// 			pickupDetails.Phone = parsePhoneNumber(phone)
// 		}
// 	}
// 	if isEnabled(customConfig, "email_enabled") {
// 		if email, ok := customConfig["email"].(string); ok {
// 			pickupDetails.Email = email
// 		}
// 	}
// }

// func (cfo *CreateForwardOrder) handleCurrencyConversion(orderData *OrderData) error {
// 	if orderData.ShipmentDetails.Currency == orderData.DropDetails.CountryCurrencyCode {
// 		return nil
// 	}

// 	// Create currency conversion object
// 	conversion := &CurrencyConversion{
// 		InvoiceValue:         orderData.ShipmentDetails.InvoiceValue,
// 		OrderValue:           orderData.ShipmentDetails.OrderValue,
// 		TotalPaid:            orderData.ShipmentDetails.TotalPaid,
// 		TotalDue:             orderData.ShipmentDetails.TotalDue,
// 		Currency:             orderData.ShipmentDetails.Currency,
// 		InvoiceValuePriceSet: convertPriceSetToMap(orderData.ShipmentDetails.InvoiceValuePriceSet),
// 		OrderValuePriceSet:   convertPriceSetToMap(orderData.ShipmentDetails.OrderValuePriceSet),
// 		TotalPaidPriceSet:    convertPriceSetToMap(orderData.ShipmentDetails.TotalPaidPriceSet),
// 		TotalDuePriceSet:     convertPriceSetToMap(orderData.ShipmentDetails.TotalDuePriceSet),
// 		ExchangeRate:         convertExchangeRateToMap(orderData.ShipmentDetails.ExchangeRate),
// 		DropCountryCurrency:  orderData.DropDetails.CountryCurrencyCode,
// 	}

// 	// Convert values
// 	convertedValues, err := convertCurrencyValues(conversion)
// 	if err != nil {
// 		return err
// 	}

// 	// Update shipment details
// 	orderData.ShipmentDetails.InvoiceValue = convertedValues["invoice_value"]
// 	orderData.ShipmentDetails.OrderValue = convertedValues["order_value"]
// 	orderData.ShipmentDetails.TotalPaid = convertedValues["total_paid"]
// 	orderData.ShipmentDetails.TotalDue = convertedValues["total_due"]
// 	orderData.ShipmentDetails.Currency = conversion.DropCountryCurrency

// 	return nil
// }

// // Helper functions for transforming between request types
// func transformPickupDetails(pd PickupDetails) shipmentRequest.PickupDetails {
// 	return shipmentRequest.PickupDetails{
// 		HubID:              pd.HubID,
// 		HubCode:            pd.HubCode,
// 		Name:               pd.Name,
// 		Phone:              pd.Phone,
// 		CountryCallingCode: pd.CountryCallingCode,
// 		Email:              pd.Email,
// 		Timezone:           pd.Timezone,
// 		CountryID:          pd.CountryID,
// 		StateID:            pd.StateID,
// 		CityID:             pd.CityID,
// 		City:               pd.City,
// 		District:           pd.District,
// 		State:              pd.State,
// 		Country:            pd.Country,
// 		Pincode:            pd.Pincode,
// 		CountryCode:        pd.CountryCode,
// 		Latitude:           pd.Latitude,
// 		Longitude:          pd.Longitude,
// 		Address:            pd.Address,
// 		StateCode:          pd.StateCode,
// 		SellerDetails:      transformSellerDetails(pd.SellerDetails),
// 		VillageDetails:     transformVillageDetails(pd.VillageDetails),
// 	}
// }

// func transformDropDetails(dd DropDetails) shipmentRequest.DropDetails {
// 	return shipmentRequest.DropDetails{
// 		Address:             dd.Address,
// 		Phone:               dd.Phone,
// 		MobileNumber:        transformMobileNumber(dd.MobileNumber),
// 		Country:             dd.Country,
// 		CountryCode:         dd.CountryCode,
// 		CountryCurrencyCode: dd.CountryCurrencyCode,
// 		State:               dd.State,
// 		City:                dd.City,
// 		CityCode:            dd.CityCode,
// 		PostalCode:          dd.PostalCode,
// 		Pincode:             dd.Pincode,
// 		Name:                dd.Name,
// 		HubCode:             dd.HubCode,
// 		HubID:               dd.HubID,
// 		Email:               dd.Email,
// 		Latitude:            dd.Latitude,
// 		Longitude:           dd.Longitude,
// 		StateCode:           dd.StateCode,
// 		Area:                dd.Area,
// 		District:            dd.District,
// 		OmnifulCountry:      dd.OmnifulCountry,
// 		OmnifulCountryID:    dd.OmnifulCountryID,
// 		OmnifulState:        dd.OmnifulState,
// 		OmnifulStateID:      dd.OmnifulStateID,
// 		OmnifulCity:         dd.OmnifulCity,
// 		OmnifulCityID:       dd.OmnifulCityID,
// 		VillageDetails:      transformVillageDetails(dd.VillageDetails),
// 		SalesChannelCity:    dd.SalesChannelCity,
// 	}
// }

// func transformCostDetails(cd *CostDetails) *shipmentRequest.CostDetails {
// 	if cd == nil {
// 		return nil
// 	}
// 	return &shipmentRequest.CostDetails{
// 		TaxableValue:         cd.TaxableValue,
// 		EwaybillSerialNumber: cd.EwaybillSerialNumber,
// 		PlaceOfSupply:        cd.PlaceOfSupply,
// 		HSNCode:              cd.HSNCode,
// 		InvoiceReference:     cd.InvoiceReference,
// 	}
// }

// func transformMetadata(md *Metadata) *shipmentRequest.Metadata {
// 	if md == nil {
// 		return nil
// 	}
// 	return &shipmentRequest.Metadata{
// 		AwbLabel:       md.AwbLabel,
// 		AwbNumber:      md.AwbNumber,
// 		DeliveryType:   md.DeliveryType,
// 		TaxNumber:      md.TaxNumber,
// 		EnableWhatsapp: md.EnableWhatsapp,
// 		IsFragile:      md.IsFragile,
// 		IsDangerous:    md.IsDangerous,
// 		Label:          md.Label,
// 		ReturnInfo:     transformReturnInfo(md.ReturnInfo),
// 		ResellerInfo:   transformResellerInfo(md.ResellerInfo),
// 	}
// }

// func transformShipmentDetails(sd ShipmentDetails) shipmentRequest.ShipmentDetails {
// 	return shipmentRequest.ShipmentDetails{
// 		CourierPartnerID:     sd.CourierPartnerID,
// 		TotalSKUCount:        sd.TotalSKUCount,
// 		TotalItemCount:       sd.TotalItemCount,
// 		CourierPartner:       transformCourierPartner(sd.CourierPartner),
// 		Slot:                 transformDeliverySlot(sd.Slot),
// 		Tags:                 transformTags(sd.Tags),
// 		OrderType:            sd.OrderType,
// 		Remarks:              sd.Remarks,
// 		NumberOfBoxes:        sd.NumberOfBoxes,
// 		Height:               sd.Height,
// 		Length:               sd.Length,
// 		Breadth:              sd.Breadth,
// 		Weight:               sd.Weight,
// 		CodValue:             float64(sd.CodValue),
// 		Count:                sd.Count,
// 		Currency:             sd.Currency,
// 		PaymentType:          sd.PaymentType,
// 		InvoiceValue:         sd.InvoiceValue,
// 		OrderValue:           sd.OrderValue,
// 		OrderNote:            sd.OrderNote,
// 		OrderCreatedAt:       sd.OrderCreatedAt,
// 		TotalPaid:            sd.TotalPaid,
// 		TotalDue:             sd.TotalDue,
// 		TotalDuePriceSet:     transformPriceSet(sd.TotalDuePriceSet),
// 		TotalPaidPriceSet:    transformPriceSet(sd.TotalPaidPriceSet),
// 		InvoiceValuePriceSet: transformPriceSet(sd.InvoiceValuePriceSet),
// 		OrderValuePriceSet:   transformPriceSet(sd.OrderValuePriceSet),
// 		ExchangeRate:         transformExchangeRate(sd.ExchangeRate),
// 		InvoiceNumber:        sd.InvoiceNumber,
// 		InvoiceDate:          sd.InvoiceDate,
// 		PackedQuantity:       sd.PackedQuantity,
// 		ContainsFrozenItems:  sd.ContainsFrozenItems,
// 		Description:          sd.Description,
// 		ServiceType:          sd.ServiceType,
// 		ShippingMethod:       sd.ShippingMethod,
// 		Items:                transformItems(sd.Items),
// 		PackageDetails:       transformPackageDetails(sd.PackageDetails),
// 	}
// }

// func transformSellerDetails(sd *SellerDetails) *shipmentRequest.SellerDetails {
// 	if sd == nil {
// 		return nil
// 	}
// 	return &shipmentRequest.SellerDetails{
// 		ID:                 sd.ID,
// 		Name:               sd.Name,
// 		Phone:              sd.Phone,
// 		CountryCode:        sd.CountryCode,
// 		CountryCallingCode: sd.CountryCallingCode,
// 		Address:            transformSellerAddress(sd.Address),
// 	}
// }

// func transformSellerAddress(sa *SellerAddress) *shipmentRequest.SellerAddress {
// 	if sa == nil {
// 		return nil
// 	}
// 	return &shipmentRequest.SellerAddress{
// 		AddressLine1: sa.AddressLine1,
// 		AddressLine2: sa.AddressLine2,
// 		Country:      transformSellerAddressCountry(sa.Country),
// 		State:        transformSellerAddressState(sa.State),
// 		City:         transformSellerAddressCity(sa.City),
// 		Pincode:      sa.Pincode,
// 	}
// }

// func transformSellerAddressCountry(sac *SellerAddressCountry) *shipmentRequest.SellerAddressCountry {
// 	if sac == nil {
// 		return nil
// 	}
// 	return &shipmentRequest.SellerAddressCountry{
// 		ID:   sac.ID,
// 		Name: sac.Name,
// 		Code: sac.Code,
// 	}
// }

// func transformSellerAddressState(sas *SellerAddressState) *shipmentRequest.SellerAddressState {
// 	if sas == nil {
// 		return nil
// 	}
// 	return &shipmentRequest.SellerAddressState{
// 		ID:   sas.ID,
// 		Name: sas.Name,
// 	}
// }

// func transformSellerAddressCity(sac *SellerAddressCity) *shipmentRequest.SellerAddressCity {
// 	if sac == nil {
// 		return nil
// 	}
// 	return &shipmentRequest.SellerAddressCity{
// 		ID:   sac.ID,
// 		Name: sac.Name,
// 	}
// }

// func transformVillageDetails(vd *VillageDetails) *shipmentRequest.VillageDetails {
// 	if vd == nil {
// 		return nil
// 	}
// 	return &shipmentRequest.VillageDetails{
// 		ID:         vd.ID,
// 		NameEn:     vd.NameEn,
// 		CityID:     vd.CityID,
// 		CityName:   vd.CityName,
// 		RegionID:   vd.RegionID,
// 		RegionName: vd.RegionName,
// 	}
// }

// func transformMobileNumber(mn *MobileNumber) *shipmentRequest.MobileNumber {
// 	if mn == nil {
// 		return nil
// 	}
// 	return &shipmentRequest.MobileNumber{
// 		Number:             mn.Number,
// 		CountryCallingCode: mn.CountryCallingCode,
// 		CountryCode:        mn.CountryCode,
// 	}
// }

// func transformReturnInfo(ri *ReturnInfo) *shipmentRequest.ReturnInfo {
// 	if ri == nil {
// 		return nil
// 	}
// 	return &shipmentRequest.ReturnInfo{
// 		PostalCode: ri.PostalCode,
// 		Address:    ri.Address,
// 		State:      ri.State,
// 		Phone:      ri.Phone,
// 		Name:       ri.Name,
// 		City:       ri.City,
// 		Country:    ri.Country,
// 	}
// }

// func transformResellerInfo(ri *ResellerInfo) *shipmentRequest.ResellerInfo {
// 	if ri == nil {
// 		return nil
// 	}
// 	return &shipmentRequest.ResellerInfo{
// 		Name:  ri.Name,
// 		Phone: ri.Phone,
// 	}
// }

// func transformCourierPartner(cp *CourierPartner) *shipmentRequest.CourierPartner {
// 	if cp == nil {
// 		return nil
// 	}
// 	return &shipmentRequest.CourierPartner{
// 		ID:   cp.ID,
// 		Name: cp.Name,
// 		Tag:  cp.Tag,
// 		Logo: cp.Logo,
// 	}
// }

// func transformDeliverySlot(ds *DeliverySlot) *shipmentRequest.DeliverySlot {
// 	if ds == nil {
// 		return nil
// 	}
// 	return &shipmentRequest.DeliverySlot{
// 		DeliveryDate: ds.DeliveryDate,
// 		StartTime:    ds.StartTime,
// 		EndTime:      ds.EndTime,
// 	}
// }

// func transformTags(tags []Tag) []shipmentRequest.Tag {
// 	if tags == nil {
// 		return nil
// 	}
// 	result := make([]shipmentRequest.Tag, len(tags))
// 	for i, tag := range tags {
// 		result[i] = shipmentRequest.Tag{
// 			ID:    tag.ID,
// 			Name:  tag.Name,
// 			Color: tag.Color,
// 		}
// 	}
// 	return result
// }

// func transformPriceSet(ps *PriceSet) *shipmentRequest.PriceSet {
// 	if ps == nil {
// 		return nil
// 	}
// 	return &shipmentRequest.PriceSet{
// 		OrderCurrency: shipmentRequest.Currency{
// 			Amount:       ps.OrderCurrency.Amount,
// 			CurrencyCode: ps.OrderCurrency.CurrencyCode,
// 		},
// 		StoreCurrency: shipmentRequest.Currency{
// 			Amount:       ps.StoreCurrency.Amount,
// 			CurrencyCode: ps.StoreCurrency.CurrencyCode,
// 		},
// 	}
// }

// func transformExchangeRate(er *ExchangeRate) *shipmentRequest.ExchangeRate {
// 	if er == nil {
// 		return nil
// 	}
// 	return &shipmentRequest.ExchangeRate{
// 		OrderCurrency: er.OrderCurrency,
// 		StoreCurrency: er.StoreCurrency,
// 		Rate:          er.Rate,
// 	}
// }

// func transformItems(items []Item) []shipmentRequest.Item {
// 	if items == nil {
// 		return nil
// 	}
// 	result := make([]shipmentRequest.Item, len(items))
// 	for i, item := range items {
// 		result[i] = shipmentRequest.Item{
// 			ProductURL:      item.ProductURL,
// 			Price:           item.Price,
// 			Description:     item.Description,
// 			Quantity:        item.Quantity,
// 			PackedQuantity:  item.PackedQuantity,
// 			SKU:             item.SKU,
// 			SKUName:         item.SKUName,
// 			CountryOfOrigin: item.CountryOfOrigin,
// 			Weight:          transformItemWeight(item.Weight),
// 			Additional:      transformItemAdditional(item.Additional),
// 		}
// 	}
// 	return result
// }

// func transformItemWeight(iw *ItemWeight) *shipmentRequest.ItemWeight {
// 	if iw == nil {
// 		return nil
// 	}
// 	return &shipmentRequest.ItemWeight{
// 		Value: iw.Value,
// 		UOM:   iw.UOM,
// 	}
// }

// func transformItemAdditional(ia *ItemAdditional) *shipmentRequest.ItemAdditional {
// 	if ia == nil {
// 		return nil
// 	}
// 	return &shipmentRequest.ItemAdditional{
// 		Height:     ia.Height,
// 		Length:     ia.Length,
// 		Breadth:    ia.Breadth,
// 		Weight:     ia.Weight,
// 		Images:     transformItemImages(ia.Images),
// 		ReturnDays: ia.ReturnDays,
// 	}
// }

// func transformItemImages(images []ItemImage) []shipmentRequest.ItemImage {
// 	if images == nil {
// 		return nil
// 	}
// 	result := make([]shipmentRequest.ItemImage, len(images))
// 	for i, img := range images {
// 		result[i] = shipmentRequest.ItemImage{
// 			URL:         img.URL,
// 			Type:        img.Type,
// 			Description: img.Description,
// 		}
// 	}
// 	return result
// }

// func transformPackageDetails(details []PackageDetail) []shipmentRequest.PackageDetail {
// 	if details == nil {
// 		return nil
// 	}
// 	result := make([]shipmentRequest.PackageDetail, len(details))
// 	for i, detail := range details {
// 		result[i] = shipmentRequest.PackageDetail{
// 			PackageID:        detail.PackageID,
// 			Dimensions:       transformDimensions(detail.Dimensions),
// 			Weight:           transformWeight(detail.Weight),
// 			InvoiceTotal:     detail.InvoiceTotal,
// 			InvoiceSubTotal:  detail.InvoiceSubTotal,
// 			InvoiceTotalPaid: detail.InvoiceTotalPaid,
// 			InvoiceTotalDue:  detail.InvoiceTotalDue,
// 			Items:            transformItems(detail.Items),
// 		}
// 	}
// 	return result
// }

// func transformDimensions(d Dimensions) shipmentRequest.Dimensions {
// 	return shipmentRequest.Dimensions{
// 		Length:  d.Length,
// 		Breadth: d.Breadth,
// 		Height:  d.Height,
// 		UOM:     d.UOM,
// 	}
// }

// func transformWeight(w Weight) shipmentRequest.Weight {
// 	return shipmentRequest.Weight{
// 		Value: w.Value,
// 		UOM:   w.UOM,
// 	}
// }

// func applyAddressToPickupDetails(pickupDetails *PickupDetails, address *SellerAddress) {
// 	if address == nil {
// 		return
// 	}

// 	pickupDetails.Address = address.AddressLine1
// 	if address.Country != nil {
// 		pickupDetails.CountryID = address.Country.ID
// 		pickupDetails.Country = address.Country.Name
// 		pickupDetails.CountryCode = address.Country.Code
// 	}
// 	pickupDetails.Pincode = address.Pincode
// 	if address.State != nil {
// 		pickupDetails.StateID = address.State.ID
// 		pickupDetails.State = address.State.Name
// 	}
// 	if address.City != nil {
// 		pickupDetails.CityID = address.City.ID
// 		pickupDetails.City = address.City.Name
// 	}
// }

// // Helper functions
// func isEnabled(config map[string]interface{}, key string) bool {
// 	if val, ok := config[key]; ok {
// 		return val != nil
// 	}
// 	return false
// }

// func parsePhoneNumber(phone string) string {
// 	return phone
// }

// func convertPriceSetToMap(priceSet *PriceSet) map[string]interface{} {
// 	if priceSet == nil {
// 		return make(map[string]interface{})
// 	}
// 	return map[string]interface{}{
// 		"order_currency": map[string]interface{}{
// 			"amount":        priceSet.OrderCurrency.Amount,
// 			"currency_code": priceSet.OrderCurrency.CurrencyCode,
// 		},
// 		"store_currency": map[string]interface{}{
// 			"amount":        priceSet.StoreCurrency.Amount,
// 			"currency_code": priceSet.StoreCurrency.CurrencyCode,
// 		},
// 	}
// }

// func convertExchangeRateToMap(exchangeRate *ExchangeRate) map[string]interface{} {
// 	if exchangeRate == nil {
// 		return make(map[string]interface{})
// 	}
// 	return map[string]interface{}{
// 		"order_currency": exchangeRate.OrderCurrency,
// 		"store_currency": exchangeRate.StoreCurrency,
// 		"rate":           exchangeRate.Rate,
// 	}
// }

// func convertCurrencyValues(conversion *CurrencyConversion) (map[string]float64, error) {
// 	// Implement actual currency conversion logic here
// 	// For now, returning original values
// 	return map[string]float64{
// 		"invoice_value": conversion.InvoiceValue,
// 		"order_value":   conversion.OrderValue,
// 		"total_paid":    conversion.TotalPaid,
// 		"total_due":     conversion.TotalDue,
// 	}, nil
// }

// type CurrencyConversion struct {
// 	InvoiceValue         float64                `json:"invoice_value"`
// 	OrderValue           float64                `json:"order_value"`
// 	TotalPaid            float64                `json:"total_paid"`
// 	TotalDue             float64                `json:"total_due"`
// 	Currency             string                 `json:"currency"`
// 	InvoiceValuePriceSet map[string]interface{} `json:"invoice_value_price_set"`
// 	OrderValuePriceSet   map[string]interface{} `json:"order_value_price_set"`
// 	TotalPaidPriceSet    map[string]interface{} `json:"total_paid_price_set"`
// 	TotalDuePriceSet     map[string]interface{} `json:"total_due_price_set"`
// 	ExchangeRate         map[string]interface{} `json:"exchange_rate"`
// 	DropCountryCurrency  string                 `json:"drop_country_currency"`
// }
