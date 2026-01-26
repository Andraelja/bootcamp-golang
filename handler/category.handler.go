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
	// validasi jika body salah
	if err := helper.DecodeJSON(r, &category); err != nil {
		Error(w, http.StatusBadRequest, "Invalid Request Body", err.Error())
	}

	create := service.StoreCategory(category)
	Success(w, http.StatusOK, "Success create data", create)
}

func UpdateCategory(w http.ResponseWriter, r *http.Request) {
	// parsing dulu string id jadi int
	id, err := helper.ParseID(r, "/api/categories/")
	if err != nil {
		Error(w, http.StatusBadRequest, "Invalid ID", err.Error())
		return
	}

	// ambil data category
	var category entity.Category
	// validasi jika body salah
	if err := helper.DecodeJSON(r, &category); err != nil {
		Error(w, http.StatusBadRequest, "Invalid request body", err.Error())
	}

	update, err := service.UpdateCategory(id, category)
	if err != nil {
		Error(w, http.StatusNotFound, "Product Not Found", nil)
		return
	}

	Success(w, http.StatusOK, "Success Update data", update)
}