package restapi

import (
	"analytics-service/internal/domain/usecase"

	"github.com/gofiber/fiber/v2"
)

func RegisterRouters(r fiber.Router, svc usecase.EventUsecase) {
	h := NewHandler(svc)
	v1 := r.Group("/api/v1")
	v1.Get("/events", h.GetUserDayEvents)
}
