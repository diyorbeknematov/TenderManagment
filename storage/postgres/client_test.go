package postgres

import (
	"database/sql"
	"log/slog"
	"tender/logs"
	"tender/model"
	"testing"
)

func ConnectDB(logger *slog.Logger) *sql.DB {
	connector := "host = localhost user = postgres port = 5432 dbname = tender password = hamidjon4424 sslmode = disable"
	db, err := sql.Open("postgres", connector)
	if err != nil {
		logger.Error(err.Error())
		return nil
	}
	return db
}

var logger = logs.InitLogger()

func Test_CreateTender(t *testing.T){
	db := ConnectDB(logger)
	c := NewClientRepo(db, logger)

	_, err := c.CreateTender(&model.CreateTenderReq{
		ClientId: "",
	})
}
