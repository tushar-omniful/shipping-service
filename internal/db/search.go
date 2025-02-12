package db

import (
	"github.com/gin-gonic/gin"
	commonError "github.com/omniful/go_commons/error"
	"github.com/omniful/shipping-service/constants"
	"gorm.io/gorm"
	"strings"
)

func GetSearchParamsScopes(ctx *gin.Context, searchColumnMap map[string]string) (scopes []func(db *gorm.DB) *gorm.DB, cusErr commonError.CustomError) {
	searchColumn := ctx.Query(constants.SearchColumn)
	searchQuery := strings.TrimSpace(ctx.Query(constants.SearchQuery))

	scopes = make([]func(db *gorm.DB) *gorm.DB, 0)
	var mappedSearchColumn string
	var ok bool

	if mappedSearchColumn, ok = searchColumnMap[searchColumn]; ok && len(searchQuery) > 0 {
		scopes = append(scopes, func(db *gorm.DB) *gorm.DB {
			return db.Where(mappedSearchColumn+" LIKE ?", searchQuery+"%")
		})
	}

	return
}

func GetCaseInsensitiveSearchParamsScopes(ctx *gin.Context, searchColumnMap map[string]string) (scopes []func(db *gorm.DB) *gorm.DB, cusErr commonError.CustomError) {
	searchColumn := ctx.Query(constants.SearchColumn)
	searchQuery := strings.TrimSpace(ctx.Query(constants.SearchQuery))

	scopes = make([]func(db *gorm.DB) *gorm.DB, 0)
	var mappedSearchColumn string
	var ok bool

	if mappedSearchColumn, ok = searchColumnMap[searchColumn]; ok && len(searchQuery) > 0 {
		scopes = append(scopes, func(db *gorm.DB) *gorm.DB {
			return db.Where(mappedSearchColumn+" ILIKE ?", searchQuery+"%")
		})
	}

	return
}
