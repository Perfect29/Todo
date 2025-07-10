package service

import (
	"context"
	log "github.com/sirupsen/logrus"
	"os"

	"github.com/redis/go-redis/v9"
)

func NewCacheRepository() *CacheRepository {
	rdb := redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_ADDR"),
		Password: "",
		DB: 0,
	})
	if err := rdb.Ping(context.Background()).Err(); err != nil {
		log.Error("Failed to connect to Redis:", err)
	}
	log.Info("Cache successfully created")
	return &CacheRepository{
		rdb: rdb,
	}
}

