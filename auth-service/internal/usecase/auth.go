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
	userRepo  repository.UserRepo
	tokenRepo repository.TokenRepo
	jwtSecret string
}

func NewAuthService(userRepo repository.UserRepo, tokenRepo repository.TokenRepo, jwtSecret string) service.AuthService {
	return &authService{userRepo: userRepo, tokenRepo: tokenRepo, jwtSecret: jwtSecret}
}

func (s *authService) Register(user entity.User) (entity.User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.PasswordHash), bcrypt.DefaultCost)
	if err != nil {
		return entity.User{}, err
	}
	user.PasswordHash = string(hash)
	return s.userRepo.Create(user)
}

func (s *authService) Login(email, pass string) (*entity.Token, error) {
	user, err := s.userRepo.GetByEmail(email)
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

	if err := s.tokenRepo.Save(user.ID, token, 7*24*time.Hour); err != nil {
		return nil, err
	}
	return &token, nil
}

func (s *authService) Logout(refreshToken string) error {
	if err := s.tokenRepo.Delete(refreshToken); err != nil {
		return err
	}
	return nil
}

func (s *authService) RefreshToken(refreshToken string) (*entity.Token, error) {
	userID, err := s.tokenRepo.Get(refreshToken)
	if err != nil {
		return nil, errors.New("Invalid or expired refresh token")
	}

	if err := s.tokenRepo.Delete(refreshToken); err != nil {
		return nil, err
	}

	now := time.Now()
	accessExpiry := now.Add(15 * time.Minute)
	claims := jwt.MapClaims{
		"sub": userID,
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
	newRefreshToken := hex.EncodeToString(b)

	token := entity.Token{
		AccessToken:  accessToken,
		RefreshToken: newRefreshToken,
		ExpiresIn:    accessExpiry.Unix(),
	}
	if err := s.tokenRepo.Save(userID, token, 7*24*time.Hour); err != nil {
		return nil, err
	}
	return &token, nil
}

func (s *authService) ValidateToken(accessToken string) (*entity.User, error) {
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(s.jwtSecret), nil
	})
	if err != nil || !token.Valid {
		return nil, errors.New("Invalid or expired access token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("Invalid claims")
	}
	userID, ok := claims["sub"].(string)
	if !ok {
		return nil, errors.New("Invalid sub claim")
	}

	user, err := s.userRepo.GetById(userID)
	if err != nil {
		return nil, errors.New("user not found")
	}
	return &user, nil
}
