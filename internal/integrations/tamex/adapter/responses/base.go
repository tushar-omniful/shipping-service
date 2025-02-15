package responses

import (
	"fmt"
	"github.com/omniful/shipping-service/internal/domain/interfaces"

	customError "github.com/omniful/go_commons/error"
	shsError "github.com/omniful/shipping-service/pkg/error"
)

const (
	TamexError = "Error from Tamex:"
)

// ResponseAdapter holds common dependencies for response handling
type ResponseAdapter struct {
	opRepo    interfaces.OrderPartnerRepository
	orderRepo interfaces.OrderRepository
}

func NewResponseAdapter(opRepo interfaces.OrderPartnerRepository, orderRepo interfaces.OrderRepository) *ResponseAdapter {
	return &ResponseAdapter{
		opRepo:    opRepo,
		orderRepo: orderRepo,
	}
}

// CommonResponse shared by all Tamex responses
type CommonResponse struct {
	Data string `json:"data,omitempty"`
	Code int    `json:"code"`
}

func (r *CommonResponse) IsError() bool {
	return r.Code == 0
}

func (r *CommonResponse) ErrorMsg() string {
	return r.Data
}

// Common error raisers
func raiseBadRequest(message string) customError.CustomError {
	return customError.NewCustomError(shsError.BadRequest, fmt.Sprintf("%s %s", TamexError, message))
}

func serverDownError(message string) customError.CustomError {
	return customError.NewCustomError(shsError.BadRequest, fmt.Sprintf("%s Service Not Available", TamexError))
}
