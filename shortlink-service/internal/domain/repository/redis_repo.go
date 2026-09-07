package repository

import (
	"errors"
	"shortlink-service/internal/domain/entity"
)

var ErrCacheMiss = errors.New("cache miss")

type LinkCache interface {
	Set(code string, link entity.Link) error
	Get(code string) (entity.Link, error)
	Delete(code string) error
}
