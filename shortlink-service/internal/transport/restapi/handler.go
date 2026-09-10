package restapi

import (
	"shortlink-service/internal/domain"
	"shortlink-service/internal/domain/entity"
	"shortlink-service/internal/domain/service"
	"shortlink-service/internal/infrastructure/analytics"
	pb "shortlink-service/proto/analytics"

	"github.com/gofiber/fiber/v2"
)

type LinkHandler struct {
	svc       service.LinkService
	analytics *analytics.Client
	log       domain.Interface
}

func NewLinkHandler(svc service.LinkService, ac *analytics.Client, log domain.Interface) LinkHandler {
	return LinkHandler{svc: svc, analytics: ac, log: log}
}

func (h *LinkHandler) Create(c *fiber.Ctx) error {
	var body CreateRequest
	userID := c.Locals("userID").(string)
	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid body"})
	}
	n := entity.Link{
		Url: body.Url,
	}
	obj, err := h.svc.Create(n)
	if err != nil {
		h.log.Error("failed to create link: %v", err)
		return c.Status(400).JSON(fiber.Map{"error": "Invalid"})
	}
	h.log.Info("link created: code=%s user=%s", obj.Code, userID)
	h.analytics.RecordEvent("shorten", pb.EventStatus_STATUS_SUCCESS, userID, c.IP(), "")
	return c.Status(201).JSON(obj)
}

func (h *LinkHandler) Delete(c *fiber.Ctx) error {
	code := c.Params("code")
	userID := c.Locals("userID").(string)
	if err := h.svc.Delete(code); err != nil {
		h.log.Error("failed to delete link code=%s: %v", code, err)
		return c.Status(404).JSON(fiber.Map{"error": "code not found"})
	}
	h.log.Info("link deleted: code=%s user=%s", code, userID)
	h.analytics.RecordEvent("delete", pb.EventStatus_STATUS_SUCCESS, userID, c.IP(), "")
	return c.Status(200).JSON(fiber.Map{"message": "code deleted"})
}

func (h *LinkHandler) Redirect(c *fiber.Ctx) error {
	code := c.Params("code")
	link, err := h.svc.GetByCode(code)
	if err != nil {
		h.log.Debug("redirect: code=%s not found", code)
		return c.Status(404).JSON(fiber.Map{"error": "not found"})
	}
	return c.Redirect(link.Url, 302)
}
