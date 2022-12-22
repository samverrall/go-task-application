package main

import (
	"context"
	"flag"
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

var opts struct {
	database struct {
		dir string
	}

	server struct {
		port int
		host string
	}
}

func run() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	log := logger.New("info")

	sqliteAdapter, err := sqliteconn.Connect(opts.database.dir)
	if err != nil {
		return fmt.Errorf("%w: failed to connect to sqlite adapter", err)
	}

	// Init right sqlite repo
	userRepo, err := sqlite.NewUserRepo(sqliteAdapter.GetDB())
	if err != nil {
		log.Error("failed to init new sqlite user repo: %v", err)
		return err
	}

	// Init business logic
	userSvc := user.NewService(userRepo, log)

	// Init gRPC adapter and inject business logic
	grpcAdapter := grpc.New(userSvc, log, opts.server.host, opts.server.port)
	if err := grpcAdapter.Run(ctx); err != nil {
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
	flag.StringVar(&opts.server.host, "host", "127.0.0.1", "The host to run the gRPC server on")
	flag.IntVar(&opts.server.port, "port", 8002, "The port to run the gRPC server on")
	flag.StringVar(&opts.database.dir, "database-dir", "./users.db", "The directory store application data")
	flag.Parse()
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
