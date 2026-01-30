package handlers

import (
	"encoding/json"
	"net/http"
	// "strconv"
	// "strings"
	// "task-session-1/models"
	"task-session-1/services"
)

type ProductHandler struct {
	service *services.ProductService
}

func NewProductHandler(service *services.ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

// get api /api/product
func (h *ProductHandler) HandleProduct(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.GetAll(w, r)
	}
}

func (h *ProductHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	product, err := h.service.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}
