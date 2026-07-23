package redis

import (
	"context"

	"github.com/redis/go-redis/v9"
)

func New(ctx context.Context) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Username: "admin",
		Password: "pass",
	})
	if err := rdb.Ping(ctx).Err(); err != nil {
		return nil, err
	}
	return rdb, nil
}
