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
	redisDB "tender/storage/redis"
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

	rdb := redisDB.Connect()
	defer rdb.Close()

	storage := storage.NewStorage(db, rdb, logger)

	service := service.NewService(storage, logger)

	router := api.Router(service, logger, storage)

	go check.StartTenderStatusUpdater(storage, time.Hour)

	log.Printf("server is running...")
	log.Fatal(router.Run(cfg.API_PORT))

}
