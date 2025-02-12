package order_repo

import (
	"context"
	"fmt"
	oerror "github.com/omniful/go_commons/error"
	"github.com/omniful/shipping-service/internal/domain/models"
	custom_error "github.com/omniful/shipping-service/pkg/error"
	"sync"

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

func (r *Repository) CreateOrder(
	ctx context.Context,
	order *models.Order,
) (*models.Order, oerror.CustomError) {
	result := r.db.GetMasterDB(ctx).Model(models.Order{}).Create(&order)
	if resultErr := result.Error; resultErr != nil {
		cusErr := oerror.NewCustomError(custom_error.SqlCreateError, fmt.Sprintf("Could not create order : %+v, err: %v", order, resultErr))
		return nil, cusErr
	}
	return order, oerror.CustomError{}
}
