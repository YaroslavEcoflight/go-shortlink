package repository

import (
	"auth-service/internal/domain/entity"
)

type UserRepo interface {
	Create(entity entity.User) (entity.User, error)
	GetById(id string) (entity.User, error)
	GetByEmail(email string) (entity.User, error)
}
