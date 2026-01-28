package repositories

import (
	"database/sql"
	"errors"
	"task-session-1/models"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository{
	return &ProductRepository{db: db}
}