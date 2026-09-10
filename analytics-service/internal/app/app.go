package app

import (
	"log"
	"net"

	"analytics-service/config"
	"analytics-service/internal/infrastructure/postgres"
	inframodels "analytics-service/internal/infrastructure/postgres/models"
	transportgrpc "analytics-service/internal/transport/grpc"
	"analytics-service/internal/transport/restapi"
	"analytics-service/internal/usecase"

	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc"
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
	db.AutoMigrate(&inframodels.Event{})

	eventRepo := postgres.NewEventRepo(db)
	uc := usecase.NewAnalyticsUsecase(eventRepo)

	handler := transportgrpc.NewHandler(uc)
	srv := transportgrpc.NewServer(handler)

	lis, err := net.Listen("tcp", ":"+cfg.Grpc.Port)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		log.Printf("gRPC server listening on :%s", cfg.Grpc.Port)
		if err := srv.Serve(lis); err != nil && err != grpc.ErrServerStopped {
			log.Fatal(err)
		}
	}()

	app := fiber.New()
	restapi.RegisterRouters(app, uc, cfg.Secret.JWTSecret)

	log.Printf("HTTP server listening on :%s", cfg.Http.Port)
	log.Fatal(app.Listen(":" + cfg.Http.Port))
}
