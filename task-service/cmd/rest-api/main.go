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
	"github.com/samverrall/task-service/internal/adapters/right/task/sqlite"
	"github.com/samverrall/task-service/internal/port/service/task"
	"github.com/samverrall/task-service/pkg/config"
	sqliteDB "github.com/samverrall/task-service/pkg/sqlite"
)

func run() error {
	c := config.New()
	log := logger.New("debug")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Init SQLite
	dbConnection, err := sqliteDB.Connect(c.DBDirectory)
	if err != nil {
		return err
	}

	if err := dbConnection.Migrate(); err != nil {
		return fmt.Errorf("%w: failed to migrate database", err)
	}

	// Init driven DB (right adapter)
	taskRepo := sqlite.NewTaskRepo(dbConnection.GetDB())

	// Init business logic
	taskService := task.NewService(taskRepo, log)

	// Init driving adapter (left adapter)
	rest := rest.New(c.Address, taskService)
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
