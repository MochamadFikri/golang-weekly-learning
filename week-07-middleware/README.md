# 🧾 Middleware Logging & Token Auth - Belajar Golang Minggu ke-7

Mini project sederhana untuk latihan memahami konsep **Middleware** di Gin Golang, termasuk pembuatan middleware custom untuk **Logging** dan **Token Authentication**.

> Dibuat sebagai bagian dari weekly learning goal saya mempelajari Golang dari nol. Minggu ke-7 fokus pada penerapan middleware di Gin menggunakan custom function untuk logging request dan validasi token sederhana.

---

## 🎯 Fitur

- ✅ Middleware logging setiap request masuk
- ✅ Middleware token authentication (`Authorization: Bearer xxx`)
- ✅ Endpoint `/user/profile` hanya bisa diakses jika token valid
- ✅ Struktur modular dan mudah dikembangkan

---

## 💡 Tujuan Pembelajaran

- Memahami konsep middleware di framework Gin
- Membuat dan menggunakan middleware global & per route
- Membuat validasi token tanpa library eksternal
- Menyusun struktur proyek backend Golang secara rapi

---

## 🧱 Struktur Proyek

```

week-07-middleware/
├── handler/
│   └── user.go              # Handler endpoint user
├── middleware/
│   ├── auth.go              # Middleware token auth
│   └── logger.go            # Middleware logger request
├── go.mod                   # Modul dan dependency Go
└── main.go                  # Entry point aplikasi

````

---

## 🛠 Cara Menjalankan

1. Masuk ke folder proyek:

```bash
cd golang-weekly-learning/week-07-middleware
````

2. Install dependency (jika belum):

```bash
go mod tidy
```

3. Jalankan aplikasi:

```bash
go run main.go
```

---

## 🔗 Endpoint

* `GET /` — Endpoint publik (tanpa middleware)
* `GET /user/profile` — Endpoint privat (butuh token)

---

## 🔐 Token

Gunakan token header sebagai berikut:

```
Authorization: Bearer secrettoken123
```

---

## 🧪 Contoh cURL

**Akses tanpa token (Gagal):**

```bash
curl http://localhost:8080/user/profile
```

**Akses dengan token (Berhasil):**

```bash
curl -H "Authorization: Bearer secrettoken123" http://localhost:8080/user/profile
```

---

## 🔁 Contoh Respon

**Berhasil:**

```json
{
  "user": "Fikri",
  "role": "Admin"
}
```

**Gagal (Unauthorized):**

```json
{
  "error": "Unauthorized"
}
```

---

## 🧠 Catatan Tambahan

* Middleware `logger.go` mencatat status dan waktu proses setiap request.
* Middleware `auth.go` mengecek header token secara sederhana (bisa dikembangkan ke JWT).
* Cocok untuk dijadikan dasar membangun API production-ready secara bertahap.