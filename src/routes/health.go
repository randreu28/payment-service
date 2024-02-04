package routes

import (
	"encoding/json"
	"net/http"
	"time"
)

// Health Returns the current date and status OK
func Health(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)

	jsonRes := map[string]string{
		"timestamp": time.Now().Format(time.RFC3339),
		"status":    "OK",
	}

	json.NewEncoder(res).Encode(jsonRes)
}
