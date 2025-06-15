# 📦 Slug Formatter - Belajar Golang Minggu ke-4

Di minggu ke-4 ini aku mulai belajar tentang hal penting di Golang yaitu **Go Modules** dan bagaimana cara mengelola library eksternal dengan `go mod`.

## 🚀 Apa itu Go Modules?

Sejak Go versi 1.11, dependency management di Go dilakukan melalui sistem `go mod`. Dengan sistem ini, kita bisa:

- Mendefinisikan nama modul / project
- Mengimpor library eksternal dengan mudah
- Menyimpan informasi versi library agar project tetap stabil dan konsisten

---

## 🔧 Tools dan Perintah yang Dipakai

| Perintah              | Fungsi                                                  |
|-----------------------|----------------------------------------------------------|
| `go mod init`         | Membuat module baru dan file `go.mod`                   |
| `go get`              | Menambahkan library eksternal                           |
| `go run`              | Menjalankan program Go                                  |
| `go mod tidy`         | Membersihkan dependensi yang tidak terpakai             |

---

## 🛠 Mini Project: CLI Slug Generator

Untuk latihan, aku membuat **Command Line App** sederhana untuk mengubah teks biasa menjadi *slug format* yang cocok untuk URL.  
Contoh: `"Halo Dunia Golang"` → `halo-dunia-golang`

### ✨ Fitur:
- Input teks via terminal
- Output slug hasil konversi
- Menggunakan library `github.com/gosimple/slug`

---

## 🧱 Struktur Proyek

```
week-04-go-modules/
├── go.mod // definisi modul & dependensi
├── go.sum // verifikasi versi library
└── main.go // kode utama program CLI
```


---

## 💻 Cara Menjalankan

1. Clone repo ini
2. Masuk ke folder `week-04-go-modules`
3. Jalankan perintah:

```bash
go mod tidy
go run main.go "Tulis teks kamu di sini"
```