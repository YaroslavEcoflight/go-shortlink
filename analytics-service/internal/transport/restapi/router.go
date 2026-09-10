package restapi

import (
	"analytics-service/internal/domain"
	"analytics-service/internal/domain/usecase"
	"analytics-service/internal/transport/restapi/middleware"

	"github.com/gofiber/fiber/v2"
)

func RegisterRouters(r fiber.Router, svc usecase.EventUsecase, jwtSecret string, log domain.Interface) {
	h := NewHandler(svc, log)
	v1 := r.Group("/api/v1")
	v1.Use(middleware.Auth(jwtSecret))
	v1.Get("/events", h.GetUserDayEvents)
}
