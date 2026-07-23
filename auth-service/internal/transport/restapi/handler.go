package restapi

import (
	"auth-service/internal/domain/service"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	svc service.AuthSerivce
}

func (h *AuthHandler) Register(c *fiber.Ctx) error {
	var body UserRegisterRequest
	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid body"})
	}
	return c.Status(201).JSON(fiber.Map{"message": "created"})
}
