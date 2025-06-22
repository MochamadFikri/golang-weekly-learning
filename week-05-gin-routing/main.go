package main

import (
	"week-05-gin-routing/books"

	"github.com/gin-gonic/gin"
)

func main() {
	// Inisialisasi router dari Gin
	r := gin.Default()

	// Registrasi semua route dari modul buku
	books.RegisterRoutes(r)

	// Jalankan server di port 8080
	r.Run(":8080")
}
