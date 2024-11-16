package storage

import (
	"database/sql"
	"log/slog"
	"tender/storage/postgres"
)

type Storage interface {
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

func (s storageImpl) RegistrationRepository() postgres.RegistrationRepository {
	return postgres.NewRegistrationRepository(s.DB, s.Log)
}
