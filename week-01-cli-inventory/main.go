package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"week01cli/product"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var products []product.Product

	for {
		fmt.Println("\n== Menu ==")
		fmt.Println("1. Lihat Produk")
		fmt.Println("2. Tambah Produk")
		fmt.Println("3. Update Stok")
		fmt.Println("4. Hapus Produk")
		fmt.Println("5. Keluar")
		fmt.Print("Pilih: ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			product.ShowProducts(products)
		case "2":
			products = product.AddProduct(reader, products)
		case "3":
			products = product.UpdateStock(reader, products)
		case "4":
			products = product.DeleteProduct(reader, products)
		case "5":
			os.Exit(0)
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
