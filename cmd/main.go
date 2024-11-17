package main

import (
	"log"
	"tender/api"
	"tender/api/middleware"
	"tender/config"
	"tender/pkg/casbin"
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

	casbin, err := casbin.CasbinEnforcer(logger)
	if err != nil {
		log.Fatal(err)
	}

	storage := storage.NewStorage(db, rdb, logger)

	service := service.NewService(storage, logger)

	router := api.Router(&api.Dependencies{
		ServiceManager: service,
		Enforcer:       casbin,
		Storage:        storage,
		RateLimiter:    *middleware.NewRateLimiter(5, 1),
		Logger:         logger,
	})

	go check.StartTenderStatusUpdater(storage, time.Hour)

	log.Printf("server is running...")
	log.Fatal(router.Run(cfg.API_PORT))

}
