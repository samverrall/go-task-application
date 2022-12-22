package server

import (
	"net/http"
	"strings"
)

type MiddlewareFunc func(h http.HandlerFunc) http.HandlerFunc

func (s *Server) useMiddleware(middleware MiddlewareFunc) {
	s.middlewares = append(s.middlewares, middleware)
}

func (s *Server) applyMiddleware(handler http.Handler) http.Handler {
	for i := 0; i < len(s.middlewares); i++ {
		middleware := s.middlewares[i]

		handler = http.HandlerFunc(middleware(handler.ServeHTTP))
	}

	return handler
}

func (s *Server) maxBytes(maxBytes func(r *http.Request) int) MiddlewareFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes(r)))

			next(w, r)
		}
	}
}

func (s *Server) withLogger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.logger.Info("Gateway: URL: %s, Method: %s", r.Method, r.URL)

		next(w, r)
	}
}

func (s *Server) allowCORS(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if origin := r.Header.Get("Origin"); origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			if r.Method == http.MethodOptions && r.Header.Get("Access-Control-Request-Method") != "" {
				s.preflightHandler(w, r)

				return
			}
		}
		h.ServeHTTP(w, r)
	}
}

func (s *Server) preflightHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	headers := []string{"*"}
	w.Header().Set("Access-Control-Allow-Headers", strings.Join(headers, ","))

	methods := []string{http.MethodGet, http.MethodHead, http.MethodPost, http.MethodPut, http.MethodDelete}
	w.Header().Set("Access-Control-Allow-Methods", strings.Join(methods, ","))

	s.logger.Info("preflight request for %s", r.URL.Path)
}
