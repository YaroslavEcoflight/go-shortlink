package app

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"

	"auth-service/config"
	infrapostgres "auth-service/internal/infrastructure/postgres"
	infraredis "auth-service/internal/infrastructure/redis"
	"auth-service/internal/transport/restapi"
	"auth-service/internal/usecase"
)

func Run() error {
	ctx := context.Background()

	cfg, err := config.NewConfig()
	if err != nil {
		return fmt.Errorf("config: %w", err)
	}

	db, err := infrapostgres.New(cfg.Pg.DSN())
	if err != nil {
		return fmt.Errorf("postgres: %w", err)
	}

	rdb, err := infraredis.New(ctx, *cfg)
	if err != nil {
		return fmt.Errorf("redis: %w", err)
	}

	userRepo := infrapostgres.NewRepository(db)
	tokenRepo := infraredis.NewTokenRepo(ctx, rdb)

	app := fiber.New()

	svc := usecase.NewAuthService(userRepo, tokenRepo, cfg.Secret.JWTSecret)

	api := app.Group("/api/v1")
	restapi.RegisterRouters(api, svc)
	return app.Listen(":" + cfg.HTTP.Port)
}
