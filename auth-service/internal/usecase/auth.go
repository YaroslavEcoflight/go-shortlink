package usecase

import (
	"auth-service/internal/domain/entity"
	"auth-service/internal/domain/repository"
)

type AuthService struct {
	repo repository.UserRepo
}

func (s *AuthService) Register(user entity.User) (entity.User, error) {
	created, err := s.repo.Create(user)
	if err != nil {
		return entity.User{}, err
	}
	return created, nil
}
