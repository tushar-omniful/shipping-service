package order_service

import (
	"github.com/omniful/shipping-service/internal/domain/models"
	"github.com/omniful/shipping-service/internal/services/shipment/requests"
	"github.com/omniful/shipping-service/internal/services/shipment/responses"
)

func TransformCreateShipmentResponseToOrderModel(req *requests.CreateForwardShipmentRequest,
	shipmentResp *responses.CreateForwardShipmentResponse, op *models.OrderPartner, psm *models.PartnerShippingMethod,
) *models.Order {
	source := req.Data.OrderSource
	if source == "" {
		source = string(models.SourceOMS)
	}

	return &models.Order{
		OrderPartnerOrderID:     req.Data.OrderPartnerOrderID,
		OrderPartnerID:          op.ID,
		PartnerShippingMethodID: psm.ID,
		ShippingPartnerID:       psm.ShippingPartnerID,
		Status:                  models.OrderStatus(shipmentResp.Status),
		ShipmentType:            req.Data.ShipmentType,
		ShippingPartnerStatus:   shipmentResp.ShippingPartnerStatus,
		PickupDetails:           req.Data.PickupDetailsAttr.PickupDetails,
		DropDetails:             req.Data.DropDetailsAttr.DropDetails,
		ShipmentDetails:         req.Data.ShipmentDetailsAttr.ShipmentDetails,
		TaxDetails:              req.Data.TaxDetails,
		Metadata: &models.OrderMetadata{
			OmnifulAwbLabel:  shipmentResp.Label,
			ShippingAwbLabel: shipmentResp.Label,
			TrackingURL:      shipmentResp.Metadata.TrackingUrl,
			AwbNumber:        shipmentResp.AwbNumber,
			NumberOfRetries:  0, // Initialize with 0 retries
			// Copy other metadata from request if needed
			DeliveryType: req.Data.Metadata.DeliveryType,
			TaxNumber:    req.Data.Metadata.TaxNumber,
			ReturnInfo:   req.Data.Metadata.ReturnInfo,
			ResellerInfo: req.Data.Metadata.ResellerInfo,
		},
		ShippingLabel: shipmentResp.Label,
		AWBNumber:     shipmentResp.AwbNumber,
		SellerID:      req.SellerID,
		Source:        models.OrderSource(source),
	}
}
