package service

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"
	"log"
	"os"

	"github.com/redis/go-redis/v9"
)

type CacheService struct {
	rdb *redis.Client
	ctx context.Context
}

func NewCacheService() *CacheService {
	rdb := redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_ADDR"),
		Password: "",
		DB: 0,
	})
	if err := rdb.Ping(context.Background()).Err(); err != nil {
		log.Fatal("Failed to connect to Redis:", err)
	}
	return &CacheService{
		rdb: rdb,
		ctx: context.Background(),
	}
}

func (c *CacheService) GetTodo(id int) (*Todo, error) {
	idStr := strconv.Itoa(id)
	val, err := c.rdb.Get(c.ctx, idStr).Result()
	if err == redis.Nil {
		return nil, nil
	} else if err != nil {
		return nil, err
	} 
	var todo Todo
	err = json.Unmarshal([]byte(val), &todo)
	if err != nil {
		return nil, err
	}
	return &todo, nil
}

func (c *CacheService) SetTodo(todo *Todo, ttl time.Duration) error {
	data, err := json.Marshal(todo)
	if err != nil {
		return err
	}
	err = c.rdb.Set(c.ctx, fmt.Sprint(todo.ID), data, ttl).Err()
	if err != nil {
		return err
	}
	return nil
}

func (c *CacheService) DeleteTodo(id int) error {
	key := fmt.Sprint(id)
	err := c.rdb.Del(c.ctx, key).Err()
	if err != nil {
		return err
	}
	return nil
}