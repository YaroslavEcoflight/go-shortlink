package app

import (
	"context"
	"shortlink-service/config"
	infraAnalytics "shortlink-service/internal/infrastructure/analytics"
	"shortlink-service/internal/infrastructure/logger"
	"shortlink-service/internal/infrastructure/postgres"
	inframodels "shortlink-service/internal/infrastructure/postgres/models"
	"shortlink-service/internal/infrastructure/redis"
	"shortlink-service/internal/transport/restapi"
	"shortlink-service/internal/usecase"

	"github.com/gofiber/fiber/v2"
)

func Run() {
	ctx := context.Background()

	cfg, err := config.NewConfig()
	if err != nil {
		panic(err)
	}

	log := logger.New(cfg.App.LogLevel)
	log.Info("starting %s v%s", cfg.App.Name, cfg.App.Version)

	analyticsClient, err := infraAnalytics.NewClient(cfg.Analytics.Addr)
	if err != nil {
		log.Fatal("analytics client: %v", err)
	}

	rdb, err := redis.New(ctx, *cfg)
	if err != nil {
		log.Fatal("redis: %v", err)
	}

	db, err := postgres.New(cfg.Pg.DSN())
	if err != nil {
		log.Fatal("postgres: %v", err)
	}
	db.AutoMigrate(&inframodels.Link{})

	linkCache := redis.NewLinkRepos(ctx, rdb)
	linkRepo := postgres.NewLinkRepo(db)
	linkSvc := usecase.NewLinkService(&linkRepo, linkCache)

	app := fiber.New()
	restapi.RegisterRouters(app, linkSvc, *cfg, analyticsClient, log)

	log.Info("HTTP server listening on :%s", cfg.Http.Port)
	if err := app.Listen(":" + cfg.Http.Port); err != nil {
		log.Fatal("HTTP server: %v", err)
	}
}
