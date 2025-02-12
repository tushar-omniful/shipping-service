package db

import "gorm.io/gorm"

func GetDeletedScope() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("deleted_at is NULL")
	}
}
