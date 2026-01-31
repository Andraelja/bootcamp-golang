// Package repositories berisi kode untuk mengakses database.
// Repository layer mengabstraksi operasi database, sehingga kode lain tidak perlu tahu detail SQL.
// Ini memudahkan perubahan database atau testing dengan mock.
package repositories

// Import library yang diperlukan.
// database/sql digunakan untuk koneksi dan query database.
// models digunakan untuk struct data seperti Product.
import (
	"database/sql"
	// "errors" // Tidak digunakan, dikomentari.
	"task-session-1/models"
)

// ProductRepository adalah struct yang menyimpan koneksi database untuk operasi produk.
// db adalah pointer ke koneksi database PostgreSQL.
type ProductRepository struct {
	db *sql.DB
}

// NewProductRepository adalah konstruktor untuk membuat instance ProductRepository.
// Menerima koneksi database sebagai parameter dan mengembalikan pointer ke ProductRepository.
func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

// GetAll mengambil semua data produk dari database dengan JOIN ke tabel category.
// Query SQL mengambil data produk dan nama kategori terkait.
// Mengembalikan slice dari Product dan error jika ada.
func (repo *ProductRepository) GetAll() ([]models.Product, error) {
	// Query SQL untuk mengambil semua produk dengan JOIN ke category.
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
	// Menjalankan query dan mendapatkan rows.
	rows, err := repo.db.Query(query)
	if err != nil {
		return nil, err
	}
	// Pastikan rows ditutup setelah selesai.
	defer rows.Close()

	// Membuat slice kosong untuk menyimpan hasil.
	product := make([]models.Product, 0)
	// Iterasi setiap row hasil query.
	for rows.Next() {
		var p models.Product
		// Scan data dari row ke struct Product.
		// Perhatian: Scan hanya mengambil field produk, tidak termasuk category name (ada kesalahan di query asli).
		err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.Stock, &p.CategoryID)
		if err != nil {
			return nil, err
		}
		// Tambahkan produk ke slice.
		product = append(product, p)
	}
	return product, nil
}

// Create menyisipkan produk baru ke database.
// Menggunakan INSERT dengan RETURNING id untuk mendapatkan ID yang dihasilkan.
// Mengembalikan error jika penyisipan gagal.
func (repo *ProductRepository) Create(product *models.Product) error {
	// Query INSERT untuk menyisipkan produk baru.
	query := "INSERT INTO product (name, price, stock, category_id) VALUES ($1, $2, $3, $4) RETURNING id"
	// Menjalankan query dan scan ID yang dihasilkan ke product.ID.
	err := repo.db.QueryRow(query, product.Name, product.Price, product.Stock, product.CategoryID).Scan(&product.ID)
	return err
}

// GetByID mengambil satu produk berdasarkan ID.
// Jika produk tidak ditemukan, mengembalikan nil tanpa error.
// Mengembalikan pointer ke Product dan error jika ada.
func (repo *ProductRepository) GetByID(id int) (*models.Product, error) {
	// Query SELECT untuk mengambil produk berdasarkan ID.
	// Perhatian: Ada kesalahan sintaks di query asli (stock FROM product p WHERE id=$1).
	query := `
			SELECT
			p.id,
			p.name,
			p.price,
			p.category_id
			stock FROM product p WHERE id=$1
			`
	var p models.Product
	// Menjalankan query dan scan hasil ke struct Product.
	// Perhatian: Scan tidak sesuai dengan field yang dipilih (kurang stock, tambah category_id).
	err := repo.db.QueryRow(query, id).Scan(&p.ID, &p.Name, &p.Price, &p.Stock, p.CategoryID)

	// Jika tidak ada row ditemukan, kembalikan nil.
	if err == sql.ErrNoRows {
		return nil, nil
	}

	// Jika ada error lain, kembalikan error.
	if err != nil {
		return nil, err
	}

	// Kembalikan pointer ke produk.
	return &p, nil
}
