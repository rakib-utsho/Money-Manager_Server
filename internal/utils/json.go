package utils

import (
	"encoding/json"
	"net/http"
)

type ApiResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func WriteJSON(
	w http.ResponseWriter,
	status int,
	success bool,
	message string,
	data interface{},
) {
	response := ApiResponse{
		Success: success,
		Message: message,
		Data:    data,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(response)
}
