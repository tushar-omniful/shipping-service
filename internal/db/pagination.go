package db

import (
	"context"
	"fmt"
	commonError "github.com/omniful/go_commons/error"
	custom_error "github.com/omniful/shipping-service/pkg/error"
	"gorm.io/gorm"
)

// this function is common function for finding the count and data.
// here countScopes is scope for finding the count and it does not contain scopes for ordering and pagination.
// scopes is scope for finding the data and it contain all scopes.
// type T should be model.

func GetPaginatedDataT[T any](ctx context.Context, db *gorm.DB, model T, condition map[string]interface{}, countScopes []func(db *gorm.DB) *gorm.DB, scopes ...func(db *gorm.DB) *gorm.DB) ([]*T, int64, commonError.CustomError) {
	var totalCount int64
	if len(countScopes) > 0 {
		err := db.Model(model).Where(condition).Scopes(countScopes...).Count(&totalCount).Error
		if err != nil {
			cusErr := commonError.NewCustomError(custom_error.SqlFetchError, fmt.Sprintf("unable to fetch :: %v", err.Error()))
			return nil, 0, cusErr
		}
	}

	var data []*T
	err := db.Model(model).Where(condition).Scopes(scopes...).Find(&data).Error
	if err != nil {
		cusErr := commonError.NewCustomError(custom_error.SqlFetchError, fmt.Sprintf("unable to fetch :: %v", err.Error()))
		return nil, 0, cusErr
	}

	return data, totalCount, commonError.CustomError{}
}
