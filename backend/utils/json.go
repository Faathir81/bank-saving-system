package utils

import (
	"encoding/json"
	"net/http"
)

func SendJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

func SendError(w http.ResponseWriter, status int, message string, err string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  "error",
		"message": message,
		"data":    err,
	})
}
