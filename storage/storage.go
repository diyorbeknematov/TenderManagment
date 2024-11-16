package storage

import (
	"database/sql"
	"log/slog"
	"tender/storage/postgres"
)

type Storage interface{
	Client()postgres.ClientRepo
	RegistrationRepository() postgres.RegistrationRepository
}

type storageImpl struct {
	DB  *sql.DB
	Log *slog.Logger
}

func NewStorage(db *sql.DB, logger *slog.Logger) Storage {
	return &storageImpl{
		DB:  db,
		Log: logger,
	}
}

func(S *storageImpl) Client()postgres.ClientRepo{
	return postgres.NewClientRepo(S.DB, S.Log)
}

func (s storageImpl) RegistrationRepository() postgres.RegistrationRepository {
	return postgres.NewRegistrationRepository(s.DB, s.Log)
}
