package main

import (
	"fmt"
)

// Fungsi untuk menghitung akar dari 2 berdasarkan rumus di gambar
func sqrt2(K int) float64 {
	result := 1.0 // Inisialisasi hasil dengan nilai 1.0 karena kita mengalikan suku-suku.
	for k := 0; k <= K; k++ {
		// Menghitung suku berdasarkan rumus f(K) yang diberikan pada soal.
		term := (float64(4*k+2) * float64(4*k+2)) / (float64(4*k+1) * float64(4*k+3))
		result *= term // Mengalikan hasil dengan suku yang baru dihitung.
	}
	return result // Mengembalikan hasil perkalian dari semua suku.
}

func main() {
	var K int // Deklarasi variabel K untuk menerima input dari pengguna.

	// Loop ini digunakan untuk menjalankan perhitungan sebanyak 3 kali.
	for i := 1; i <= 3; i++ {
		// Menampilkan ke pengguna untuk memasukkan nilai K.
		fmt.Print("Nilai K = ")
		fmt.Scan(&K) // Menerima input dari pengguna dan menyimpan dalam variabel K.

		// Menghitung akar 2 dengan K iterasi
		hasil := sqrt2(K) // Memanggil fungsi sqrt2 untuk menghitung akar 2 dengan nilai K.

		// Menampilkan hasil dengan presisi 10 angka di belakang koma
		fmt.Printf("Nilai akar 2 = %.10f\n\n", hasil)
	}
}
