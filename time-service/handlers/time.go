package handlers

import (
	"encoding/json"
	"net/http"
	"time"
)

// TimeResponse is the JSON response for the /time endpoint.
type TimeResponse struct {
	Now string `json:"now"`
}

// ErrorResponse is the JSON response for errors.
type ErrorResponse struct {
	Error string `json:"error"`
}

// TimeHandler handles GET /time requests.
func TimeHandler(w http.ResponseWriter, r *http.Request) {
	now := time.Now().UTC().Format(time.RFC3339)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(TimeResponse{Now: now}); err != nil {
		http.Error(w, `{"error":"Internal server error"}`, http.StatusInternalServerError)
		return
	}
}
