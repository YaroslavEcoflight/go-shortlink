package redis

import (
	"context"
	"encoding/json"
	"errors"
	"shortlink-service/internal/domain/entity"
	"shortlink-service/internal/domain/repository"
	"time"

	"github.com/redis/go-redis/v9"
)

const ttl = 24 * time.Hour


type LinkCache struct {
	client *redis.Client
	ctx    context.Context
}

func NewLinkRepos(ctx context.Context, client *redis.Client) repository.LinkCache {
	return &LinkCache{ctx: ctx, client: client}
}

func (r *LinkCache) Get(code string) (entity.Link, error) {
	data, err := r.client.Get(r.ctx, "link:"+code).Bytes()
	if errors.Is(err, redis.Nil) {
		return entity.Link{}, repository.ErrCacheMiss
	}
	if err != nil {
		return entity.Link{}, err
	}

	var link entity.Link

	return link, json.Unmarshal(data, &link)
}

func (r *LinkCache) Set(code string, link entity.Link) error {
	data, err := json.Marshal(link)
	if err != nil {
		return err
	}
	return r.client.Set(r.ctx, "link:"+code, data, ttl).Err()
}

func (r *LinkCache) Delete(code string) error {
	return r.client.Del(r.ctx, "link:"+code).Err()
}
