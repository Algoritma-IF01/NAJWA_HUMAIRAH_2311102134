package main

import (
	"fmt"
)

// Fungsi untuk memeriksa apakah suatu bilangan adalah perfect number
func isPerfectNumber(x int) bool {
	sum := 0

	// Cari faktor-faktor dari x (kecuali x sendiri)
	for i := 1; i < x; i++ {
		if x%i == 0 {
			sum += i
		}
	}
	// Periksa apakah penjumlahan faktor sama dengan x
	return sum == x
}

// Fungsi untuk menampilkan perfect number dalam rentang a hingga b
func perfectNumbersInRange(a int, b int) {
	fmt.Printf("Perfect number dari %d hingga %d:\n", a, b)

	// Iterasikan melalui bilangan dari a hingga b
	found := false
	for i := a; i <= b; i++ {
		if isPerfectNumber(i) {
			fmt.Println(i)
			found = true
		}
	}

	// Jika perfect number tidak ditemukan
	if !found {
		fmt.Println("Tidak ada perfect number dalam rentang ini")
	}
}

func main() {
	// Meminta input bilanga bulat a dan b dari pengguna
	var a, b int

	fmt.Print("Masukkan nilai a : ")
	fmt.Scan(&a)
	fmt.Print("Masukkan nilai b : ")
	fmt.Scan(&b)

	// Validasi input a harus lebih kecil atau sama dengan b
	if a > b {
		fmt.Println("Nilai a harus lebih kecil atau sama dengan b")
		return
	}
	// Memanggil fungsi untuk menampilkan perfect number dalam rentang a hingga b
	perfectNumbersInRange(a, b)
}