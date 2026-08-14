package restapi

import (
	"auth-service/internal/domain/entity"
	"auth-service/internal/domain/service"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	svc service.AuthService
}

func NewAuthHandler(svc service.AuthService) AuthHandler {
	return AuthHandler{svc: svc}
}

func (h *AuthHandler) Register(c *fiber.Ctx) error {
	var body UserRegisterRequest
	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid body"})
	}
	n := entity.User{
		Username:     body.Username,
		PasswordHash: body.Password,
		Email:        body.Email,
	}
	_, err := h.svc.Register(n)
	if err != nil {
		return c.Status(409).JSON(fiber.Map{"error": "user exist"})
	}
	return c.Status(201).JSON(fiber.Map{"message": "created"})
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var body UserLoginRequest
	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid credentials"})
	}

	token, err := h.svc.Login(body.Email, body.Password)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"error": "invalid credentials"})
	}

	return c.Status(200).JSON(token)
}

func (h *AuthHandler) Logout(c *fiber.Ctx) error {
	var body UserLogoutRequest
	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid body"})
	}
	err := h.svc.Logout(body.RefreshToken)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "token not found"})
	}
	return c.Status(200).JSON(fiber.Map{"message": "logged out"})
}

func (h *AuthHandler) Refresh(c *fiber.Ctx) error {
	var body UserRefreshRequest
	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid body"})
	}

	token, err := h.svc.RefreshToken(body.RefreshToken)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "token not found"})
	}
	return c.JSON(UserRefreshTokenResponse{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
	})
}

func (h *AuthHandler) ValidateToken(c *fiber.Ctx) error {
	var body UserValidateTokenRequest
	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid body"})
	}
	user, err := h.svc.ValidateToken(body.AccessToken)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"error": "Unauthorized"})
	}
	return c.Status(200).JSON(UserValidateTokenResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	})
}
