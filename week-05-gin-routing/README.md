# 📚 Book API - Belajar Golang Minggu ke-5

Mini project sederhana untuk latihan memahami konsep **Routing dengan Gin** di Golang.

> Dibuat sebagai bagian dari weekly learning goal saya mempelajari Golang dari nol. Minggu ke-5 fokus pada pemahaman cara kerja **routing**, **parameter**, dan **handler modular** menggunakan framework web **Gin**.

---

## 🎯 Fitur

- ✅ Lihat semua buku
- ✅ Lihat detail buku berdasarkan ID
- ✅ Tambah buku baru
- ✅ Perbarui buku berdasarkan ID
- ✅ Hapus buku berdasarkan ID

---

## 💡 Tujuan Pembelajaran

- Memahami penggunaan **Gin** sebagai web framework di Golang
- Mengimplementasikan **GET, POST, PUT, DELETE** route
- Mengelola **parameter dinamis (`/buku/:id`)** dan **body JSON**
- Menerapkan struktur folder yang rapi dan modular di proyek web Go

---

## 🧱 Struktur Proyek
```
week-05-gin-routing/
├── books/
│   ├── handler.go   # Handler untuk masing-masing endpoint buku
│   ├── model.go     # Struktur data Book & data dummy
│   └── router.go    # Registrasi route grup /buku
├── go.mod           # Modul dan dependency Go
├── go.sum           # Checksum dependency
└── main.go          # Entry point aplikasi

````

---

## 🛠 Cara Menjalankan

1. Clone repositori atau buat folder project:

```bash
git clone <url-repo>
cd week-05-gin-routing
````

2. Inisialisasi module dan install dependencies:

```bash
go mod tidy
```

3. Jalankan aplikasi:

```bash
go run main.go
```

4. Akses endpoint via Postman, browser, atau cURL:

* `GET     http://localhost:8080/buku`
* `GET     http://localhost:8080/buku/:id`
* `POST    http://localhost:8080/buku`
* `PUT     http://localhost:8080/buku/:id`
* `DELETE  http://localhost:8080/buku/:id`

---

## 📌 Contoh JSON Tambah Buku

```json
{
  "id": "11",
  "judul": "Golang Untuk Pemula",
  "penulis": "Fikri Fadila"
}
```