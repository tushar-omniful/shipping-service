package tenant_city_mapping_repo

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

// Repository handles queries for tenant city mappings.
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

// CreateTenantCityMapping creates a new tenant city mapping record.
func (r *Repository) CreateTenantCityMapping(ctx context.Context, tenantCityMapping *models.TenantCityMapping) (cusErr commonError.CustomError) {
	logTag := fmt.Sprintf("RequestID: %s Function: CreateTenantCityMapping", env.GetRequestID(ctx))
	if resultErr := r.db.GetMasterDB(ctx).Create(tenantCityMapping).Error; resultErr != nil {
		cusErr = commonError.NewCustomError(customError.SqlCreateError, fmt.Sprintf("%s unable to create", logTag))
		return
	}
	return
}

// GetTenantCityMapping retrieves a single tenant city mapping record based on the provided condition and scopes.
func (r *Repository) GetTenantCityMapping(
	ctx context.Context,
	condition map[string]any,
	scopes ...func(db *gorm.DB) *gorm.DB,
) (tenantCityMapping *models.TenantCityMapping, cusErr commonError.CustomError) {
	var tenantCityMappings []*models.TenantCityMapping
	logTag := fmt.Sprintf("RequestID: %s Function: GetTenantCityMapping", env.GetRequestID(ctx))
	if resultErr := r.db.GetMasterDB(ctx).Where(condition).Scopes(scopes...).Find(&tenantCityMappings).Error; resultErr != nil {
		cusErr = commonError.NewCustomError(customError.SqlFetchError, fmt.Sprintf("%s unable to read", logTag))
		return
	}

	if len(tenantCityMappings) == 0 {
		cusErr = commonError.NewCustomError(customError.NotFound, "TenantCityMapping not found")
		return
	}

	tenantCityMapping = tenantCityMappings[0]
	return
}

// GetTenantCityMappingByTenantID retrieves a tenant city mapping record by the given tenant ID.
func (r *Repository) GetTenantCityMappingByTenantID(
	ctx context.Context,
	tenantID string,
	scopes ...func(db *gorm.DB) *gorm.DB,
) (tenantCityMapping *models.TenantCityMapping, cusErr commonError.CustomError) {
	conditions := map[string]any{
		"tenant_id": tenantID,
	}

	tenantCityMapping, cusErr = r.GetTenantCityMapping(ctx, conditions, scopes...)
	if cusErr.Exists() {
		return
	}

	if tenantCityMapping == nil {
		cusErr = commonError.NewCustomError(customError.NotFound, "TenantCityMapping not found")
		return
	}
	return
}
