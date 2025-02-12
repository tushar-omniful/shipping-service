package interfaces

import (
	"context"

	oerror "github.com/omniful/go_commons/error"
	"github.com/omniful/shipping-service/internal/domain/models"
)

type OrderPartnerController interface{}

type OrderPartnerService interface{}

type OrderPartnerRepository interface {
	GetOrderPartnerByTenantID(ctx context.Context, tenantID string) (*models.OrderPartner, oerror.CustomError)
}
