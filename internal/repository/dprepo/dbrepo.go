package dbrepo

import (
	"database/sql"
	"github.com/AidosK/GoFinal/internal/models"

	"github.com/AidosK/GoFinal/internal/config"
	"github.com/AidosK/GoFinal/internal/repository"
)

type PostgresDBRepo struct {
	App *config.AppConfig
	DB  *sql.DB
}

func (repo *PostgresDBRepo) GetUserByEmail(email string) (interface{}, error) {

	panic("implement me")
}

type testDBRepo struct {
	App *config.AppConfig
	DB  *sql.DB
}

func (repo *testDBRepo) GetUserByEmail(email string) (interface{}, error) {

	panic("implement me")
}

func (repo *testDBRepo) InsertUser(user models.User) (models.User, error) {

	panic("implement me")
}

func NewPostgresRepo(dbConnection *sql.DB, appConfig *config.AppConfig) *PostgresDBRepo {
	return &PostgresDBRepo{
		App: appConfig,
		DB:  dbConnection,
	}
}

func NewTestRepo(appConfig *config.AppConfig) repository.DatabaseRepo {
	return &testDBRepo{
		App: appConfig,
	}
}
