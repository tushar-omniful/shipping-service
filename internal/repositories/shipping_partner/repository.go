package shipping_partner_repo

import (
	"context"
	"fmt"
	"sync"

	"github.com/omniful/go_commons/db/sql/postgres"
	"github.com/omniful/go_commons/env"
	commonError "github.com/omniful/go_commons/error"
	"github.com/omniful/shipping-service/internal/domain/models"
	customError "github.com/omniful/shipping-service/pkg/error"
	"gorm.io/gorm"
)

// Repository handles queries for shipping partners.
type Repository struct {
	db *postgres.DbCluster
}

var (
	repository     *Repository
	repositoryOnce sync.Once
)

// NewRepository initializes and returns a singleton instance of Repository.
func NewRepository(db *postgres.DbCluster) *Repository {
	repositoryOnce.Do(func() {
		repository = &Repository{db: db}
	})
	return repository
}

// CreateShippingPartner creates a new shipping partner record.
func (r *Repository) CreateShippingPartner(ctx context.Context, shippingPartner *models.ShippingPartner) (cusErr commonError.CustomError) {
	logTag := fmt.Sprintf("RequestID: %s Function: CreateShippingPartner", env.GetRequestID(ctx))
	if resultErr := r.db.GetMasterDB(ctx).Create(shippingPartner).Error; resultErr != nil {
		cusErr = commonError.NewCustomError(customError.SqlCreateError, fmt.Sprintf("%s unable to create", logTag))
		return
	}
	return
}

// GetShippingPartner retrieves a single shipping partner record based on the provided condition and scopes.
func (r *Repository) GetShippingPartner(
	ctx context.Context,
	condition map[string]any,
	scopes ...func(db *gorm.DB) *gorm.DB,
) (shippingPartner *models.ShippingPartner, cusErr commonError.CustomError) {
	var shippingPartners []*models.ShippingPartner
	logTag := fmt.Sprintf("RequestID: %s Function: GetShippingPartner", env.GetRequestID(ctx))
	if resultErr := r.db.GetMasterDB(ctx).Where(condition).Scopes(scopes...).Find(&shippingPartners).Error; resultErr != nil {
		cusErr = commonError.NewCustomError(customError.SqlFetchError, fmt.Sprintf("%s unable to read", logTag))
		return
	}

	if len(shippingPartners) == 0 {
		cusErr = commonError.NewCustomError(customError.NotFound, "ShippingPartner not found")
		return
	}

	shippingPartner = shippingPartners[0]
	return
}

// GetShippingPartnerByTenantID retrieves a shipping partner record by the given tenant ID.
func (r *Repository) GetShippingPartnerByTenantID(
	ctx context.Context,
	tenantID string,
	scopes ...func(db *gorm.DB) *gorm.DB,
) (shippingPartner *models.ShippingPartner, cusErr commonError.CustomError) {
	conditions := map[string]any{
		"tenant_id": tenantID,
	}

	shippingPartner, cusErr = r.GetShippingPartner(ctx, conditions, scopes...)
	if cusErr.Exists() {
		return
	}

	if shippingPartner == nil {
		cusErr = commonError.NewCustomError(customError.NotFound, "ShippingPartner not found")
		return
	}
	return
}

// GetShippingPartnerByTag retrieves a shipping partner record by the given tag.
func (r *Repository) GetShippingPartnerByTag(
	ctx context.Context,
	tag string,
	scopes ...func(db *gorm.DB) *gorm.DB,
) (shippingPartner *models.ShippingPartner, cusErr commonError.CustomError) {
	conditions := map[string]any{
		"tag": tag,
	}

	shippingPartner, cusErr = r.GetShippingPartner(ctx, conditions, scopes...)
	if cusErr.Exists() {
		return
	}

	if shippingPartner == nil {
		cusErr = commonError.NewCustomError(customError.NotFound, fmt.Sprintf("ShippingPartner not found for tag %s", tag))
		return
	}
	return
}
