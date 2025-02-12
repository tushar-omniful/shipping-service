package order_partner_repo

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

// Repository handles queries for order partners.
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

// CreateOrderPartner creates a new order partner record.
//func (r *Repository) CreateOrderPartner(ctx context.Context, orderPartner *models.OrderPartner) (cusErr commonError.CustomError) {
//	logTag := fmt.Sprintf("RequestID: %s Function: CreateOrderPartner", env.GetRequestID(ctx))
//	if resultErr := r.db.GetMasterDB(ctx).Create(orderPartner).Error; resultErr != nil {
//		cusErr = commonError.NewCustomError(customError.SqlCreateError, fmt.Sprintf("%s unable to create", logTag))
//		return
//	}
//	return
//}

// GetOrderPartner retrieves a single order partner record based on the provided condition and scopes.
// It returns a NotFound error if no matching record is found.
func (r *Repository) getOrderPartner(
	ctx context.Context,
	condition map[string]any,
	scopes ...func(db *gorm.DB) *gorm.DB,
) (orderPartner *models.OrderPartner, cusErr commonError.CustomError) {
	var orderPartners []*models.OrderPartner
	logTag := fmt.Sprintf("RequestID: %s Function: GetOrderPartner", env.GetRequestID(ctx))
	if resultErr := r.db.GetMasterDB(ctx).Where(condition).Scopes(scopes...).Find(&orderPartners).Error; resultErr != nil {
		cusErr = commonError.NewCustomError(customError.SqlFetchError, fmt.Sprintf("%s unable to read", logTag))
		return
	}

	if len(orderPartners) == 0 {
		cusErr = commonError.NewCustomError(customError.NotFound, "OrderPartner not found")
		return
	}

	orderPartner = orderPartners[0]
	return
}

// GetOrderPartnerByTenantID retrieves a single order partner record based on the tenant ID.
func (r *Repository) GetOrderPartnerByTenantID(
	ctx context.Context,
	tenantID string,
) (orderPartner *models.OrderPartner, cusErr commonError.CustomError) {
	conditions := map[string]any{
		"tenant_id": tenantID,
	}

	orderPartner, cusErr = r.getOrderPartner(ctx, conditions)
	if cusErr.Exists() {
		return
	}

	if orderPartner == nil {
		cusErr = commonError.NewCustomError(customError.NotFound, "OrderPartner not found")
		return
	}

	return
}
