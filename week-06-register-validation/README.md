# 🧾 Register API - Belajar Golang Minggu ke-6

Mini project sederhana untuk latihan memahami konsep **JSON Binding** dan **Validasi Manual** di Gin Golang.

> Dibuat sebagai bagian dari weekly learning goal saya mempelajari Golang dari nol. Minggu ke-6 fokus pada pembuatan endpoint `POST /register` dengan **binding JSON** ke struct dan **validasi input manual** tanpa menggunakan library tambahan.

---

## 🎯 Fitur

- ✅ Endpoint register user (`POST /register`)
- ✅ Binding data JSON ke struct
- ✅ Validasi manual untuk nama, email, dan password

---

## 💡 Tujuan Pembelajaran

- Memahami penggunaan `ShouldBindJSON()` di Gin
- Menerapkan validasi manual tanpa package validator eksternal
- Menangani input dan error dengan respon API yang ramah pengguna
- Menyusun struktur proyek web sederhana secara modular

---

## 🧱 Struktur Proyek
```

week-06-register-validation/
├── controllers/
│   └── user\_controller.go   # Handler untuk endpoint register
├── models/
│   └── user.go              # Struct input user
├── routes/
│   └── user\_routes.go       # Routing endpoint register
├── go.mod                   # Modul dan dependency Go
├── go.sum                   # Checksum dependency
└── main.go                  # Entry point aplikasi

````

---

## 🛠 Cara Menjalankan

1. Clone repositori atau masuk ke folder proyek:

```bash
cd week-06-register-validation
````

2. Inisialisasi module dan install dependency:

```bash
go mod tidy
```

3. Jalankan aplikasi:

```bash
go run main.go
```

4. Akses endpoint via Postman atau cURL:

* `POST http://localhost:8080/register`

---

## 📌 Contoh JSON Input

**Sukses:**

```json
{
  "nama": "Fikri",
  "email": "fikri@mail.com",
  "password": "rahasia123"
}
```

**Gagal (Validasi):**

```json
{
  "nama": "",
  "email": "salahformat",
  "password": "123"
}
```

---

## 🔁 Contoh Respon

**Berhasil:**

```json
{
  "pesan": "Registrasi berhasil",
  "user": {
    "nama": "Fikri",
    "email": "fikri@mail.com",
    "password": "rahasia123"
  }
}
```

**Gagal (Error Validasi):**

```json
{
  "error": "Nama wajib diisi"
}
```