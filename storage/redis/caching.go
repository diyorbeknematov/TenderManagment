package redisDB

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/redis/go-redis/v9"
)

type CachingRepo interface {
	GetCache(key string) (string, error)
	SetCache(key string, value string, expiration time.Duration) error
}

type cachingImpl struct {
	RDB *redis.Client
	Log *slog.Logger
}

func NewCacingRepo(rdb *redis.Client, logger *slog.Logger) CachingRepo {
	return &cachingImpl{
		RDB: rdb,
		Log: logger,
	}
}

func (C *cachingImpl) GetCache(key string) (string, error) {
	val, err := C.RDB.Get(context.Background(), key).Result()
	if err == redis.Nil {
		// Agar key Redisda topilmasa
		return "", nil
	} else if err != nil {
		return "", fmt.Errorf("redisdan o'qish xatosi: %v", err)
	}
	return val, nil
}

func (C *cachingImpl) SetCache(key string, value string, expiration time.Duration) error {
	err := C.RDB.Set(context.Background(), key, value, expiration).Err()
	if err != nil {
		return fmt.Errorf("redisga yozishda xato: %v", err)
	}
	return nil
}
