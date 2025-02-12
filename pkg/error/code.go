package ss_error

import (
	oerror "github.com/omniful/go_commons/error"
	"github.com/omniful/go_commons/http"
)

var CustomCodeToHttpCodeMapping = map[oerror.Code]http.StatusCode{
	NotFound:        http.StatusBadRequest,
	RequestNotValid: http.StatusForbidden,
	BadRequest:      http.StatusBadRequest,
	RequestInvalid:  http.StatusBadRequest,
}
