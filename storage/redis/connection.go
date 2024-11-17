package redisDB

import "github.com/redis/go-redis/v9"

func Connect()*redis.Client{
	return redis.NewClient(&redis.Options{
		Addr: "redis:6379",
		Password: "",
		DB: 0,
	})
}