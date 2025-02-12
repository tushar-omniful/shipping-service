package errors

import "fmt"

const (
	OrderIdExists                 = "order_id_exists"
	ServiceTemporarilyUnavailable = "temporarily unavailable"
	Unauthorized                  = "unauthorized"
	NotFound                      = "not found"
)

type TokenError struct {
	Message string
	Code    string
}

type ReferenceExistsError struct {
	Message string
	Code    string
}

func (e *TokenError) Error() string {
	return fmt.Sprintf("token error: %s", e.Message)
}

func (e *ReferenceExistsError) Error() string {
	return fmt.Sprintf("reference exists error: %s", e.Message)
}
