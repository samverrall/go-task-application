package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/samverrall/go-task-application/logger"
	"github.com/samverrall/go-task-application/user-service/internal/adapters/left/grpc"
	"github.com/samverrall/go-task-application/user-service/internal/adapters/right/repo/user/sqlite"
	"github.com/samverrall/go-task-application/user-service/internal/port/service/user"
	sqliteconn "github.com/samverrall/go-task-application/user-service/pkg/sqlite"
)

func run(grpcPort int, database string) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	log := logger.New("debug")

	sqliteAdapter, err := sqliteconn.Connect(database)
	if err != nil {
		return fmt.Errorf("%w: failed to connect to sqlite adapter", err)
	}

	// Init right sqlite repo
	userRepo, err := sqlite.NewUserRepo(sqliteAdapter.GetDB())
	if err != nil {
		log.Error("failed to init new sqlite user repo: %w", err)
		return err
	}

	// Init business logic
	userSvc := user.NewService(userRepo, log)

	// Init gRPC adapter and inject business logic
	grpcAdapter := grpc.New(userSvc, log, grpcPort)
	if err := grpcAdapter.Run(); err != nil {
		return err
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	select {
	case v := <-quit:
		log.Error("signal.Notify: %v", v)
	case done := <-ctx.Done():
		log.Error("ctx.Done: %v", done)
	}

	return nil
}

func main() {
	if err := run(8000, "users.db"); err != nil {
		log.Fatal(err)
	}
}
