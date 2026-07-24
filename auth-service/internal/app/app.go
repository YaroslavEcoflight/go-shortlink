package app

import (
	"github.com/gofiber/fiber/v2"

	"auth-service/internal/transport/restapi"
	"auth-service/internal/usecase"
)

func Run() {
	app := fiber.New()

	svc := usecase.NewAuthService(nil) // TODO: передать репозиторий

	api := app.Group("/api/v1")
	restapi.RegisterRouters(api, svc)
	app.Listen(":3000")
}
