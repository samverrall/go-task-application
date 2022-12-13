package app

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/samverrall/go-task-application/logger"
	"github.com/samverrall/go-task-application/user-service/internal/grpc"
	"github.com/samverrall/go-task-application/user-service/internal/repository"
	"github.com/samverrall/go-task-application/user-service/internal/service"
	"github.com/samverrall/go-task-application/user-service/internal/sqlite"
)

type App struct {
	logger logger.Logger
}

func New(log logger.Logger) *App {
	return &App{
		logger: log,
	}
}

func (a *App) Run() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Init DB adapter
	sqliteAdapter, err := sqlite.Connect("")
	if err != nil {
		return fmt.Errorf("%w: failed to connect to sqlite adapter", err)
	}

	// Init Repos
	userRepo := repository.NewUserRepo(sqliteAdapter.GetDB())

	// Init business logic
	userSvc := service.NewUserService(userRepo, a.logger)

	// Init gRPC adapter and inject business logic
	grpcAdapter := grpc.New(userSvc, a.logger, 8000)
	if err := grpcAdapter.Run(); err != nil {
		return err
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	select {
	case v := <-quit:
		a.logger.Error("signal.Notify: %v", v)
	case done := <-ctx.Done():
		a.logger.Error("ctx.Done: %v", done)
	}

	return nil
}
