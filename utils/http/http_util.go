package httputil

import (
	"encoding/json"
	"net/http"
)

// WriteJSON encodes a HTTP JSON response with the supplied status and
// response body.
func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}
