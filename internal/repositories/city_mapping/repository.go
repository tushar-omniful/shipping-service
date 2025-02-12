package city_mapping_repo

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

// Repository handles queries for city mappings.
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

// CreateCityMapping creates a new city mapping record.
func (r *Repository) CreateCityMapping(ctx context.Context, cityMapping *models.CityMapping) (cusErr commonError.CustomError) {
	logTag := fmt.Sprintf("RequestID: %s Function: CreateCityMapping", env.GetRequestID(ctx))
	if resultErr := r.db.GetMasterDB(ctx).Create(cityMapping).Error; resultErr != nil {
		cusErr = commonError.NewCustomError(customError.SqlCreateError, fmt.Sprintf("%s unable to create", logTag))
		return
	}
	return
}

// GetCityMapping retrieves a single city mapping record based on the provided condition and scopes.
func (r *Repository) GetCityMapping(
	ctx context.Context,
	condition map[string]any,
	scopes ...func(db *gorm.DB) *gorm.DB,
) (cityMapping *models.CityMapping, cusErr commonError.CustomError) {
	var cityMappings []*models.CityMapping
	logTag := fmt.Sprintf("RequestID: %s Function: GetCityMapping", env.GetRequestID(ctx))
	if resultErr := r.db.GetMasterDB(ctx).Where(condition).Scopes(scopes...).Find(&cityMappings).Error; resultErr != nil {
		cusErr = commonError.NewCustomError(customError.SqlFetchError, fmt.Sprintf("%s unable to read", logTag))
		return
	}

	if len(cityMappings) == 0 {
		cusErr = commonError.NewCustomError(customError.NotFound, "CityMapping not found")
		return
	}

	cityMapping = cityMappings[0]
	return
}

// GetCityMappingByCityID retrieves a city mapping record by the given city ID.
func (r *Repository) GetCityMappingByCityID(
	ctx context.Context,
	cityID string,
	scopes ...func(db *gorm.DB) *gorm.DB,
) (cityMapping *models.CityMapping, cusErr commonError.CustomError) {
	conditions := map[string]any{
		"city_id": cityID,
	}

	cityMapping, cusErr = r.GetCityMapping(ctx, conditions, scopes...)
	if cusErr.Exists() {
		return
	}

	if cityMapping == nil {
		cusErr = commonError.NewCustomError(customError.NotFound, "CityMapping not found")
		return
	}
	return
}

// GetCityMappingByShippingPartnerAndOmnifulCityID retrieves a city mapping record based on the provided shipping partner ID and omniful city ID.
func (r *Repository) GetCityMappingByShippingPartnerAndOmnifulCityID(
	ctx context.Context,
	shippingPartnerID int64,
	omnifulCityID string,
	scopes ...func(db *gorm.DB) *gorm.DB,
) (cityMapping *models.CityMapping, cusErr commonError.CustomError) {
	conditions := map[string]any{
		"shipping_partner_id": shippingPartnerID,
		"omniful_city_id":     omnifulCityID,
	}

	cityMapping, cusErr = r.GetCityMapping(ctx, conditions, scopes...)
	if cusErr.Exists() {
		return
	}

	if cityMapping == nil {
		cusErr = commonError.NewCustomError(customError.NotFound, "CityMapping not found")
		return
	}

	return
}
