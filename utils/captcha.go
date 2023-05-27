package utils

import (
	"context"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"project/database"
	"time"
)

var Redis_Store *RedisStore

func init() {
	Redis_Store = &RedisStore{Client: database.NewRedisClient()}
}

type RedisStore struct {
	Client *redis.Client
}

// Set 根据键值设置到Redis中
func (r *RedisStore) Set(id string, value string) error {
	return r.Client.Set(context.Background(), id, value, time.Minute*5).Err()
}

// Get 根据键从Redis获取值
// 如果clear为true,则在获取值后删除该键
func (r *RedisStore) Get(id string, clear bool) string {
	result, err := r.Client.Get(context.Background(), id).Result()
	if err != nil {
		zap.S().Info("Redis Get Error：", err)
		return ""
	}
	if clear {
		if err = r.Client.Get(context.Background(), id).Err(); err != nil {
			zap.S().Info("Redis failed to delete data:", err)
			return ""
		}
	}
	return result
}

func (r *RedisStore) Verify(id, answer string, clear bool) bool {
	val := r.Get(id, clear)
	return val == answer
}
