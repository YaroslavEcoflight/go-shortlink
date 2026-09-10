package app

import (
	"context"
	"log"
	"shortlink-service/config"
	infraAnalytics "shortlink-service/internal/infrastructure/analytics"
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
		log.Fatal(err)
	}

	analyticsClient, err := infraAnalytics.NewClient(cfg.Analytics.Addr)
	if err != nil {
		log.Fatal(err)
	}

	rdb, err := redis.New(ctx, *cfg)
	if err != nil {
		log.Fatal(err)
	}

	db, err := postgres.New(cfg.Pg.DSN())
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&inframodels.Link{})

	linkCache := redis.NewLinkRepos(ctx, rdb)
	linkRepo := postgres.NewLinkRepo(db)
	linkSvc := usecase.NewLinkService(&linkRepo, linkCache)

	app := fiber.New()
	restapi.RegisterRouters(app, linkSvc, *cfg, analyticsClient)

	log.Fatal(app.Listen(":" + cfg.Http.Port))
}
