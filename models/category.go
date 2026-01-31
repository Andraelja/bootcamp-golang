// Package models berisi definisi struct data untuk aplikasi.
// Struct ini merepresentasikan tabel di database dan digunakan untuk serialisasi JSON.
// Ini memisahkan logika data dari logika bisnis.
package models

// Category adalah struct yang merepresentasikan data kategori.
// Struct ini memiliki field ID, Name, dan Description.
// Tag json digunakan untuk serialisasi/deserialisasi JSON saat API response/request.
type Category struct {
	// ID adalah identifier unik untuk kategori, biasanya auto-increment dari database.
	ID int `json:"id"`
	// Name adalah nama kategori, seperti "Elektronik" atau "Pakaian".
	Name string `json:"name"`
	// Description adalah deskripsi kategori, menjelaskan apa yang ada di kategori tersebut.
	Description string `json:"description"`
}
