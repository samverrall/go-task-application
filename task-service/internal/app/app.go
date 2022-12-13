package app

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/samverrall/go-task-application/logger"
	"github.com/samverrall/task-service/internal/repository"
	"github.com/samverrall/task-service/internal/rest"
	"github.com/samverrall/task-service/internal/service"
	"github.com/samverrall/task-service/internal/sqlite"
	"github.com/samverrall/task-service/pkg/config"
)

type App struct {
	logger logger.Logger
	config *config.Config
}

func New(logger logger.Logger, config *config.Config) *App {
	return &App{
		logger: logger,
		config: config,
	}
}

func (a *App) Run() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Init DB adapter
	dbConnection, err := sqlite.Connect(a.config.DBDirectory)
	if err != nil {
		return err
	}

	if err := dbConnection.Migrate(); err != nil {
		return fmt.Errorf("%w: failed to migrate database", err)
	}

	// Init repos
	taskRepo := repository.NewTaskRepo(dbConnection.GetDB())

	// Init business logic
	taskService := service.NewTaskService(taskRepo, a.logger)

	// Init rest adapter
	rest := rest.New(a.config.Address)
	rest.InitMiddleware()
	rest.InitHandlers(ctx, taskService)
	if err := rest.Start(ctx); err != nil {
		return fmt.Errorf("%w: failed to start http server", err)
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	select {
	case v := <-quit:
		a.logger.Error("signal.Notify: %v", v)
	case done := <-ctx.Done():
		a.logger.Error("ctx.Done: %v", done)
	}

	a.logger.Info("REST HTTP server listening at %s", a.config.Address)

	return nil
}
