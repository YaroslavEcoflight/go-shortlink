package restapi

import (
	"shortlink-service/internal/domain/service"

	"github.com/gofiber/fiber/v2"
)

func RegisterRouters(r fiber.Router, svc service.LinkService) {
	h := NewLinkHandler(svc)
	link := r.Group("/link")
	link.Post("/shorten", h.Create)
	link.Delete(":code", h.Delete)
}

