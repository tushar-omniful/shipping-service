package order_repo

import (
	"context"
	"fmt"
	"sync"

	oerror "github.com/omniful/go_commons/error"
	"github.com/omniful/shipping-service/internal/domain/models"
	custom_error "github.com/omniful/shipping-service/pkg/error"

	"github.com/omniful/go_commons/db/sql/postgres"
)

type Repository struct {
	db *postgres.DbCluster
}

var repo *Repository
var repoOnce sync.Once

func NewRepository(db *postgres.DbCluster) *Repository {
	repoOnce.Do(func() {
		repo = &Repository{
			db: db,
		}
	})

	return repo
}

func (r *Repository) CreateOrder(ctx context.Context, order *models.Order) (*models.Order, oerror.CustomError) {
	result := r.db.GetMasterDB(ctx).Model(models.Order{}).Create(&order)
	if resultErr := result.Error; resultErr != nil {
		cusErr := oerror.NewCustomError(custom_error.SqlCreateError, fmt.Sprintf("Could not create order : %+v, err: %v", order, resultErr))
		return nil, cusErr
	}
	return order, oerror.CustomError{}
}

// CheckExistingOrder checks if an order exists with the given order partner order ID and order partner ID,
// excluding certain statuses
func (r *Repository) CheckExistingOrder(ctx context.Context, orderPartnerOrderID string, orderPartnerID uint64) (exists bool, cusErr oerror.CustomError) {
	var count int64
	excludedStatuses := []models.OrderStatus{
		models.Created,
		models.NewOrder,
		models.ReturnToOrigin,
		models.Cancelled,
	}

	result := r.db.GetMasterDB(ctx).Model(&models.Order{}).
		Where("order_partner_order_id = ? AND order_partner_id = ?", orderPartnerOrderID, orderPartnerID).
		Not("status IN ?", excludedStatuses).
		Count(&count)

	if result.Error != nil {
		cusErr = oerror.NewCustomError(custom_error.SqlFetchError, fmt.Sprintf("Failed to check existing order: %v", result.Error))
		return
	}

	exists = count > 0
	return
}
