package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"week02interface/interfaces"
	"week02interface/models"
	"week02interface/utils"
)

func printArea(s interfaces.Shape) {
	fmt.Println("Area:", s.Area())
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Input nilai jari-jari untuk hitung luas lingkaran
	fmt.Print("Masukkan jari-jari lingkaran: ")
	rInput, _ := reader.ReadString('\n')
	rInput = strings.TrimSpace(rInput)
	radius, err := strconv.ParseFloat(rInput, 64)
	if err != nil {
		fmt.Println("Input jari-jari tidak valid:", err)
		return
	}

	c := models.Circle{Radius: radius}
	printArea(c)

	// Input angka untuk pembagian
	fmt.Print("Masukkan angka pertama: ")
	input1, _ := reader.ReadString('\n')
	input1 = strings.TrimSpace(input1)
	a, err := strconv.ParseFloat(input1, 64)
	if err != nil {
		fmt.Println("Input pertama tidak valid:", err)
		return
	}

	fmt.Print("Masukkan angka kedua: ")
	input2, _ := reader.ReadString('\n')
	input2 = strings.TrimSpace(input2)
	b, err := strconv.ParseFloat(input2, 64)
	if err != nil {
		fmt.Println("Input kedua tidak valid:", err)
		return
	}

	result, err := utils.Divide(a, b)
	if err != nil {
		fmt.Println("Terjadi error saat membagi:", err)
	} else {
		fmt.Println("Hasil pembagian:", result)
	}
}
