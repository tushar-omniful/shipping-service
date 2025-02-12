package http_response

import (
	"github.com/gin-gonic/gin"
	"github.com/omniful/shipping-service/constants"
	"math"
	"net/http"
)

type SuccessResponse struct {
	IsSuccess  bool        `json:"is_success"`
	StatusCode int         `json:"status_code"`
	Data       interface{} `json:"data"`
	Meta       interface{} `json:"meta"`
}

type Meta struct {
	CurrentPage int64  `json:"current_page"`
	PerPage     int64  `json:"per_page"`
	LastPage    int64  `json:"last_page"`
	Total       uint64 `json:"total"`
}

type ErrorResponse struct {
	IsSuccess  bool        `json:"is_success"`
	StatusCode int         `json:"status_code"`
	Error      ErrorStruct `json:"wms_error"`
}

type ErrorStruct struct {
	Message string   `json:"message"`
	Errors  []string `json:"errors,omitempty"`
}

func Metadata(ctx *gin.Context, total uint64) *Meta {
	limit := ctx.MustGet(constants.Limit).(int64)
	page := ctx.MustGet(constants.Page).(int64)
	lastPage := int64(math.Ceil(float64(total) / float64(limit)))
	if total == 0 {
		page = 0
	}
	result := Meta{
		CurrentPage: page,
		LastPage:    lastPage,
		PerPage:     limit,
		Total:       total,
	}
	return &result
}

func RaiseBadRequest(errorMessage string) ErrorResponse {
	return ErrorResponse{IsSuccess: false, StatusCode: http.StatusBadRequest, Error: ErrorStruct{Message: errorMessage}}
}

func RaiseBadRequestErrors(errorMessage string, errors []string) ErrorResponse {
	return ErrorResponse{IsSuccess: false, StatusCode: http.StatusBadRequest, Error: ErrorStruct{Message: errorMessage, Errors: errors}}
}

func RaiseNotFound(errorMessage string) ErrorResponse {
	return ErrorResponse{IsSuccess: false, StatusCode: http.StatusNotFound, Error: ErrorStruct{Message: errorMessage}}
}

func RaiseUnauthorized() ErrorResponse {
	return ErrorResponse{IsSuccess: false, StatusCode: http.StatusUnauthorized, Error: ErrorStruct{Message: "Unauthorized"}}
}

func Success(data interface{}) SuccessResponse {
	return SuccessResponse{IsSuccess: true, StatusCode: http.StatusOK, Data: data}
}

func SuccessGet(data interface{}, meta map[string]interface{}) SuccessResponse {
	value, ok := meta["meta"]
	if ok {
		return SuccessResponse{IsSuccess: true, StatusCode: http.StatusOK, Data: data, Meta: value}
	} else {
		return SuccessResponse{IsSuccess: true, StatusCode: http.StatusOK, Data: data}
	}
}

type StatusMessage struct {
	Message string `json:"message,omitempty"`
}

func SuccessMessage(message string) StatusMessage {
	return StatusMessage{
		Message: message,
	}
}
