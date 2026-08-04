package usecase

import (
	"auth-service/internal/domain/entity"
	"auth-service/internal/domain/repository"
	"auth-service/internal/domain/service"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type authService struct {
	user_repo  repository.UserRepo
	token_repo repository.TokenRepo
	jwtSecret  string
}

func NewAuthService(repo repository.UserRepo) service.AuthSerivce {
	return &authService{user_repo: repo}
}

func (s *authService) Register(user entity.User) (entity.User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.PasswordHash), bcrypt.DefaultCost)
	if err != nil {
		return entity.User{}, err
	}
	user.PasswordHash = string(hash)
	return s.user_repo.Create(user)
}

func (s *authService) Login(email, pass string) (*entity.Token, error) {
	user, err := s.user_repo.GetByEmail(email)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(pass)); err != nil {
		return nil, errors.New("invalid credentials")
	}

	now := time.Now()
	accessExpiry := now.Add(15 * time.Minute)
	claims := jwt.MapClaims{
		"sub": user.ID,
		"exp": accessExpiry.Unix(),
		"iat": now.Unix(),
	}
	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(s.jwtSecret))
	if err != nil {
		return nil, err
	}

	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return nil, err
	}
	refreshToken := hex.EncodeToString(b)

	token := entity.Token{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresIn:    accessExpiry.Unix(),
	}

	if err := s.token_repo.Save(user.ID, token, 7*24*time.Hour); err != nil {
		return nil, err
	}
	return &token, nil
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
