package db

import (
	"github.com/gin-gonic/gin"
	commonError "github.com/omniful/go_commons/error"
	"github.com/omniful/shipping-service/constants"
	customError "github.com/omniful/shipping-service/pkg/error"
	"gorm.io/gorm"
	"strings"
)

type SortableMap map[string]string

type SortableInterface interface {
	GetDefaultSort() string
	GetField(field string) (string, bool)
}

type OrderSort string

const (
	Asc  OrderSort = "asc"
	Desc OrderSort = "desc"
)

type Sortable struct {
	field string
	order OrderSort
}

var stringToSortOrderType = map[string]OrderSort{
	"asc":  Asc,
	"desc": Desc,
}

var SortOrderTypeToString = map[OrderSort]string{
	Asc:  "asc",
	Desc: "desc",
}

func GetSortingScopes(ctx *gin.Context, sortMap SortableInterface) (scopes []func(db *gorm.DB) *gorm.DB, cusErr commonError.CustomError) {
	scopes = make([]func(db *gorm.DB) *gorm.DB, 0)
	sortArray := ctx.QueryArray(constants.Sort)
	sort, cusErr := ParseSortArray(sortArray)
	if cusErr.Exists() {
		return
	}

	for _, value := range sort {
		key, found := sortMap.GetField(value.field)
		if !found {
			continue
		}

		if _, ok := SortOrderTypeToString[value.order]; !ok {
			continue
		}

		scopes = append(scopes, func(db *gorm.DB) *gorm.DB {
			return db.Order(key + " " + SortOrderTypeToString[value.order])
		})
	}

	scopes = append(scopes, func(db *gorm.DB) *gorm.DB {
		return db.Order(sortMap.GetDefaultSort())
	})

	return
}

func ParseSortArray(sortArray []string) ([]Sortable, commonError.CustomError) {
	sort := make([]Sortable, 0)
	for _, value := range sortArray {
		parts := strings.Split(value, "_")
		if len(parts) != 2 {
			cusErr := commonError.NewCustomError(customError.BadRequest, "invalid sort order")
			return nil, cusErr
		}

		order, ok := stringToSortOrderType[parts[1]]
		if !ok {
			cusErr := commonError.NewCustomError(customError.BadRequest, "invalid sort order")
			return nil, cusErr
		}

		sort = append(sort, Sortable{
			field: parts[0],
			order: order,
		})
	}
	return sort, commonError.CustomError{}
}
