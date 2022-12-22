package server

import (
	"net/http"

	httputil "github.com/samverrall/go-task-application/utils/http"
)

type APIError struct {
	Error string `json:"error"`
}

func NewAPIError(err string) APIError {
	return APIError{
		Error: err,
	}
}

type apiErrorFunc func(http.ResponseWriter, *http.Request) error

func (s *Server) errorHandlerFunc(f apiErrorFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			s.logger.LogError(err)
			httputil.WriteJSON(w, http.StatusInternalServerError, NewAPIError(err.Error()))
		}
	}
}
