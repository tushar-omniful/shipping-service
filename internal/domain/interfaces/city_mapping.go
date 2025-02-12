package interfaces

type CityMappingController interface {
	//GetCityMapping(ctx *gin.Context)
	//CreateCityMapping(ctx *gin.Context)
	//UpdateCityMapping(ctx *gin.Context)
	//DeleteCityMapping(ctx *gin.Context)
}

type CityMappingService interface {
	//GetCityMapping(ctx context.Context, id int64) (*models.CityMapping, oerror.CustomError)
	//CreateCityMapping(ctx context.Context, cityMapping *models.CityMapping) oerror.CustomError
	//UpdateCityMapping(ctx context.Context, condition map[string]interface{}, cityMapping *models.CityMapping) oerror.CustomError
	//DeleteCityMapping(ctx context.Context, id int64) oerror.CustomError
}

type CityMappingRepository interface {
	//GetCityMapping(ctx context.Context, condition map[string]interface{}, scopes ...func(db *gorm.DB) *gorm.DB) (*models.CityMapping, oerror.CustomError)
	//GetCityMappingByID(ctx context.Context, id int64) (*models.CityMapping, oerror.CustomError)
	//CreateCityMapping(ctx context.Context, cityMapping *models.CityMapping) oerror.CustomError
	//UpdateCityMapping(ctx context.Context, condition map[string]interface{}, cityMapping *models.CityMapping) oerror.CustomError
	//DeleteCityMapping(ctx context.Context, id int64) oerror.CustomError
}
