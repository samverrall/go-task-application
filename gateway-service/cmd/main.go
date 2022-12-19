package main

import (
	"context"
	"errors"
	"flag"
	"net/http"

	"github.com/samverrall/go-task-application/gateway-service/gateway"
	"github.com/samverrall/go-task-application/gateway-service/server"
	"github.com/samverrall/go-task-application/logger"
)

func main() {
	var host string
	var port int
	flag.StringVar(&host, "host", "127.0.0.1", "Host to start the gateway HTTP server on")
	flag.IntVar(&port, "port", 5000, "Port to start gateway HTTP server on")
	flag.Parse()

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	log := logger.New("info")

	gw := gateway.New(log)

	gatewayHandler, err := gw.Handler(ctx, nil)
	if err != nil {
		log.Fatal("failed to create gateway handler: %s", err.Error())
	}

	mux := http.NewServeMux()

	mux.Handle("/", gatewayHandler)

	s := server.New(log, host, port, mux)

	go func() {
		<-ctx.Done()
		log.Info("Shutting down the http server")

		if err := s.Shutdown(context.Background()); err != nil {
			log.Error("Failed to shutdown http server: %v", err)
		}
	}()

	if err := s.Start(ctx); errors.Is(err, http.ErrServerClosed) {
		log.Error("Failed to listen and serve: %v", err)
	}
}
