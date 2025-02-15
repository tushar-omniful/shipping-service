package requests

import (
	"fmt"
	customError "github.com/omniful/go_commons/error"
	"strings"
	"time"

	"github.com/omniful/shipping-service/internal/domain/models"
	"github.com/omniful/shipping-service/internal/services/shipment/requests"
)

const (
	CODServiceCode        = "COD"
	AddressTypeFreeText   = "FREE_TEXT"
	ContactTypeEnterprise = "ENTERPRISE"
	ContactTypeIndividual = "INDIVIDUAL"
)

var CountriesWithoutPostalCode = map[string]bool{
	"AE": true,
	"SA": true,
}

func (a *RequestAdapter) FormatCreateShipment(req *requests.CreateForwardShipmentRequest) (requestData map[string]interface{}, err customError.CustomError) {
	details := req.PartnerShippingMethod.Details
	requestData = map[string]interface{}{
		"orderId":            formatOrderID(req.Data.OrderPartnerOrderID, req.PartnerShippingMethod.ShippingPartnerID),
		"orderTime":          time.Now().UTC().Format("2006-01-02T15:04:05.000Z"),
		"productCode":        details["product_code"].(string),
		"expressType":        details["express_type"].(string),
		"totalDeclaredValue": getInvoiceOrOrderValue(req.Data.ShipmentDetailsAttr.ShipmentDetails),
		"declaredCurrency":   req.Data.ShipmentDetailsAttr.ShipmentDetails.Currency,
		"parcelTotalWeight":  req.Data.ShipmentDetailsAttr.ShipmentDetails.Weight,
		"pickupMethod":       details["pickup_method"].(string),
		"paymentMethod":      details["payment_method"].(string),
		"customerAccount":    details["customer_account"].(string),
		"buyerName":          req.Data.DropDetailsAttr.DropDetails.Name,
		"senderInfo":         formatSenderInfo(&req.Data.PickupDetailsAttr, details),
		"receiverInfo":       formatReceiverInfo(req.Data.DropDetailsAttr.DropDetails),
		"parcels":            formatParcels(req.Data.ShipmentDetailsAttr.ShipmentDetails, details),
		"orderAlias":         req.Data.OrderAlias,
	}

	// Add COD service if total due is greater than 0
	if req.Data.ShipmentDetailsAttr.ShipmentDetails.TotalDue > 0 {
		requestData["addedServices"] = []map[string]interface{}{{
			"serviceName": CODServiceCode,
			"val1":        req.Data.ShipmentDetailsAttr.ShipmentDetails.TotalDue,
			"val2":        req.Data.ShipmentDetailsAttr.ShipmentDetails.Currency,
		}}
	}

	return
}

func formatOrderID(orderID string, partnerID uint64) string {
	reference := fmt.Sprintf("%s-%d", orderID, partnerID)
	return strings.ReplaceAll(reference, " ", "-")
}

func formatSenderInfo(details *requests.PickupDetailsAttrs, partnerDetails map[string]interface{}) map[string]interface{} {
	addressType := AddressTypeFreeText

	if v, ok := partnerDetails["default_address_type"].(string); ok && v != "" {
		addressType = v
	}

	info := map[string]interface{}{
		"name":            details.PickupDetails.Name,
		"phone":           formatPhone(details.PickupDetails.Phone),
		"email":           details.PickupDetails.Email,
		"contactType":     ContactTypeEnterprise,
		"addressType":     addressType,
		"country":         details.PickupDetails.Country,
		"countryCode":     details.PickupDetails.CountryCode,
		"province":        details.PickupDetails.State,
		"city":            details.PickupDetails.City,
		"detailedAddress": details.PickupDetails.Address,
		"longitude":       details.PickupDetails.Longitude,
		"latitude":        details.PickupDetails.Latitude,
	}

	if !CountriesWithoutPostalCode[details.PickupDetails.CountryCode] {
		info["postalCode"] = details.PickupDetails.Pincode
	}

	return info
}

func formatReceiverInfo(details *models.DropDetails) map[string]interface{} {
	info := map[string]interface{}{
		"name":            details.Name,
		"phone":           formatPhone(details.Phone),
		"email":           details.Email,
		"contactType":     ContactTypeIndividual,
		"addressType":     AddressTypeFreeText,
		"country":         details.OmnifulCountry,
		"countryCode":     details.CountryCode,
		"province":        details.OmnifulState,
		"city":            details.OmnifulCity,
		"detailedAddress": details.Address,
		"longitude":       details.Longitude,
		"latitude":        details.Latitude,
	}

	if !CountriesWithoutPostalCode[details.CountryCode] {
		info["postalCode"] = details.Pincode
	}

	return info
}

func formatParcels(details *models.ShipmentDetails, partnerDetails map[string]interface{}) []map[string]interface{} {
	parcels := make([]map[string]interface{}, len(details.PackageDetails))
	for i, pkg := range details.PackageDetails {
		parcels[i] = map[string]interface{}{
			"weight":    pkg.Weight.Value,
			"quantity":  pkg.Items[0].Quantity,
			"cargoInfo": formatCargoInfo(pkg.Items, details.Currency, isInternationalShipment(details), partnerDetails),
		}
	}
	return parcels
}

func formatCargoInfo(items []models.Item, currency string, isInternational bool, partnerDetails map[string]interface{}) []map[string]interface{} {
	cargoInfo := make([]map[string]interface{}, len(items))
	for i, item := range items {
		name := item.SKUName
		if name == "" {
			name = item.SKU
		}
		if len(name) > 200 {
			name = name[:200]
		}

		info := map[string]interface{}{
			"name":  name,
			"count": item.Quantity,
		}

		if isInternational {
			info["totalValue"] = float64(item.Price)
			info["hsCode"] = partnerDetails["hs_code"]
			info["countryOfOrigin"] = partnerDetails["country_of_origin"]
		}

		cargoInfo[i] = info
	}
	return cargoInfo
}

func formatPhone(phone string) string {
	if strings.Contains(phone, "-") {
		parts := strings.Split(phone, "-")
		if len(parts) > 1 {
			return parts[1]
		}
	}
	return phone
}

func getInvoiceOrOrderValue(details *models.ShipmentDetails) float64 {
	if details.InvoiceValue > 0 {
		return details.InvoiceValue
	}
	return details.OrderValue
}

func isInternationalShipment(details *models.ShipmentDetails) bool {
	// This would need to be implemented based on your business logic
	// for determining international shipments
	return false
}
