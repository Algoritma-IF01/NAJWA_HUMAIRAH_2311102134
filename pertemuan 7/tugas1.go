package main

import (
	"fmt"
	"math"
)

func main() {
	var totalIkan, ikanPerWadah int
	fmt.Print("Masukkan jumlah ikan yang dijual (x): ")
	fmt.Scan(&totalIkan) // Input jumlah ikan yang dijual (x) 
	fmt.Print("Masukkan jumlah ikan per wadah (y): ")
	fmt.Scan(&ikanPerWadah) // Input jumlah ikan per wadah (y) 

	// Validasi input agar nilai tidak kurang dari atau sama dengan 0
	if totalIkan <= 0 || ikanPerWadah <= 0 {
		fmt.Println("Jumlah ikan dan jumlah ikan per wadah harus lebih dari 0.")
		return // Jika input tidak valid, program dihentikan
	}

	// Input berat ikan masing-masing
	beratIkan := make([]float64, totalIkan) // Array untuk menyimpan berat masing-masing ikan
	fmt.Printf("Masukkan berat ikan (sejumlah x): ")
	for i := 0; i < totalIkan; i++ {
		fmt.Scan(&beratIkan[i]) // Menyimpan berat ikan ke dalam array
	}

	// Hitung jumlah wadah yang diperlukan
	jumlahWadah := int(math.Ceil(float64(totalIkan) / float64(ikanPerWadah))) // Jumlah wadah dihitung dengan pembulatan ke atas

	// Distribusi ikan ke dalam wadah dan hitung total berat per wadah
	totalBeratwadah := make([]float64, jumlahWadah) // Array untuk menyimpan total berat setiap wadah
	for i := 0; i < totalIkan; i++ {
		indekswadah := i / ikanPerWadah // Tentukan indeks wadah berdasarkan pembagian
		totalBeratwadah[indekswadah] += beratIkan[i] // Tambahkan berat ikan ke wadah yang sesuai
	}

	// Menampilkan total berat setiap wadah
	fmt.Println("\nTotal berat setiap wadah: ")
	for i, berat := range totalBeratwadah { //Mengiterasikan total berat setiap wadah
		fmt.Printf("Wadah %d: %.2f\n", i+1, berat) // Cetak total berat untuk setiap wadah
	}

	// Hitung rata-rata berat ikan di setiap wadah
	var totalBeratSemuaWadah float64
	for _, berat := range totalBeratwadah {
		totalBeratSemuaWadah += berat // Hitung total berat semua wadah
	}
	rataRataBerat := totalBeratSemuaWadah / float64(jumlahWadah) // Rata-rata berat dihitung dengan membagi total berat dengan jumlah wadah

	// Menampilkan rata-rata berat ikan di setiap wadah
	fmt.Printf("\nBerat rata-rata ikan di setiap wadah: %.2f\n", rataRataBerat) // Cetak rata-rata berat
}
