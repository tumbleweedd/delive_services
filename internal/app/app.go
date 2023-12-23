package app

import (
	"context"
	httpapp "github.com/tumbleweedd/delive_services/sso/internal/app/http"
	"github.com/tumbleweedd/delive_services/sso/internal/repository"
	"github.com/tumbleweedd/delive_services/sso/internal/services"
	"github.com/tumbleweedd/delive_services/sso/pkg/databases/postgres"
	"log/slog"
	"time"
)

type App struct {
	HTTPServer *httpapp.App
}

func NewApp(
	log *slog.Logger,
	httpPort int,
	storagePath string,
	tokenTTL time.Duration,
) (*App, error) {
	postgresDB, err := postgres.NewPostgresDB(context.Background(), storagePath)
	if err != nil {
		log.Error("failed to connect to postgres", err)
		return nil, err
	}

	repo := repository.NewRepository(postgresDB)

	svc := services.NewServices(log, tokenTTL, repo)

	httpApp := httpapp.NewApp(log, httpPort, svc)

	return &App{
		HTTPServer: httpApp,
	}, nil
}
