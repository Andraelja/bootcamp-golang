package router

import (
	"net/http"
	"task-session-1/handler"
)

func RegisterRoutes() {
	http.HandleFunc("/health", handler.HealtHeandler)

	http.HandleFunc("/api/categories", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handler.GetAllCategory(w, r)
		case http.MethodPost:
			handler.StoreCategory(w, r)
		}
	})

	http.HandleFunc("/api/categories/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handler.GetCategoryById(w, r)
		case http.MethodPut:
			handler.UpdateCategory(w, r)
		case http.MethodDelete:
			handler.DeleteCategory(w, r)
		}
	})
}
