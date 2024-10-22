package main

import (
	"fmt"
)

// Fungsi untuk menangani angka genap
func genap(n int) int {
	return n / 2 
}

// Fungsi untuk menangani angka ganjil
func ganjil(n int) int {
	return 3*n + 1 
}

// Fungsi untuk menjalankan urutan Skiena
func skiena(n int) {
	fmt.Println("Urutan Skiena : ", n) // Menampilkan angka awal
	for n != 1 { // Looping hingga n menjadi 1
		fmt.Printf("%d ", n) 
		if n%2 == 0 { 
			n = genap(n) // Memanggil fungsi genap
		} else { 
			n = ganjil(n) // Memanggil fungsi ganjil
		}
	}
	fmt.Println(1) // Menampilkan angka akhir 1
}

func main() {
	const nmax = 1000000 // Batas maksimal input n
	var n int
	// Meminta input dari pengguna
	fmt.Print("Masukkan Angka : ")
	fmt.Scan(&n) 

	if n < nmax { // Jika angka kurang dari nmax
		skiena(n) // Panggil fungsi Skiena untuk menampilkan urutan
	} else {
		fmt.Println("Input harus bilangan positif dan kurang dari 1 juta!") 
	}
}
