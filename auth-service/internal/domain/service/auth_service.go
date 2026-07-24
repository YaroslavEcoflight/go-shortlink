package service

import "auth-service/internal/domain/entity"

type AuthSerivce interface {
	Register(entity.User) (entity.User, error)
	Login(email, pass string) (*entity.Token, error)
	Logout(refreshToken string) error
	RefreshToken(refreshToken string) (*entity.Token, error)
	ValidateToken(accessToken string) (*entity.User, error)
}
