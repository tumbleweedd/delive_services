package grpc

import (
	"fmt"
	authgrpc "github.com/tumbleweedd/delive_services/sso/internal/grpc"
	"github.com/tumbleweedd/delive_services/sso/internal/grpc/auth_server"
	"google.golang.org/grpc"
	"log/slog"
	"net"
)

type App struct {
	log        *slog.Logger
	gRPCServer *grpc.Server
	port       int
}

func NewApp(log *slog.Logger, authService auth_server.AuthService, port int) *App {
	gRPCServer := grpc.NewServer()

	authgrpc.Register(gRPCServer, authService)

	return &App{
		log:        log,
		gRPCServer: gRPCServer,
		port:       port,
	}
}

func (a *App) RunWithPanic() {
	if err := a.run(); err != nil {
		panic("failed to run grpc server")
	}
}

func (a *App) run() error {
	const grpcAppRunInfo = "grpcapp.run"
	log := a.log.With(
		slog.String("grpcAppRunInfo", grpcAppRunInfo),
		slog.Int("port", a.port),
	)
	log.Info("starting gRPC server")

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		return fmt.Errorf("%s: %w", grpcAppRunInfo, err)
	}

	log.Info("grpc server is running", slog.String("addr", l.Addr().String()))

	if err := a.gRPCServer.Serve(l); err != nil {
		return fmt.Errorf("%s: %w", grpcAppRunInfo, err)
	}

	return nil
}

func (a *App) Stop() {
	const grpcAppStopInfo = "grpcapp.run"
	a.log.With(slog.String("grpcAppStopInfo", grpcAppStopInfo)).Info("stopping grpc server")

	a.gRPCServer.GracefulStop()

}
