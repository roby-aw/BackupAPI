package admin

import (
	"api-redeem-point/business/admin"
	"api-redeem-point/utils"
)

func RepositoryFactory(dbCon *utils.DatabaseConnection) admin.Repository {
	adminRepo := NewPosgresRepository(dbCon.Postgres)
	return adminRepo
}
