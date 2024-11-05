package main

import (
	"fmt"
)

// Fungsi rekursif untuk menghitung jumlah pertemuan
func hitungPertemuan(hari int, x int, y int, count int) int {
	// Basis rekursi: Jika hari > 365, kembalikan count
	if hari > 365 {
		return count
	}

	// Memeriksa apakah hari adalah kelipatan x tapi bukan kelipatan y
	if hari%x == 0 && hari%y != 0 {
		// Menambah 1 ke count jika hari memenuhi syarat
		count++
	}

	// Melanjutkan ke hari berikutnya dengan rekursi
	return hitungPertemuan(hari+1, x, y, count)
}

func main() {
	// Meminta input bilangan positif x dan y dari pengguna
	var x, y int

	fmt.Print("Masukkan nilai x : ")
	fmt.Scan(&x)
	fmt.Print("Masukkan nilai y : ")
	fmt.Scan(&y)

	// Memanggil fungsi rekursif dari hari ke-1 dengan jumlah pertemuan awal 0
	jumlahPertemuan := hitungPertemuan(1, x, y, 0)

	// Menampilkan jumlah pertemuan dalam setahun
	fmt.Printf("Jumlah pertemuan rahasia dalam setahun: %d\n", jumlahPertemuan)
}