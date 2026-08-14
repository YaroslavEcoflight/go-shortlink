package postgres

import (
	"auth-service/internal/domain/entity"
	"auth-service/internal/domain/repository"
	"auth-service/internal/infrastructure/postgres/models"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) repository.UserRepo {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user entity.User) (entity.User, error) {
	m := to_model(user)
	if err := r.db.Create(&m).Error; err != nil {
		return entity.User{}, err
	}
	return to_entity(m), nil
}

func (r *userRepository) GetById(id string) (entity.User, error) {
	var m models.User
	if err := r.db.First(&m, "id = ?", id).Error; err != nil {
		return entity.User{}, err
	}
	return to_entity(m), nil
}

func (r *userRepository) GetByEmail(email string) (entity.User, error) {
	var m models.User
	if err := r.db.Where("email = ?", email).First(&m).Error; err != nil {
		return entity.User{}, err
	}
	return to_entity(m), nil
}

func to_entity(user models.User) entity.User {
	return entity.User{
		ID:           user.ID,
		Username:     user.Username,
		PasswordHash: user.PasswordHash,
		Email:        user.Email,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
	}
}

func to_model(user entity.User) models.User {
	return models.User{
		ID:           user.ID,
		Username:     user.Username,
		PasswordHash: user.PasswordHash,
		Email:        user.Email,
	}
}
