package healthservice

import (
	"encoding/json"
	"net/http"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx := r.Body
	res := map[string]interface{}{
		"status":  "ok",
		"message": "Service is running",
		"context": ctx,
	}
	json.NewEncoder(w).Encode(res)
}
