package restapi

import (
	"shortlink-service/internal/domain/service"

	"github.com/gofiber/fiber/v2"
)

func RegisterRouters(r fiber.Router, svc service.LinkService) {
	h := NewLinkHandler(svc)

	r.Get("/:code", h.Redirect)

	v1 := r.Group("/api/v1")
	link := v1.Group("/link")
	link.Post("/shorten", h.Create)
	link.Delete("/:code", h.Delete)
}
