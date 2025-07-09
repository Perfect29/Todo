package service

import (
	"context"
	"log"
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
		log.Fatal("Failed to connect to Redis:", err)
	}
	return &CacheRepository{
		rdb: rdb,
	}
}

