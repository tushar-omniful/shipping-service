package requests

import (
	"fmt"
	"strings"

	customError "github.com/omniful/go_commons/error"

	"github.com/omniful/shipping-service/internal/domain/models"
	"github.com/omniful/shipping-service/internal/services/shipment/requests"
)

const (
	PackTypeDelivery = "1" // 1 for delivery, 2 for pickup
)

func (a *RequestAdapter) FormatCreateShipment(req *requests.CreateForwardShipmentRequest) (requestData map[string]interface{}, err customError.CustomError) {
	credentials := req.PartnerShippingMethod.Credentials
	apiKey, ok := credentials["api_key"].(string)
	if !ok {
		return nil, customError.NewCustomError("invalid_credentials", "Missing API key in credentials")
	}

	shipmentDetails := req.Data.ShipmentDetailsAttr.ShipmentDetails
	pickupDetails := req.Data.PickupDetailsAttr.PickupDetails
	dropDetails := req.Data.DropDetailsAttr.DropDetails

	requestData = map[string]interface{}{
		"apikey":                 apiKey,
		"pack_type":              PackTypeDelivery,
		"pack_awb":               strings.ToUpper(formatOrderID(req.Data.OrderPartnerOrderID, req.PartnerShippingMethod.ShippingPartnerID)),
		"pack_reciver_name":      dropDetails.Name,
		"pack_reciver_phone":     formatPhone(dropDetails.Phone),
		"pack_reciver_country":   dropDetails.CountryCode,
		"pack_reciver_city":      dropDetails.OmnifulCity,
		"pack_reciver_dist":      getValueOrDefault(dropDetails.OmnifulState, dropDetails.OmnifulCity),
		"pack_reciver_email":     dropDetails.Email,
		"pack_reciver_street":    dropDetails.Address,
		"pack_reciver_zipcode":   dropDetails.Pincode,
		"pack_reciver_longitude": dropDetails.Longitude,
		"pack_reciver_latitude":  dropDetails.Latitude,
		"pack_desc":              formatDescription(shipmentDetails),
		"pack_num_pcs":           shipmentDetails.NumberOfBoxes,
		"pack_weight":            shipmentDetails.Weight,
		"pack_cod_amount":        getCODAmount(shipmentDetails),
		"pack_currency_code":     getCurrencyCode(shipmentDetails, dropDetails),
		"pack_extra_note":        getValueOrDefault(shipmentDetails.Remarks, ""),
		"pack_vendor_id":         pickupDetails.Name,
		"pack_sender_name":       pickupDetails.Name,
		"pack_sender_phone":      formatPhone(pickupDetails.Phone),
		"pack_sender_email":      pickupDetails.Email,
		"pack_send_country":      pickupDetails.CountryCode,
		"pack_send_city":         pickupDetails.City,
		"pack_sender_dist":       getSenderDistrict(pickupDetails),
		"pack_sender_street":     pickupDetails.Address,
		"pack_sender_zipcode":    pickupDetails.Pincode,
		"pack_sender_longitude":  pickupDetails.Longitude,
		"pack_sender_latitude":   pickupDetails.Latitude,
	}

	return
}

func formatDescription(details *models.ShipmentDetails) string {
	if len(details.Items) == 0 {
		return ""
	}

	var desc string
	for _, item := range details.Items {
		if desc != "" {
			desc += ", "
		}
		desc += fmt.Sprintf("%s (x%d)", item.SKUName, item.Quantity)
	}

	if len(desc) > 250 {
		desc = desc[:250]
	}

	return desc
}

func getCODAmount(details *models.ShipmentDetails) float64 {
	if details.TotalDue > 0 {
		return details.TotalDue
	}
	return 0
}

func getCurrencyCode(shipmentDetails *models.ShipmentDetails, dropDetails *models.DropDetails) string {
	if shipmentDetails.TotalDue == 0 {
		return "SAR"
	}

	if shipmentDetails.Currency == "SAR" {
		return "SAR"
	}

	if dropDetails.CountryCurrencyCode == "SAR" {
		if shipmentDetails.TotalDuePriceSet != nil {
			if shipmentDetails.TotalDuePriceSet.OrderCurrency.CurrencyCode == "SAR" ||
				shipmentDetails.TotalDuePriceSet.StoreCurrency.CurrencyCode == "SAR" {
				return "SAR"
			}
		}
	}

	return "SAR" // TAMEX only supports SAR for COD
}

func getSenderDistrict(details *models.PickupDetails) string {
	if details.State != "" {
		return details.State
	}
	if details.District != "" {
		return details.District
	}
	return details.City
}

func getValueOrDefault(value, defaultValue string) string {
	if value != "" {
		return value
	}
	return defaultValue
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

func formatOrderID(orderID string, partnerID uint64) string {
	reference := fmt.Sprintf("%s-%d", orderID, partnerID)
	return strings.ReplaceAll(reference, " ", "-")
}
