package redis

import (
	"auth-service/config"
	"context"

	"github.com/redis/go-redis/v9"
)

func New(ctx context.Context, cfg config.Config) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Addr,
		Username: cfg.Redis.User,
		Password: "pass",
	})
	if err := rdb.Ping(ctx).Err(); err != nil {
		return nil, err
	}
	return rdb, nil
}
