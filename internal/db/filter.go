package db

import (
	"github.com/gin-gonic/gin"
	commonError "github.com/omniful/go_commons/error"
	"gorm.io/gorm"
)

type FilterScopes interface {
	ToFilterScopes() []func(db *gorm.DB) *gorm.DB
}

func GetFilterParamScope(ctx *gin.Context, filter FilterScopes) (scopes []func(db *gorm.DB) *gorm.DB, cusErr commonError.CustomError) {
	scopes = filter.ToFilterScopes()
	return
}
