package order_controller

import "gorm.io/gorm"

type orderFilters struct {
	SellerID string `form:"seller_id" json:"seller_id"`
}

func (p orderFilters) ToFilterScopes() (scopes []func(db *gorm.DB) *gorm.DB) {
	scopes = make([]func(db *gorm.DB) *gorm.DB, 0)

	if len(p.SellerID) > 0 {
		scopes = append(scopes, func(db *gorm.DB) *gorm.DB {
			return db.Where("seller_id = ?", p.SellerID)
		})
	}
	return
}
