package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int
	fmt.Print("Masukkan jumlah ikan yang dijual (x): ")
	fmt.Scan(&x) // Input jumlah ikan yang dijual (x) 
	fmt.Print("Masukkan jumlah ikan per wadah (y): ")
	fmt.Scan(&y) // Input jumlah ikan per wadah (y) 

	// Validasi input agar nilai tidak kurang dari atau sama dengan 0
	if x <= 0 || y <= 0 {
		fmt.Println("Jumlah ikan dan jumlah ikan per wadah harus lebih dari 0.")
		return // Jika input tidak valid, program dihentikan
	}

	// Input berat ikan masing-masing
	weights := make([]float64, x) // Array untuk menyimpan berat masing-masing ikan
	fmt.Printf("Masukkan berat ikan (sejumlah x): ")
	for i := 0; i < x; i++ {
		fmt.Scan(&weights[i]) // Menyimpan berat ikan ke dalam array
	}

	// Hitung jumlah wadah yang diperlukan
	numContainers := int(math.Ceil(float64(x) / float64(y))) // Jumlah wadah dihitung dengan pembulatan ke atas

	// Distribusi ikan ke dalam wadah dan hitung total berat per wadah
	containerWeights := make([]float64, numContainers) // Array untuk menyimpan total berat setiap wadah
	for i := 0; i < x; i++ {
		containerIndex := i / y // Tentukan indeks wadah berdasarkan pembagian
		containerWeights[containerIndex] += weights[i] // Tambahkan berat ikan ke wadah yang sesuai
	}

	// Menampilkan total berat setiap wadah
	fmt.Println("\nTotal berat setiap wadah: ")
	for i, berat := range containerWeights { //Mengiterasikan total berat setiap wadah
		fmt.Printf("Wadah %d: %.2f\n", i+1, berat) // Cetak total berat untuk setiap wadah
	}

	// Hitung rata-rata berat ikan di setiap wadah
	var totalWeight float64
	for _, berat := range containerWeights {
		totalWeight += berat // Hitung total berat semua wadah
	}
	averageWeight := totalWeight / float64(numContainers) // Rata-rata berat dihitung dengan membagi total berat dengan jumlah wadah

	// Menampilkan rata-rata berat ikan di setiap wadah
	fmt.Printf("\nBerat rata-rata ikan di setiap wadah: %.2f\n", averageWeight) // Cetak rata-rata berat
}
