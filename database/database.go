// Package database berisi fungsi untuk inisialisasi koneksi database.
// Ini memisahkan logika koneksi database dari kode utama aplikasi.
package database

// Import library yang diperlukan.
// database/sql adalah package standar Go untuk berinteraksi dengan database.
// log digunakan untuk mencetak pesan ke konsol.
// _ "github.com/lib/pq" adalah driver PostgreSQL, underscore (_) berarti hanya untuk side effect (registrasi driver).
import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

// InitDB adalah fungsi untuk menginisialisasi koneksi ke database PostgreSQL.
// Menerima connectionString sebagai parameter, yang berisi detail koneksi seperti host, port, user, password, dll.
// Mengembalikan pointer ke sql.DB dan error jika ada.
func InitDB(connectionString string) (*sql.DB, error) {
	// Membuka koneksi ke database PostgreSQL menggunakan driver "postgres".
	// sql.Open tidak langsung membuat koneksi, hanya mempersiapkan.
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	// Menguji koneksi dengan Ping untuk memastikan database dapat diakses.
	// Ini akan membuat koneksi fisik jika belum ada.
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	// Mengatur pengaturan connection pool (opsional tapi direkomendasikan).
	// SetMaxOpenConns membatasi jumlah koneksi maksimal yang terbuka.
	db.SetMaxOpenConns(25)
	// SetMaxIdleConns membatasi jumlah koneksi idle (tidak digunakan) yang disimpan.
	db.SetMaxIdleConns(5)

	// Mencetak pesan sukses ke log.
	log.Println("Database connected successfully")
	return db, nil
}
