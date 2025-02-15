package responses

import (
	"fmt"

	customError "github.com/omniful/go_commons/error"
	"github.com/omniful/shipping-service/internal/domain/interfaces"
	shsError "github.com/omniful/shipping-service/pkg/error"
	"gopkg.in/guregu/null.v4"
)

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

type LoginResponse struct {
	Status      null.Bool `json:"status"`
	Response    any       `json:"response"`
	AccessToken any       `json:"access_token"`
	TokenType   any       `json:"token_type"`
}

type CommonResponse struct {
	IsSuccess    bool      `json:"is_success"`
	StatusCode   int       `json:"status_code"`
	Status       null.Bool `json:"status"`
	ResponseCode string    `json:"responseCode"`
	ResponseMsg  string    `json:"responseMessage"`
}

func (r *CommonResponse) IsError() bool {
	if r.ResponseCode != "100" || r.ResponseMsg != "Success" {
		return false
	}
	return true
}

func (r *CommonResponse) ErrorMsg() string {
	return r.ResponseMsg
}

func (r *CommonResponse) IsUnauthorized() bool {
	if r == nil {
		return false
	}

	return r.StatusCode == 401
}

func raiseBadRequest(message string) customError.CustomError {
	return customError.NewCustomError(shsError.BadRequest, fmt.Sprintf("Error from Ajex: %s", message))
}

func raiseOrderReferenceExist(message string) customError.CustomError {
	return customError.NewCustomError(shsError.OrderReferenceExistError, message)
}
