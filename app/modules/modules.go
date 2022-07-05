package modules

import (
	"api-redeem-point/api"
	adminApi "api-redeem-point/api/admin"
	customermitraApi "api-redeem-point/api/customermitra"
	adminBusiness "api-redeem-point/business/admin"
	customermitraBusiness "api-redeem-point/business/customermitra"
	"api-redeem-point/config"
	adminRepo "api-redeem-point/repository/admin"
	customermitraRepo "api-redeem-point/repository/customermitra"
	"api-redeem-point/utils"
)

func RegistrationModules(dbCon *utils.DatabaseConnection, config *config.AppConfig) api.Controller {

	adminPermitRepository := adminRepo.RepositoryFactory(dbCon)
	adminPermitService := adminBusiness.NewService(adminPermitRepository)
	adminPermitController := adminApi.NewController(adminPermitService)

	customermitraPermitRepository := customermitraRepo.RepositoryFactory(dbCon)
	customermitraPermitService := customermitraBusiness.NewService(customermitraPermitRepository)
	customermitraPermitController := customermitraApi.NewController(customermitraPermitService)

	controller := api.Controller{
		AdminControlller:        adminPermitController,
		CustomerMitraController: customermitraPermitController,
	}
	return controller
}
