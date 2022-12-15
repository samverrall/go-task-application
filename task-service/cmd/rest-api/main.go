package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/samverrall/go-task-application/logger"
	"github.com/samverrall/task-service/internal/adapters/left/rest"
	sqliteAdapter "github.com/samverrall/task-service/internal/adapters/right/sqlite"
	"github.com/samverrall/task-service/internal/repository/sqlite"
	"github.com/samverrall/task-service/internal/service/task"
	"github.com/samverrall/task-service/pkg/config"
)

func run() error {
	c := config.New()
	log := logger.New("debug")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Init DB adapter
	dbConnection, err := sqliteAdapter.Connect(c.DBDirectory)
	if err != nil {
		return err
	}

	if err := dbConnection.Migrate(); err != nil {
		return fmt.Errorf("%w: failed to migrate database", err)
	}

	// Init repos
	taskRepo := sqlite.NewTaskRepo(dbConnection.GetDB())

	// Init business logic
	taskService := task.NewService(taskRepo, log)

	// Init rest adapter
	rest := rest.New(c.Address)
	rest.InitMiddleware()
	rest.InitHandlers(ctx, taskService)
	if err := rest.Start(ctx); err != nil {
		return fmt.Errorf("%w: failed to start http server", err)
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	select {
	case v := <-quit:
		log.Error("signal.Notify: %v", v)
	case done := <-ctx.Done():
		log.Error("ctx.Done: %v", done)
	}

	log.Info("REST HTTP server listening at %s", c.Address)

	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
