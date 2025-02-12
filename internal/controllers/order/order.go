package order_controller

import (
	"sync"

	"github.com/gin-gonic/gin"
	oerror "github.com/omniful/go_commons/error"
	"github.com/omniful/go_commons/response"
	ctrlRequest "github.com/omniful/shipping-service/internal/controllers/order/requests"
	"github.com/omniful/shipping-service/internal/domain/interfaces"
	customError "github.com/omniful/shipping-service/pkg/error"
	"github.com/omniful/shipping-service/pkg/http_response"
	validator "github.com/omniful/shipping-service/pkg/validate"
)

type Controller struct {
	service interfaces.OrderService
}

var ctrl *Controller
var ctrlOnce sync.Once

func NewController(service interfaces.OrderService) *Controller {
	ctrlOnce.Do(func() {
		ctrl = &Controller{
			service: service,
		}
	})

	return ctrl
}

func (sc *Controller) CreateOrder(ctx *gin.Context) {
	var createOrderReq ctrlRequest.CreateForwardOrder

	err := ctx.ShouldBindJSON(&createOrderReq)
	if err != nil {
		cusErr := oerror.NewCustomError(customError.BadRequest, err.Error())
		customError.NewErrorResponse(ctx, cusErr)
		return
	}

	//tenantID, entityID, entityType, cusErr := validatePathParams(ctx)
	//if cusErr.Exists() {
	//	customError.NewErrorResponse(ctx, cusErr)
	//	return
	//}

	validationErr := validator.Get().Struct(createOrderReq)
	if validationErr != nil {
		cusErr := oerror.NewCustomError(customError.BadRequest, validationErr.Error())
		customError.NewErrorResponse(ctx, cusErr)
		return
	}

	// Transform controller request to service request using model types
	serviceRequest := createOrderReq.ToServiceRequest()

	res, cusErr := sc.service.CreateOrder(ctx, serviceRequest)
	if cusErr.Exists() {
		customError.NewErrorResponse(ctx, cusErr)
		return
	}

	response.NewSuccessResponse(ctx, http_response.Success(res))
}
