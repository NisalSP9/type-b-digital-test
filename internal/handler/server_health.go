package handler

import (
	"encoding/json"
	"net/http"
	"runtime"
	"time"
)

func ServerHealth(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":               "OK",
		"message":              "TypeBDigital test server is up and running",
		"time":                 time.Now().Format(time.RFC3339),
		"go version":           runtime.Version(),
		"number of goroutines": runtime.NumGoroutine(),
	})

}
