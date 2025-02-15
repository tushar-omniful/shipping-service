package responses

import (
	"fmt"
	customError "github.com/omniful/go_commons/error"
	"github.com/omniful/shipping-service/internal/services/shipment/responses"
	shsError "github.com/omniful/shipping-service/pkg/error"
	"strings"
)

const (
	OrderIdExists = "order_id_exists"
)

type CreateResponse struct {
	CommonResponse
	OrderID       string `json:"orderId"`
	WayBillNumber string `json:"waybillNumber"`
	LabelFormat   string `json:"labelFormat"`
}

func (a *ResponseAdapter) ParseCreateShipment(respData CreateResponse) (resp responses.CreateForwardShipmentResponse, err customError.CustomError) {

	if respData.ResponseCode != "100" || respData.ResponseMsg != "Success" {
		err = customError.RequestInvalidError(fmt.Sprintf("Error from Ajex: %v", respData.ResponseMsg))
		return
	}

	if respData.ResponseCode != "200" {
		return handleError(respData)
	}

	resp = responses.CreateForwardShipmentResponse{
		AwbNumber: respData.WayBillNumber,
	}

	return
}

func handleError(respData CreateResponse) (resp responses.CreateForwardShipmentResponse, err customError.CustomError) {
	if strings.Contains(strings.ToLower(respData.ResponseMsg), OrderIdExists) {
		err = customError.NewCustomError(shsError.OrderReferenceExistError, respData.ResponseMsg)
	}

	return responses.CreateForwardShipmentResponse{}, customError.RequestInvalidError(respData.ResponseMsg)
}
