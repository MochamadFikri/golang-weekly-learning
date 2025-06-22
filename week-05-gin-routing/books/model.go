package books

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"judul"`
	Author string `json:"penulis"`
}

var BookList = []Book{
	{ID: "1", Title: "Filosofi Teras", Author: "Henry Manampiring"},
	{ID: "2", Title: "Negeri 5 Menara", Author: "Ahmad Fuadi"},
	{ID: "3", Title: "Laskar Pelangi", Author: "Andrea Hirata"},
	{ID: "4", Title: "Bumi", Author: "Tere Liye"},
	{ID: "5", Title: "Pulang", Author: "Tere Liye"},
	{ID: "6", Title: "Dilan: Dia adalah Dilanku Tahun 1990", Author: "Pidi Baiq"},
	{ID: "7", Title: "Perahu Kertas", Author: "Dewi Lestari"},
	{ID: "8", Title: "Sang Pemimpi", Author: "Andrea Hirata"},
	{ID: "9", Title: "9 Summers 10 Autumns", Author: "Iwan Setyawan"},
	{ID: "10", Title: "Kambing Jantan", Author: "Raditya Dika"},
}
