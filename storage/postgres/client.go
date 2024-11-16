package postgres

import (
	"database/sql"
	"log/slog"
)

type ClientRepo interface{

}

type clientImpl struct{
	DB *sql.DB
	Log *slog.Logger
}

func NewClientRepo(db *sql.DB, logger *slog.Logger)ClientRepo{
	return &clientImpl{
		DB: db,
		Log: logger,
	}
}

func (C *clientImpl) CreateTender()