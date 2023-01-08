package cache

import (
	"context"
	"time"
)

type Repository interface {
	create
	read
	del
}

type create interface {
	StoreToCache(ctx context.Context, key string, val interface{}, exp time.Duration) error
}

type read interface {
	GetStringByKey(ctx context.Context, key string) (string, error)
}

type del interface {
	DeleteByKey(ctx context.Context, key string) error
}
