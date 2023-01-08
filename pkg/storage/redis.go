package storage

import (
	"contactsList/config"
	"context"
	"fmt"
	"github.com/go-redis/redis/v9"
	"time"
)

func InitRedisClient(cfg *config.Redis) (*redis.Client, error) {
	opts := &redis.Options{
		Addr:         fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
		MinIdleConns: cfg.MinIdleConns,
		PoolSize:     cfg.PoolSize,
		PoolTimeout:  time.Duration(cfg.PoolTimeout) * time.Second,
		Password:     cfg.Password,
		DB:           cfg.DB,
	}

	client := redis.NewClient(opts)
	pingCtx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	if result := client.Ping(pingCtx); result.Err() != nil {
		return nil, result.Err()
	}

	return client, nil
}
