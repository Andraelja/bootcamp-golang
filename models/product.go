// Package models berisi definisi struct data untuk aplikasi.
// Struct ini merepresentasikan tabel di database dan digunakan untuk serialisasi JSON.
// Ini memisahkan logika data dari logika bisnis.
package models

// Product adalah struct yang merepresentasikan data produk.
// Struct ini memiliki field ID, Name, Price, Stock, CategoryID, dan relasi ke Category.
// Tag json digunakan untuk serialisasi/deserialisasi JSON saat API response/request.
type Product struct {
	// ID adalah identifier unik untuk produk, biasanya auto-increment dari database.
	ID int `json:"id"`
	// Name adalah nama produk, seperti "Laptop" atau "Baju".
	Name string `json:"name"`
	// Price adalah harga produk dalam rupiah.
	Price int `json:"price"`
	// Stock adalah jumlah stok produk yang tersedia.
	Stock int `json:"stock"`
	// CategoryID adalah ID kategori yang terkait dengan produk ini.
	CategoryID int `json:"category_id"`

	// Category adalah relasi ke tabel category, opsional (tidak selalu di-include dalam JSON).
	// Tag omitempty berarti field ini tidak akan muncul di JSON jika nil.
	Category *Category `json:"category,omitempty"`
}
