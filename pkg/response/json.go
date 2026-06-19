// Package response provides utility functions for sending HTTP responses in JSON format.
package response

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status  bool   `json:"status"`
	Message string `json:"message,omitempty"`
	Error   string `json:"error,omitempty"`
	Data    any    `json:"data,omitempty"`
}

// JSON writes a JSON response with the given status code and data.
func JSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

// Error writes a JSON error response with the given status code and error message.
func Error(w http.ResponseWriter, status int, err error) {
	response := Response{
		Status:  false,
		Message: http.StatusText(status),
		Error:   err.Error(),
	}
	JSON(w, status, response)
}

func Success(w http.ResponseWriter, status int, message string, data any) {
	response := Response{
		Status:  true,
		Message: message,
		Data:    data,
	}
	JSON(w, status, response)
}
