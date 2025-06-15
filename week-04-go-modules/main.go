package main

import (
	"fmt"
	"os"

	"github.com/gosimple/slug"
)

func main() {
	// Cek apakah pengguna sudah memasukkan input dari terminal
	if len(os.Args) < 2 {
		fmt.Println("Cara pakai:")
		fmt.Println("go run main.go \"Teks yang ingin diubah menjadi slug\"")
		return
	}

	// Ambil input dari argument pertama setelah nama file
	teksAsli := os.Args[1]

	// Ubah teks menjadi slug
	slugHasil := slug.Make(teksAsli)

	// Tampilkan hasil slug
	fmt.Println("Slug yang dihasilkan:", slugHasil)
}
