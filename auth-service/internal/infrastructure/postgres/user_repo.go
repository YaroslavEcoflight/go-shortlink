package postgres

import (
	"auth-service/internal/domain/entity"
	"auth-service/internal/domain/repository"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) repository.UserRepo {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user entity.User) (entity.User, error) {
	return entity.User{}, nil
}

func (r *userRepository) GetById(id string) (entity.User, error) {
	return entity.User{}, nil
}

func (r *userRepository) GetByEmail(email string) (entity.User, error) {
	return entity.User{}, nil
}
