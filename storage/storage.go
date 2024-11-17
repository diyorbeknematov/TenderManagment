package storage

import (
	"database/sql"
	"log/slog"
	"tender/storage/postgres"
	redisDB "tender/storage/redis"

	"github.com/redis/go-redis/v9"
)

type Storage interface {
	Client() postgres.ClientRepo
	RegistrationRepository() postgres.RegistrationRepository
	Contractor() postgres.BidRepository
	NotificationRepository() postgres.NotificationRepository
	Caching() redisDB.CachingRepo
}

type storageImpl struct {
	DB  *sql.DB
	RDB *redis.Client
	Log *slog.Logger
}

func NewStorage(db *sql.DB, rdb *redis.Client, logger *slog.Logger) Storage {
	return &storageImpl{
		DB:  db,
		RDB: rdb,
		Log: logger,
	}
}

func (S *storageImpl) Client() postgres.ClientRepo {
	return postgres.NewClientRepo(S.DB, S.Log)
}

func (S *storageImpl) RegistrationRepository() postgres.RegistrationRepository {
	return postgres.NewRegistrationRepository(S.DB, S.Log)
}

func (S *storageImpl) NotificationRepository() postgres.NotificationRepository {
	return postgres.NewNotificationRepository(S.DB, S.Log)
}

func (s *storageImpl) Contractor() postgres.BidRepository {
	return postgres.NewBidRepository(s.DB)
}

func (s *storageImpl) Caching() redisDB.CachingRepo {
	return redisDB.NewCacingRepo(s.RDB, s.Log)
}
