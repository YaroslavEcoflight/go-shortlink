package restapi

import (
	"shortlink-service/config"
	"shortlink-service/internal/domain/service"
	"shortlink-service/internal/transport/restapi/middleware"

	"github.com/gofiber/fiber/v2"
)

func RegisterRouters(r fiber.Router, svc service.LinkService, cfg config.Config) {
	h := NewLinkHandler(svc)

	r.Get("/:code", h.Redirect)

	v1 := r.Group("/api/v1")
	link := v1.Group("/link")
	link.Use(middleware.Auth(cfg.Secret.JWTSecret))
	link.Post("/shorten", h.Create)
	link.Delete("/:code", h.Delete)
}
