package app

import (
	"context"
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
) (*App, error) {
	postgresDB, err := postgres.NewPostgresDB(context.Background(), storagePath)
	if err != nil {
		log.Error("failed to connect to postgres", err)
		return nil, err
	}

	repo := repository.NewRepository(postgresDB)

	svc := services.NewService(log, tokenTTL, repo)

	grpcApp := grpcapp.NewApp(log, svc, grpcPort)

	return &App{
		GRPCServer: grpcApp,
	}, nil
}
