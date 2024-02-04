package routes

import (
	"encoding/json"
	"net/http"
	"time"
)

// Health Returns the current date and status OK
func Health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := map[string]string{
		"timestamp": time.Now().Format(time.RFC3339),
		"status":    "OK",
	}

	json.NewEncoder(w).Encode(response)
}
