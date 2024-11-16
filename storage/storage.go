package storage

import (
	"database/sql"
	"log/slog"
)

type Storage interface{

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