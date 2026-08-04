package restapi

import (
	"auth-service/internal/domain/service"

	"github.com/gofiber/fiber/v2"
)

func RegisterRouters(r fiber.Router, svc service.AuthSerivce) {
	h := AuthHandler{svc: svc}
	auth := r.Group("/auth")
	auth.Post("", h.Register)
	auth.Post("/login", h.Login)
}
