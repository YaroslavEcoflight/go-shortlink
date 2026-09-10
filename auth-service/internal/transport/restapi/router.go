package restapi

import (
	"auth-service/internal/domain/service"
	"auth-service/internal/infrastructure/analytics"

	"github.com/gofiber/fiber/v2"
)

func RegisterRouters(r fiber.Router, svc service.AuthService, ac *analytics.Client) {
	h := NewAuthHandler(svc, ac)
	auth := r.Group("/auth")
	auth.Post("", h.Register)
	auth.Post("/login", h.Login)
	auth.Post("/logout", h.Logout)
	auth.Post("/refresh", h.Refresh)
	auth.Post("/validate", h.ValidateToken)
}
