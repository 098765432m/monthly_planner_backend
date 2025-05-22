package utils

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func WriteError(w http.ResponseWriter, status int, msg string) {
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(ErrorResponse{
		Success: false,
		Message: msg,
	})
}

type SuccessResponse struct {
	Success bool   `json:"success,omitempty"`
	Message string `json:"message,omitempty"`
	Status  int    `json:"status"`
	Data    any    `json:"data,omitempty"`
}

func WriteJSON(w http.ResponseWriter, statusCode int, data any, msg string) {
	w.Header().Set("Content-Type/", "application/json")
	w.WriteHeader(statusCode)

	resp := SuccessResponse{
		Success: true,
		Message: msg,
		Data:    data,
	}

	json.NewEncoder(w).Encode(resp)
}
