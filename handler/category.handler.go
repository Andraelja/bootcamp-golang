package handler

import (
	// "task-session-1/entity"
	"task-session-1/entity"
	"task-session-1/helper"
	"task-session-1/service"

	// "task-session-1/helper"
	"net/http"
)

func HealtHeandler(w http.ResponseWriter, r *http.Request) {
	Success(w, http.StatusOK, "API Running!", nil)
}

func GetAllCategory(w http.ResponseWriter, r *http.Request) {
	Success(w, http.StatusOK, "Get all categories", service.GetAllCategory())
}

func GetCategoryById(w http.ResponseWriter, r *http.Request) {
	// ubah dulu string id jadi integer
	id, err := helper.ParseID(r, "/api/categories/")
	if err != nil {
		Error(w, http.StatusBadRequest, "Invalid Category ID", err.Error())
		return
	}

	category, err := service.GetCategoryById(id)
	if err != nil {
		Error(w, http.StatusNotFound, "Category Not Found", nil)
		return
	}

	Success(w, http.StatusOK, "Category Found", category)
}

func StoreCategory(w http.ResponseWriter, r *http.Request) {
	var category entity.Category
	if err := helper.DecodeJSON(r, &category); err != nil {
		Error(w, http.StatusBadRequest, "Invalid Request Body", err.Error())
	}

	create := service.StoreCategory(category)
	Success(w, http.StatusOK, "Success create data", create)
}