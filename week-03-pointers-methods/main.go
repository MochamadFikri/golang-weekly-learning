package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"week-03-pointers-methods/todo"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	list := todo.NewTodoList()

	for {
		fmt.Println("\nMenu:")
		fmt.Println("1. Tambah Todo")
		fmt.Println("2. Tandai Selesai")
		fmt.Println("3. Hapus Todo")
		fmt.Println("4. Lihat Semua Todo")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih: ")

		choiceStr, _ := reader.ReadString('\n')
		choice := strings.TrimSpace(choiceStr)

		switch choice {
		case "1":
			fmt.Print("Masukkan tugas: ")
			task, _ := reader.ReadString('\n')
			list.Add(strings.TrimSpace(task))
		case "2":
			fmt.Print("ID todo yang selesai: ")
			idStr, _ := reader.ReadString('\n')
			id, _ := strconv.Atoi(strings.TrimSpace(idStr))
			if err := list.Complete(id); err != nil {
				fmt.Println("Error:", err)
			}
		case "3":
			fmt.Print("ID todo yang dihapus: ")
			idStr, _ := reader.ReadString('\n')
			id, _ := strconv.Atoi(strings.TrimSpace(idStr))
			if err := list.Delete(id); err != nil {
				fmt.Println("Error:", err)
			}
		case "4":
			list.ShowAll()
		case "0":
			fmt.Println("Keluar. Sampai jumpa!")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
