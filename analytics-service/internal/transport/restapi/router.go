package restapi

import (
	"analytics-service/internal/domain/usecase"
	"analytics-service/internal/transport/restapi/middleware"

	"github.com/gofiber/fiber/v2"
)

func RegisterRouters(r fiber.Router, svc usecase.EventUsecase, jwtSecret string) {
	h := NewHandler(svc)
	v1 := r.Group("/api/v1")
	v1.Use(middleware.Auth(jwtSecret))
	v1.Get("/events", h.GetUserDayEvents)
}
