package database

import (
	"github.com/redis/go-redis/v9"
	"project/config"
)

func NewRedisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     config.Redis.Addr,
		Password: config.Redis.Password,
		DB:       0,
	})
}
