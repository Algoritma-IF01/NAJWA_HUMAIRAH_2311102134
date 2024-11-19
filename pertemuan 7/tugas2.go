package main

import (
	"fmt"
)

// Definisi tipe array dengan kapasitas maksimum 100 elemen
type arrBalita [100]float64

// Fungsi untuk menghitung nilai minimum dan maksimum dalam array
func hitungMinMax(arrBerat arrBalita, bMin, bMax *float64, n int) {
	// Inisialisasi nilai awal bMin dan bMax dengan elemen pertama array
	*bMin = arrBerat[0]
	*bMax = arrBerat[0]

	// Iterasi untuk mencari nilai minimum dan maksimum dari elemen array
	for i := 0; i < n; i++ {
		if arrBerat[i] < *bMin {
			*bMin = arrBerat[i] // Jika elemen lebih kecil dari bMin, perbarui bMin
		}
		if arrBerat[i] > *bMax {
			*bMax = arrBerat[i] // Jika elemen lebih besar dari bMax, perbarui bMax
		}
	}
}

// Fungsi untuk menghitung rata-rata berat balita dalam array
func rerata(arrBerat arrBalita, n int) float64 {
	total := 0.0 // Variabel untuk menyimpan jumlah total berat balita

	// Iterasi untuk menjumlahkan semua elemen array
	for i := 0; i < n; i++ {
		total += arrBerat[i] // Tambahkan berat balita ke dalam total
	}

	// Mengembalikan rata-rata (total berat dibagi jumlah elemen)
	return total / float64(n)
}

func main() {
	var n int              // Variabel untuk jumlah data berat balita
	var arrBerat arrBalita // Array untuk menyimpan berat balita
	var bMin, bMax float64 // Variabel untuk menyimpan nilai minimum dan maksimum

	// Meminta input jumlah data berat balita
	fmt.Print("Masukkan banyak data berat balita: ")
	fmt.Scan(&n)

	// Validasi jumlah data (harus antara 1 dan 100)
	if n <= 0 || n > 100 {
		fmt.Println("Jumlah data harus antara 1 hingga 100.")
		return // Jika input tidak valid, hentikan program
	}

	// Input data berat balita
	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan berat balita ke-%d: ", i+1)
		fmt.Scan(&arrBerat[i]) // Menyimpan input ke dalam array
	}

	// Memanggil fungsi untuk menghitung nilai minimum dan maksimum
	hitungMinMax(arrBerat, &bMin, &bMax, n)

	// Memanggil fungsi untuk menghitung rata-rata berat balita
	rataRata := rerata(arrBerat, n)

	// Menampilkan hasil
	fmt.Printf("\nBerat balita minimum: %.2f kg\n", bMin)  // Menampilkan berat minimum balita
	fmt.Printf("Berat balita maksimum: %.2f kg\n", bMax)   // Menampilkan berat maksimum balita
	fmt.Printf("Rerata berat balita: %.2f kg\n", rataRata) // Menampilkan rata-rata berat balita
}