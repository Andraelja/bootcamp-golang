package repositories

import (
	"database/sql"
	// "errors"
	"task-session-1/models"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (repo *ProductRepository) GetAll() ([]models.Product, error) {
	query := `
			SELECT 
			p.id, 
			p.name, 
			p.price, 
			p.stock, 
			p.category_id,
			c.id,
			c.name
			FROM product p JOIN category c ON c.id = p.category_id`
	rows, err := repo.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	product := make([]models.Product, 0)
	for rows.Next() {
		var p models.Product
		err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.Stock, &p.CategoryID)
		if err != nil {
			return nil, err
		}
		product = append(product, p)
	}
	return product, nil
}
