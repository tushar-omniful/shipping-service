package ss_error

import (
	"context"
	"github.com/gin-gonic/gin"
	oerror "github.com/omniful/go_commons/error"
	"github.com/omniful/go_commons/i18n"
	"github.com/omniful/go_commons/response"
)

const (
	BadRequest             oerror.Code = "BAD_REQUEST"
	NotFound               oerror.Code = "NOT_FOUND"
	RequestNotValid        oerror.Code = "REQUEST_NOT_VALID"
	RequestInvalid         oerror.Code = "REQUEST_INVALID"
	FailedToCreateShipment oerror.Code = "FAILED_TO_CREATE_SHIPMENT"
	SqlCreateError         oerror.Code = "SQL_CREATE_ERROR"
	SqlUpdateError         oerror.Code = "SQL_UPDATE_ERROR"
	SqlFetchError          oerror.Code = "SQL_FETCH_ERROR"
	SqlDeleteError         oerror.Code = "SQL_DELETE_ERROR"
)

func NewErrorResponse(ctx *gin.Context, cusErr oerror.CustomError) {
	response.NewErrorResponse(ctx, cusErr, CustomCodeToHttpCodeMapping)
	return
}

func NewReDirectResponse(ctx *gin.Context, redirectData response.Redirect) {
	response.NewRedirectResponse(ctx, redirectData)
	return
}

func InvalidRequest(ctx context.Context, key string) oerror.CustomError {
	message := i18n.Translate(ctx, key)
	return oerror.NewCustomError(RequestInvalid, message)
}
