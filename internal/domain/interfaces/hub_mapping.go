package interfaces

type HubMappingController interface {
	//GetHubMapping(ctx *gin.Context)
	//CreateHubMapping(ctx *gin.Context)
	//UpdateHubMapping(ctx *gin.Context)
	//DeleteHubMapping(ctx *gin.Context)
}

type HubMappingService interface {
	//GetHubMapping(ctx context.Context, id int64) (*models.HubMapping, oerror.CustomError)
	//CreateHubMapping(ctx context.Context, hubMapping *models.HubMapping) oerror.CustomError
	//UpdateHubMapping(ctx context.Context, condition map[string]interface{}, hubMapping *models.HubMapping) oerror.CustomError
	//DeleteHubMapping(ctx context.Context, id int64) oerror.CustomError
}

type HubMappingRepository interface {
	//GetHubMapping(ctx context.Context, condition map[string]interface{}, scopes ...func(db *gorm.DB) *gorm.DB) (*models.HubMapping, oerror.CustomError)
	//GetHubMappingByID(ctx context.Context, id int64) (*models.HubMapping, oerror.CustomError)
	//CreateHubMapping(ctx context.Context, hubMapping *models.HubMapping) oerror.CustomError
	//UpdateHubMapping(ctx context.Context, condition map[string]interface{}, hubMapping *models.HubMapping) oerror.CustomError
	//DeleteHubMapping(ctx context.Context, id int64) oerror.CustomError
}
