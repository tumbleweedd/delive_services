package app

import (
	grpcapp "github.com/tumbleweedd/delive_services/sso/internal/app/grpc"
	"github.com/tumbleweedd/delive_services/sso/internal/repository"
	"github.com/tumbleweedd/delive_services/sso/internal/services"
	"github.com/tumbleweedd/delive_services/sso/pkg/databases/postgres"
	"log/slog"
	"time"
)

type App struct {
	GRPCServer *grpcapp.App
}

func NewApp(
	log *slog.Logger,
	grpcPort int,
	storagePath string,
	tokenTTL time.Duration,
) *App {
	postgresDB, err := postgres.NewPostgresDB(storagePath)
	if err != nil {
		//return nil
	}

	repo := repository.NewRepository(postgresDB)

	svc := services.NewService(log, repo, repo, tokenTTL)

	grpcApp := grpcapp.NewApp(log, svc, grpcPort)

	return &App{
		GRPCServer: grpcApp,
	}
}
