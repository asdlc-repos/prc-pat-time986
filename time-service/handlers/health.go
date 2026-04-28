package handlers

import (
	"encoding/json"
	"net/http"
)

// HealthResponse is the JSON response for the /health endpoint.
type HealthResponse struct {
	Status string `json:"status"`
}

// HealthHandler handles GET /health requests.
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(HealthResponse{Status: "ok"}) //nolint:errcheck
}
