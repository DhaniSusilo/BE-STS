package wizards

import (
	"learning-backend/config"
	"learning-backend/domains/handlers/http"
	"learning-backend/domains/repositories"
	"learning-backend/domains/usecases"
	"learning-backend/infrastructures"
)

var (
	Config           = config.GetConfig()
	PostgresDatabase = infrastructures.NewPostgresDatabase(Config)
	//
	DatabaseRepo = repositories.NewDatabaseRepository(PostgresDatabase)
	UseCase = usecases.NewUseCase(DatabaseRepo)
	Http = http.NewHttp(UseCase)
)
