package redis

import (
	"auth-service/internal/domain/entity"
	"auth-service/internal/domain/repository"
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type TokenRepo struct {
	client *redis.Client
	ctx    context.Context
}

func NewTokenRepo(ctx context.Context, client *redis.Client) repository.TokenRepo {
	return &TokenRepo{client: client, ctx: ctx}
}

func (r *TokenRepo) Save(userID string, toekn entity.Token, ttl time.Duration) error {
	return r.client.Set(r.ctx, "refresh:"+toekn.RefreshToken, userID, ttl).Err()
}

func (r *TokenRepo) Get(refreshToken string) (string, error) {
	return r.client.Get(r.ctx, "refresh:"+refreshToken).Result()
}

func (r *TokenRepo) Delete(refreshToken string) error {
	return r.client.Del(r.ctx, "refresh:"+refreshToken).Err()
}
