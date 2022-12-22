package main

import (
	"context"
	"errors"
	"flag"
	"net/http"
	"time"

	"github.com/samverrall/go-task-application/gateway-service/config"
	"github.com/samverrall/go-task-application/gateway-service/gateway"
	"github.com/samverrall/go-task-application/gateway-service/server"
	"github.com/samverrall/go-task-application/logger"
)

var opts struct {
	log struct {
		level string
	}

	config struct {
		path string
	}

	server struct {
		host string
		port int
	}
}

func main() {
	flag.StringVar(&opts.config.path, "config-path", "../config/local.config.yaml", "Level of output logs written to stdout")
	flag.StringVar(&opts.log.level, "log-level", "info", "Level of output logs written to stdout")
	flag.StringVar(&opts.server.host, "host", "127.0.0.1", "Host to start the gateway HTTP server on")
	flag.IntVar(&opts.server.port, "port", 5000, "Port to start gateway HTTP server on")
	flag.Parse()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	log := logger.New(opts.log.level)

	config := config.New(opts.config.path)
	if err := config.ParseConfig(); err != nil {
		log.Fatal("failed to parse config: %s", err.Error())
	}

	// Initialise new gateway
	gw := gateway.New(log, config)

	// Return the mux gateway handler
	gatewayHandler, err := gw.Handler(ctx, nil)
	if err != nil {
		log.Fatal("failed to create gateway handler: %s", err.Error())
	}

	s := server.New(log, opts.server.host, opts.server.port, gatewayHandler)

	go func() {
		<-ctx.Done()
		log.Info("Shutting down the http server")

		ctxShutdown, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := s.Shutdown(ctxShutdown); err != nil {
			log.Error("Failed to shutdown http server: %v", err)
		}
	}()

	if err := s.Start(ctx); errors.Is(err, http.ErrServerClosed) {
		log.Error("Failed to listen and serve: %v", err)
	}
}
