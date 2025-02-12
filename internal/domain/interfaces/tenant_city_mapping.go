package interfaces

type TenantCityMappingController interface {
	//GetTenantCityMapping(ctx *gin.Context)
	//CreateTenantCityMapping(ctx *gin.Context)
	//UpdateTenantCityMapping(ctx *gin.Context)
	//DeleteTenantCityMapping(ctx *gin.Context)
}

type TenantCityMappingService interface {
	//GetTenantCityMapping(ctx context.Context, id int64) (*models.TenantCityMapping, oerror.CustomError)
	//CreateTenantCityMapping(ctx context.Context, mapping *models.TenantCityMapping) oerror.CustomError
	//UpdateTenantCityMapping(ctx context.Context, condition map[string]interface{}, mapping *models.TenantCityMapping) oerror.CustomError
	//DeleteTenantCityMapping(ctx context.Context, id int64) oerror.CustomError
}

type TenantCityMappingRepository interface {
	//GetTenantCityMapping(ctx context.Context, condition map[string]interface{}, scopes ...func(db *gorm.DB) *gorm.DB) (*models.TenantCityMapping, oerror.CustomError)
	//GetTenantCityMappingByID(ctx context.Context, id int64) (*models.TenantCityMapping, oerror.CustomError)
	//CreateTenantCityMapping(ctx context.Context, mapping *models.TenantCityMapping) oerror.CustomError
	//UpdateTenantCityMapping(ctx context.Context, condition map[string]interface{}, mapping *models.TenantCityMapping) oerror.CustomError
	//DeleteTenantCityMapping(ctx context.Context, id int64) oerror.CustomError
}
