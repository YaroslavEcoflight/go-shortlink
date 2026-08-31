package app

import (
	"log"
	"shortlink-service/config"
	"shortlink-service/internal/infrastructure/postgres"
	inframodels "shortlink-service/internal/infrastructure/postgres/models"
	"shortlink-service/internal/transport/restapi"
	"shortlink-service/internal/usecase"

	"github.com/gofiber/fiber/v2"
)

func Run() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	db, err := postgres.New(cfg.Pg.DSN())
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&inframodels.Link{})

	linkRepo := postgres.NewLinkRepo(db)
	linkSvc := usecase.NewLinkService(&linkRepo)

	app := fiber.New()
	restapi.RegisterRouters(app, linkSvc)

	log.Fatal(app.Listen(":" + cfg.Http.Port))
}
