package router

import (
	"task-session-1/handler"
	"net/http"
)

func RegisterRoutes() {
	http.HandleFunc("/health", handler.HealtHeandler)
}