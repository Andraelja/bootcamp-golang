package handler

import (
	// "task-session-1/entity"
	"task-session-1/service"
	// "task-session-1/helper"
	"net/http"
)

func HealtHeandler(w http.ResponseWriter, r *http.Request) {
	Success(w, http.StatusOK, "API Running!", nil)
}

func getAllCategory(w http.ResponseWriter, r *http.Request) {
	Success(w, http.StatusOK, "Get all categories", service.GetAllCategory())
}