package repository

import (
	"auth-service/internal/domain/entity"
	"time"
)

type TokenRepo interface {
	Save(userID string, token entity.Token, ttl time.Duration) error
	Get(refreshToken string) (string, error)
	Delete(refreshToken string) error
}
