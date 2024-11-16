package main

import (
	"log"
	"tender/api"
	"tender/api/handler"
	"tender/config"
	"tender/logs"
	"tender/service"
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
	service := service.NewService(storage, logger)
	hand := handler.NewHandler(*service, logger)
	router := api.Router(hand)
	log.Printf("server is running...")
	log.Fatal(router.Run(cfg.API_PORT))
}
