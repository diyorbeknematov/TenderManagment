package main

import (
	"log"
	"tender/api"
	"tender/config"
	"tender/pkg/check"
	"tender/pkg/logs"
	"tender/service"
	"tender/storage"
	"tender/storage/postgres"
	"time"
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

	service := service.NewService(storage, logger)

	router := api.Router(service, logger)

	go check.StartTenderStatusUpdater(storage, time.Hour)

	log.Printf("server is running...")
	log.Fatal(router.Run(cfg.API_PORT))

}
