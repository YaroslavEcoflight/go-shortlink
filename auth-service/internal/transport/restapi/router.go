package restapi

import (
	"auth-service/internal/domain/service"

	"github.com/gofiber/fiber/v2"
)

func RegisterRouters(r fiber.Router, svc service.AuthService) {
	h := AuthHandler{svc: svc}
	auth := r.Group("/auth")
	auth.Post("", h.Register)
	auth.Post("/login", h.Login)
	auth.Post("/logout", h.Logout)
	auth.Post("/refresh", h.Refresh)
	auth.Post("/validate", h.ValidateToken)
}
