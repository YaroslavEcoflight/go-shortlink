package app

import (
	"net"

	"analytics-service/config"
	"analytics-service/internal/infrastructure/logger"
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
		panic(err)
	}

	log := logger.New(cfg.App.LogLevel)
	log.Info("starting %s v%s", cfg.App.Name, cfg.App.Version)

	db, err := postgres.New(cfg.Pg.DSN())
	if err != nil {
		log.Fatal("postgres: %v", err)
	}
	db.AutoMigrate(&inframodels.Event{})

	eventRepo := postgres.NewEventRepo(db)
	uc := usecase.NewAnalyticsUsecase(eventRepo)

	handler := transportgrpc.NewHandler(uc, log)
	srv := transportgrpc.NewServer(handler)

	lis, err := net.Listen("tcp", ":"+cfg.Grpc.Port)
	if err != nil {
		log.Fatal("grpc listener: %v", err)
	}

	go func() {
		log.Info("gRPC server listening on :%s", cfg.Grpc.Port)
		if err := srv.Serve(lis); err != nil && err != grpc.ErrServerStopped {
			log.Fatal("gRPC server: %v", err)
		}
	}()

	app := fiber.New()
	restapi.RegisterRouters(app, uc, cfg.Secret.JWTSecret, log)

	log.Info("HTTP server listening on :%s", cfg.Http.Port)
	if err := app.Listen(":" + cfg.Http.Port); err != nil {
		log.Fatal("HTTP server: %v", err)
	}
}
