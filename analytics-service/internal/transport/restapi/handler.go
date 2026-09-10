package restapi

import (
	"time"

	"analytics-service/internal/domain/usecase"

	"github.com/gofiber/fiber/v2"
)

type AnalyticsEventHandler struct {
	svc usecase.EventUsecase
}

func NewHandler(svc usecase.EventUsecase) *AnalyticsEventHandler {
	return &AnalyticsEventHandler{svc: svc}
}

// GET /api/v1/events?user_id=123&date=2026-09-10
func (h *AnalyticsEventHandler) GetUserDayEvents(c *fiber.Ctx) error {
	userID := c.Query("user_id")
	if userID == "" {
		return c.Status(400).JSON(fiber.Map{"error": "user_id is required"})
	}

	dateStr := c.Query("date")
	if dateStr == "" {
		return c.Status(400).JSON(fiber.Map{"error": "date is required"})
	}

	day, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid date format, use YYYY-MM-DD"})
	}

	events, err := h.svc.GetUserDayEvents(c.Context(), userID, day)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "internal error"})
	}

	return c.Status(200).JSON(events)
}
