package repositories

import (
	"database/sql"
	"fmt"
	"task-session-1/models"
	"time"
)

type TransactionRepository struct {
	db *sql.DB
}

func NewTransactionRepository(db *sql.DB) *TransactionRepository {
	return &TransactionRepository{db: db}
}

func (repo *TransactionRepository) CreateTransaction(items []models.CheckoutItem) (*models.Transaction, error) {
	tx, err := repo.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	totalAmount := 0
	var details []models.TransactionDetail

	for _, item := range items {
		var productPrice, stock int
		var productName string

		err := tx.QueryRow(
			"SELECT name, price, stock FROM product WHERE id = $1",
			item.ProductID,
		).Scan(&productName, &productPrice, &stock)

		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("product id %d not found", item.ProductID)
		}
		if err != nil {
			return nil, err
		}

		if stock < item.Quantity {
			return nil, fmt.Errorf("insufficient stock for product id %d", item.ProductID)
		}

		subtotal := productPrice * item.Quantity
		totalAmount += subtotal

		_, err = tx.Exec(
			"UPDATE product SET stock = stock - $1 WHERE id = $2",
			item.Quantity,
			item.ProductID,
		)
		if err != nil {
			return nil, err
		}

		details = append(details, models.TransactionDetail{
			ProductID:   item.ProductID,
			ProductName: productName,
			Quantity:    item.Quantity,
			Subtotal:    subtotal,
		})
	}

	var transactionID int
	err = tx.QueryRow(
		"INSERT INTO transactions (total_amount) VALUES ($1) RETURNING id",
		totalAmount,
	).Scan(&transactionID)

	if err != nil {
		return nil, err
	}

	for i := range details {
		details[i].TransactionID = transactionID

		_, err = tx.Exec(
			"INSERT INTO transaction_details (transaction_id, product_id, quantity, subtotal) VALUES ($1, $2, $3, $4)",
			details[i].TransactionID,
			details[i].ProductID,
			details[i].Quantity,
			details[i].Subtotal,
		)
		if err != nil {
			return nil, err
		}
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return &models.Transaction{
		ID:          transactionID,
		TotalAmount: totalAmount,
		Details:     details,
	}, nil
}

func (repo *TransactionRepository) GetReport(startDate, endDate time.Time) (*models.ReportResponse, error) {
	// Total revenue and total transactions
	var totalRevenue, totalTransaksi int
	err := repo.db.QueryRow(`
		SELECT COALESCE(SUM(total_amount), 0), COUNT(*)
		FROM transactions
		WHERE DATE(created_at) >= $1 AND DATE(created_at) <= $2
	`, startDate.Format("2006-01-02"), endDate.Format("2006-01-02")).Scan(&totalRevenue, &totalTransaksi)
	if err != nil {
		return nil, err
	}

	// Best-selling product
	var nama string
	var qtyTerjual int
	err = repo.db.QueryRow(`
		SELECT p.name, COALESCE(SUM(td.quantity), 0) as total_qty
		FROM product p
		LEFT JOIN transaction_details td ON p.id = td.product_id
		LEFT JOIN transactions t ON td.transaction_id = t.id
		WHERE DATE(t.created_at) >= $1 AND DATE(t.created_at) <= $2
		GROUP BY p.id, p.name
		ORDER BY total_qty DESC
		LIMIT 1
	`, startDate.Format("2006-01-02"), endDate.Format("2006-01-02")).Scan(&nama, &qtyTerjual)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	if err == sql.ErrNoRows {
		nama = ""
		qtyTerjual = 0
	}

	return &models.ReportResponse{
		TotalRevenue:   totalRevenue,
		TotalTransaksi: totalTransaksi,
		ProdukTerlaris: models.ProdukTerlarisResponse{
			Nama:       nama,
			QtyTerjual: qtyTerjual,
		},
	}, nil
}
