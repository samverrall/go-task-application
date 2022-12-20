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
	address string
	server  *http.Server
	logger  logger.Logger
}

func New(logger logger.Logger, host string, port int, mux *http.ServeMux) *Server {
	address := fmt.Sprintf("%s:%d", host, port)

	s := &Server{
		logger:  logger,
		address: address,
	}
	s.server = &http.Server{
		Handler: s.maxBytes(func(r *http.Request) int {
			switch r.Method {
			case http.MethodPatch, http.MethodPost, http.MethodPut:
				return 10
			default:
				return 0
			}
		}, s.allowCORS(s.withLogger(mux))),
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
