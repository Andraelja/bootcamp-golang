package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	// "strconv"
	// "strings"
)

type Category struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

var category = []Category{
	{ID: 1, Name: "Hardware", Description: "Kategori pada hardware"},
	{ID: 2, Name: "Software", Description: "Kategori pada software"},
	{ID: 3, Name: "Lainnya", Description: "Kategori pada lainnya"},
}

func main() {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "Ok",
			"message": "Api is running!",
		})
	})

	http.HandleFunc("/api/categories", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(category)
		} else if r.Method == "POST" {
			var newCategory Category
			err := json.NewDecoder(r.Body).Decode(&newCategory)
			if err != nil {
				http.Error(w, "Invalid Request!", http.StatusBadRequest)
				return
			}

			newCategory.ID = len(category) + 1
			category = append(category, newCategory)

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(newCategory)
		}
	})

	fmt.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
