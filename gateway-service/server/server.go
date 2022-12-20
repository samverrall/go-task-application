package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
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
		Handler:      s.allowCORS(s.withLogger(mux)),
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

func (s *Server) withLogger(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s.logger.Info("Gateway: URL: %s, Method: %s", r.Method, r.URL)

		h.ServeHTTP(w, r)
	})
}

func (s *Server) allowCORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if origin := r.Header.Get("Origin"); origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			if r.Method == http.MethodOptions && r.Header.Get("Access-Control-Request-Method") != "" {
				s.preflightHandler(w, r)

				return
			}
		}
		h.ServeHTTP(w, r)
	})
}

func (s *Server) preflightHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	headers := []string{"*"}
	w.Header().Set("Access-Control-Allow-Headers", strings.Join(headers, ","))

	methods := []string{http.MethodGet, http.MethodHead, http.MethodPost, http.MethodPut, http.MethodDelete}
	w.Header().Set("Access-Control-Allow-Methods", strings.Join(methods, ","))

	s.logger.Info("preflight request for %s", r.URL.Path)
}
