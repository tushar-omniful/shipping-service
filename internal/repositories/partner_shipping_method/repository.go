package partner_shipping_method_repo

import (
	"context"
	"fmt"
	"sync"

	"github.com/omniful/go_commons/db/sql/postgres"
	"github.com/omniful/go_commons/env"
	commonError "github.com/omniful/go_commons/error"
	"github.com/omniful/shipping-service/internal/db"
	"github.com/omniful/shipping-service/internal/domain/models"
	customError "github.com/omniful/shipping-service/pkg/error"
	"gorm.io/gorm"
)

// Repository handles queries for partner shipping methods.
// You can use helper functions defined in internal/db/sort.go and internal/db/search.go
// (which use constant keys defined in constants/pagination.go) to generate sorting and searching scopes.
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
		repository = &Repository{
			db: db,
		}
	})
	return repository
}

// CreatePartnerShippingMethod creates a new partner shipping method record.
func (r *Repository) CreatePartnerShippingMethod(ctx context.Context, partnerShippingMethod *models.PartnerShippingMethod) (cusErr commonError.CustomError) {
	logTag := fmt.Sprintf("RequestID: %s Function: CreatePartnerShippingMethod", env.GetRequestID(ctx))
	if resultErr := r.db.GetMasterDB(ctx).Create(partnerShippingMethod).Error; resultErr != nil {
		cusErr = commonError.NewCustomError(customError.SqlCreateError, fmt.Sprintf("%s unable to create", logTag))
		return
	}
	return
}

// UpdatePartnerShippingMethod updates an existing partner shipping method record.
func (r *Repository) UpdatePartnerShippingMethod(ctx context.Context, condition map[string]interface{}, partnerShippingMethod *models.PartnerShippingMethod) (cusErr commonError.CustomError) {
	logTag := fmt.Sprintf("RequestID: %s Function: UpdatePartnerShippingMethod", env.GetRequestID(ctx))
	if resultErr := r.db.GetMasterDB(ctx).Model(partnerShippingMethod).Where(condition).Updates(partnerShippingMethod).Error; resultErr != nil {
		cusErr = commonError.NewCustomError(customError.SqlUpdateError, fmt.Sprintf("%s unable to update", logTag))
		return
	}
	return
}

// GetPaginatedPartnerShippingMethods retrieves partner shipping methods in a paginated format.
// The caller is expected to build any necessary scopes for sorting and searching using
// functions like db.GetSortingScopes(ctx, sortMap) from internal/db/sort.go and
// db.GetSearchParamsScopes(ctx, searchColumnMap) from internal/db/search.go.
func (r *Repository) GetPaginatedPartnerShippingMethods(
	ctx context.Context,
	condition map[string]any,
	countScopes []func(db *gorm.DB) *gorm.DB,
	scopes ...func(db *gorm.DB) *gorm.DB,
) (partnerShippingMethods models.PartnerShippingMethod, count int64, cusErr commonError.CustomError) {
	var psm []*models.PartnerShippingMethod
	psm, count, cusErr = db.GetPaginatedDataT[models.PartnerShippingMethod](
		ctx,
		r.db.GetMasterDB(ctx),
		models.PartnerShippingMethod{},
		condition,
		countScopes,
		scopes...,
	)
	if cusErr.Exists() {
		return
	}

	if len(psm) > 0 {
		partnerShippingMethods = *psm[0]
	}
	return
}

// GetPartnerShippingMethods retrieves a list of partner shipping method records based on the provided condition and scopes.
func (r *Repository) GetPartnerShippingMethods(
	ctx context.Context,
	condition map[string]any,
	scopes ...func(db *gorm.DB) *gorm.DB,
) (partnerShippingMethods []*models.PartnerShippingMethod, cusErr commonError.CustomError) {
	logTag := fmt.Sprintf("RequestID: %s Function: GetPartnerShippingMethods", env.GetRequestID(ctx))
	if resultErr := r.db.GetMasterDB(ctx).Where(condition).Scopes(scopes...).Find(&partnerShippingMethods).Error; resultErr != nil {
		cusErr = commonError.NewCustomError(customError.SqlFetchError, fmt.Sprintf("%s unable to read", logTag))
		return
	}

	return
}

// GetPartnerShippingMethod retrieves a single partner shipping method record based on the provided condition and scopes.
// It returns a NotFound error if no matching record is found.
func (r *Repository) GetPartnerShippingMethod(
	ctx context.Context,
	condition map[string]any,
	scopes ...func(db *gorm.DB) *gorm.DB,
) (partnerShippingMethod *models.PartnerShippingMethod, cusErr commonError.CustomError) {
	partnerShippingMethods, cusErr := r.GetPartnerShippingMethods(ctx, condition, scopes...)
	if cusErr.Exists() {
		return
	}

	if len(partnerShippingMethods) == 0 {
		cusErr = commonError.NewCustomError(customError.NotFound, "Account not found")
		return
	}

	partnerShippingMethod = partnerShippingMethods[0]
	return
}

func (r *Repository) GetPartnerShippingMethodByID(ctx context.Context, id string, orderPartnerID string) (partnerShippingMethod *models.PartnerShippingMethod, cusErr commonError.CustomError) {
	// Include both "id" and "order_partner_id" in the conditions.
	conditions := map[string]any{
		"id":               id,
		"order_partner_id": orderPartnerID,
	}

	partnerShippingMethod, cusErr = r.GetPartnerShippingMethod(ctx, conditions)
	if cusErr.Exists() {
		return
	}

	if partnerShippingMethod == nil {
		cusErr = commonError.NewCustomError(customError.NotFound, "PartnerShippingMethod not found")
		return
	}

	return
}
