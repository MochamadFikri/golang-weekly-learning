package product

import (
    "bufio"
    "fmt"
    "strconv"
    "strings"
)

type Product struct {
    ID    int
    Name  string
    Stock int
}

func ShowProducts(products []Product) {
    fmt.Println("== Daftar Produk ==")
    for _, p := range products {
        fmt.Printf("ID: %d | Nama: %s | Stok: %d\n", p.ID, p.Name, p.Stock)
    }
}

func AddProduct(reader *bufio.Reader, products []Product) []Product {
    fmt.Print("Nama Produk: ")
    name, _ := reader.ReadString('\n')
    name = strings.TrimSpace(name)

    fmt.Print("Stok: ")
    stockInput, _ := reader.ReadString('\n')
    stockInput = strings.TrimSpace(stockInput)
    stock, _ := strconv.Atoi(stockInput)

    id := len(products) + 1
    newProduct := Product{ID: id, Name: name, Stock: stock}
    return append(products, newProduct)
}

func UpdateStock(reader *bufio.Reader, products []Product) []Product {
    fmt.Print("ID Produk: ")
    idInput, _ := reader.ReadString('\n')
    id, _ := strconv.Atoi(strings.TrimSpace(idInput))

    fmt.Print("Stok Baru: ")
    stockInput, _ := reader.ReadString('\n')
    stock, _ := strconv.Atoi(strings.TrimSpace(stockInput))

    for i, p := range products {
        if p.ID == id {
            products[i].Stock = stock
        }
    }
    return products
}

func DeleteProduct(reader *bufio.Reader, products []Product) []Product {
    fmt.Print("ID Produk yang Dihapus: ")
    idInput, _ := reader.ReadString('\n')
    id, _ := strconv.Atoi(strings.TrimSpace(idInput))

    for i, p := range products {
        if p.ID == id {
            return append(products[:i], products[i+1:]...)
        }
    }
    return products
}
