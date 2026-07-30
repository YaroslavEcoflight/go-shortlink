package usecase

import (
	"auth-service/internal/domain/entity"
	"auth-service/internal/domain/repository"
	"auth-service/internal/domain/service"

	"golang.org/x/crypto/bcrypt"
)

type authService struct {
	repo repository.UserRepo
}

func NewAuthService(repo repository.UserRepo) service.AuthSerivce {
	return &authService{repo: repo}
}

func (s *authService) Register(user entity.User) (entity.User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.PasswordHash), bcrypt.DefaultCost)
	if err != nil {
		return entity.User{}, err
	}
	user.PasswordHash = string(hash)
	return s.repo.Create(user)
}

func (s *authService) Login(email, pass string) (*entity.Token, error) {
	return nil, nil
}

func (s *authService) Logout(refreshToken string) error {
	return nil
}

func (s *authService) RefreshToken(refreshToken string) (*entity.Token, error) {
	return nil, nil
}

func (s *authService) ValidateToken(accessToken string) (*entity.User, error) {
	return nil, nil
}
