package storage

import (
	"database/sql"
	"log/slog"
	"tender/storage/postgres"
)

type Storage interface{
	Client()postgres.ClientRepo
}

type storageImpl struct{
	DB *sql.DB
	Log *slog.Logger
}

func NewStorage(db *sql.DB, logger *slog.Logger)Storage{
	return &storageImpl{
		DB: db,
		Log: logger,
	}
}

func(S *storageImpl) Client()postgres.ClientRepo{
	return postgres.NewClientRepo(S.DB, S.Log)
}