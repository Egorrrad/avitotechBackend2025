package http

import (
	"avitotechBackend2025/internal/dto"
	"encoding/json"
	"net/http"
)

func SendError(w http.ResponseWriter, code int, message string) {
	RespondJSON(w, code, dto.Error{Message: message})
}

func RespondJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
