package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/samverrall/go-task-application/logger"
)

type Server struct {
	address     string
	server      *http.Server
	logger      logger.Logger
	middlewares []MiddlewareFunc
}

func New(logger logger.Logger, host string, port int, gatewayHandler http.Handler) *Server {
	address := fmt.Sprintf("%s:%d", host, port)

	s := &Server{
		logger:  logger,
		address: address,
	}

	mux := http.NewServeMux()

	// Register the gateway handler on the root so that request paths can be
	// forwarded to our gRPC proxy.
	mux.Handle("/", gatewayHandler)

	// Health endpoint
	mux.HandleFunc("/healthz", s.errorHandlerFunc(s.healthHandler))

	s.useMiddleware(s.withLogger)
	s.useMiddleware(s.allowCORS)
	s.useMiddleware(s.maxBytes(func(r *http.Request) int {
		switch r.Method {
		case http.MethodPatch, http.MethodPost, http.MethodPut:
			kilobyte := 1000 * 1
			return 100 * kilobyte

		default:
			return 0
		}
	}))

	s.server = &http.Server{
		Handler:      s.applyMiddleware(mux),
		Addr:         address,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	return s
}

// Start starts the HTTP server and listens on the server's address.
func (s *Server) Start(ctx context.Context) error {
	if err := s.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		s.logger.Error("Failed to listen and serve: %v", err)
		return err
	}
	s.logger.Info("Gateway HTTP server listening on %s", s.address)
	return nil
}

func (s *Server) Shutdown(ctx context.Context) error {
	if err := s.server.Shutdown(ctx); err != nil {
		s.logger.Error("Failed to shutdown http server: %v", err)
		return err
	}
	return nil
}
