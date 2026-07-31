package app

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"

	"auth-service/config"
	"auth-service/internal/infrastructure/postgres"
	"auth-service/internal/infrastructure/redis"
	"auth-service/internal/transport/restapi"
	"auth-service/internal/usecase"
)

func Run() error {
	ctx := context.Background()

	cfg, err := config.NewConfig()
	if err != nil {
		return fmt.Errorf("config: %w", err)
	}

	_, err = postgres.New(cfg.Pg.DSN())
	if err != nil {
		return fmt.Errorf("postgres: %w", err)
	}

	_, err = redis.New(ctx, *cfg)
	if err != nil {
		return fmt.Errorf("redis: %w", err)
	}

	app := fiber.New()

	svc := usecase.NewAuthService(nil) // TODO: передать репозиторий

	api := app.Group("/api/v1")
	restapi.RegisterRouters(api, svc)
	return app.Listen(":" + cfg.HTTP.Port)
}
