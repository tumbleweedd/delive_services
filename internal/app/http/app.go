package http_app

import (
	"context"
	"fmt"
	sso_http "github.com/tumbleweedd/delive_services/sso/internal/delivery/http"
	"github.com/tumbleweedd/delive_services/sso/internal/services"
	"log/slog"
	"net/http"
)

type App struct {
	log        *slog.Logger
	httpServer *http.Server
	port       int
}

func NewApp(log *slog.Logger, port int, svc *services.Services) *App {
	handler := sso_http.NewHandler(log, svc)

	httpServer := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: handler.InitRoutes(),
	}

	return &App{
		log:        log,
		httpServer: httpServer,
		port:       port,
	}
}

func (a *App) RunWithPanic() {
	if err := a.run(); err != nil {
		panic("failed to run http server")
	}
}

func (a *App) run() error {
	const op = "httpapp.run"

	log := a.log.With(
		slog.String("op", op),
		slog.Int("port", a.port),
	)

	log.Info("starting http server")

	return a.httpServer.ListenAndServe()
}

func (a *App) Stop() {
	const op = "httpapp.stop"

	a.log.With(slog.String("op", op)).Info("stopping http server")

	a.httpServer.Shutdown(context.Background())
}
