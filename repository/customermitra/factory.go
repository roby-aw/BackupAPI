package customermitra

import (
	"api-redeem-point/business/customermitra"
	"api-redeem-point/utils"
)

func RepositoryFactory(dbCon *utils.DatabaseConnection) customermitra.Repository {
	dummyRepo := NewPostgresRepository(dbCon.Postgres)
	return dummyRepo
}
