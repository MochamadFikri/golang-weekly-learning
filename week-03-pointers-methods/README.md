# 📝 Todo CLI App - Belajar Golang Minggu ke-3

Mini project sederhana untuk latihan pemahaman konsep **Pointer** dan **Receiver Method** di Golang.

> Dibuat sebagai bagian dari weekly learning goal saya mempelajari Golang dari nol. Minggu ke-3 fokus pada pemahaman cara kerja method receiver (`value` vs `pointer`) melalui studi kasus nyata.

---

## 🎯 Fitur

- ✅ Tambah tugas
- ✅ Tandai tugas selesai
- ✅ Hapus tugas
- ✅ Tampilkan semua tugas

---

## 💡 Tujuan Pembelajaran

- Memahami perbedaan **value receiver (`T`)** dan **pointer receiver (`*T`)**
- Membangun aplikasi CLI sederhana menggunakan **struct** dan **slice**
- Modularisasi kode dengan package lokal di Golang

---

## 🧱 Struktur Proyek
week-03-todo/
├── main.go # Entry point aplikasi CLI
└── todo/
└── todo.go # Struct Todo & TodoList serta method-method receiver


---

## 🛠 Cara Menjalankan

1. Clone repo ini:

```bash
go run main.go
