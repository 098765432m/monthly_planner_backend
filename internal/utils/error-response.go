package error_response

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func WriteError(w http.ResponseWriter, status int, msg string) {
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(ErrorResponse{
		Message: msg,
		Status:  status,
	})
}
