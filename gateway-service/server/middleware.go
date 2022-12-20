package server

import (
	"net/http"
	"strings"
)

func (s *Server) maxBytes(maxBytes func(r *http.Request) int, h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes(r)))

		h.ServeHTTP(w, r)
	})
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
