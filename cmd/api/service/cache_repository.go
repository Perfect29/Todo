package service

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

type CacheRepository struct {
	rdb *redis.Client
}

type Cache interface {
	GetTodo(ctx context.Context, id int) (*Todo, error)
	SetTodo(ctx context.Context, todo *Todo, ttl time.Duration) error
	DeleteTodo(ctx context.Context, id int) error
}

func (c *CacheRepository) GetTodo(ctx context.Context, id int) (*Todo, error) {
	idStr := strconv.Itoa(id)
	val, err := c.rdb.Get(ctx, idStr).Result()
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

func (c *CacheRepository) SetTodo(ctx context.Context, todo *Todo, ttl time.Duration) error {
	data, err := json.Marshal(todo)
	if err != nil {
		return err
	}
	err = c.rdb.Set(ctx, fmt.Sprint(todo.ID), data, ttl).Err()
	if err != nil {
		return err
	}
	return nil
}

func (c *CacheRepository) DeleteTodo(ctx context.Context, id int) error {
	key := fmt.Sprint(id)
	err := c.rdb.Del(ctx, key).Err()
	if err != nil {
		return err
	}
	return nil
}