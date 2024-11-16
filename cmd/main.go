package main

import (
	"tender/config"
	"tender/logs"
	"tender/storage"
	"tender/storage/postgres"
)

func main() {
	cfg := config.LoadConfig()
	logger := logs.InitLogger()

	db, err := postgres.Connect(cfg)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	defer db.Close()

	storage := storage.NewStorage(db, logger)
}
