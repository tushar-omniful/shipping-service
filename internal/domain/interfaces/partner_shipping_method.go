package interfaces

import (
	"context"

	oerror "github.com/omniful/go_commons/error"
	"github.com/omniful/shipping-service/internal/domain/models"
)

type PartnerShippingMethodController interface {
	//GetPartnerShippingMethod(ctx *gin.Context)
	//CreatePartnerShippingMethod(ctx *gin.Context)
	//UpdatePartnerShippingMethod(ctx *gin.Context)
}

type PartnerShippingMethodService interface {
	//GetPartnerShippingMethod(ctx context.Context, id int64, orderPartnerID int64) (*models.PartnerShippingMethod, oerror.CustomError)
	//CreatePartnerShippingMethod(ctx context.Context, psm *models.PartnerShippingMethod) oerror.CustomError
	//UpdatePartnerShippingMethod(ctx context.Context, condition map[string]interface{}, psm *models.PartnerShippingMethod) oerror.CustomError
}

type PartnerShippingMethodRepository interface {
	GetPartnerShippingMethodByID(ctx context.Context, id string, orderPartnerID string) (*models.PartnerShippingMethod, oerror.CustomError)
	UpdatePartnerShippingMethod(ctx context.Context, condition map[string]interface{}, psm *models.PartnerShippingMethod) oerror.CustomError
}
