package cache

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v9"
	"time"
)

type RedisRepository struct {
	client *redis.Client
}

func NewRedisRepository(client *redis.Client) *RedisRepository {
	return &RedisRepository{client: client}
}

var _ Repository = (*RedisRepository)(nil)

func (r *RedisRepository) StoreToCache(ctx context.Context, key string, val interface{}, exp time.Duration) error {
	if err := r.client.Set(ctx, key, val, exp).Err(); err != nil {
		return fmt.Errorf("r.client.Set: %w", err)
	}
	return nil
}

func (r *RedisRepository) GetStringByKey(ctx context.Context, key string) (string, error) {
	val, err := r.client.Get(ctx, key).Result()
	if err != nil {
		return "", fmt.Errorf("r.client.Get: %w", err)
	}
	return val, nil
}

func (r *RedisRepository) DeleteByKey(ctx context.Context, key string) error {
	if err := r.client.Del(ctx, key).Err(); err != nil {
		return fmt.Errorf("r.client.Del: %w", err)
	}
	return nil
}
