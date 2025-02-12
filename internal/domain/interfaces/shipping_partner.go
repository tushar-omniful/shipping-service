package interfaces

import (
	"context"
	oerror "github.com/omniful/go_commons/error"
	"github.com/omniful/shipping-service/internal/domain/models"
	"gorm.io/gorm"
)

type ShippingPartnerController interface {
	//GetShippingPartner(ctx *gin.Context)
	//CreateShippingPartner(ctx *gin.Context)
	//UpdateShippingPartner(ctx *gin.Context)
}

type ShippingPartnerService interface {
	//GetShippingPartner(ctx context.Context, id int64) (*models.ShippingPartner, oerror.CustomError)
	//CreateShippingPartner(ctx context.Context, partner *models.ShippingPartner) oerror.CustomError
	//UpdateShippingPartner(ctx context.Context, condition map[string]interface{}, partner *models.ShippingPartner) oerror.CustomError
}

type ShippingPartnerRepository interface {
	GetShippingPartnerByTag(ctx context.Context, tag string, scopes ...func(db *gorm.DB) *gorm.DB) (*models.ShippingPartner, oerror.CustomError)
	//GetShippingPartner(ctx context.Context, condition map[string]interface{}, scopes ...func(db *gorm.DB) *gorm.DB) (*models.ShippingPartner, oerror.CustomError)
	//CreateShippingPartner(ctx context.Context, partner *models.ShippingPartner) oerror.CustomError
	//UpdateShippingPartner(ctx context.Context, condition map[string]interface{}, partner *models.ShippingPartner) oerror.CustomError
}
