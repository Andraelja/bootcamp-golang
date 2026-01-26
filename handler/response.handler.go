package handler

import (
	"encoding/json"
	"net/http"
	"task-session-1/entity"
)

func Success(w http.ResponseWriter, status int, message string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	_ = json.NewEncoder(w).Encode(entity.APIResponse{
		Success: true,
		Message: message,
		Data:    data,
	})
}

func Error(w http.ResponseWriter, status int, message string, err interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	_ = json.NewEncoder(w).Encode(entity.APIResponse{
		Success: false,
		Message: message,
		Error:   err,
	})
}
