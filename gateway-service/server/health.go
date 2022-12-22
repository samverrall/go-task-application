package server

import (
	"net/http"

	httputil "github.com/samverrall/go-task-application/utils/http"
)

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) error {
	s.logger.Info("healthHandler Invoked")

	resp := map[string]string{
		"message": "healthy",
	}
	return httputil.WriteJSON(w, http.StatusOK, resp)
}
