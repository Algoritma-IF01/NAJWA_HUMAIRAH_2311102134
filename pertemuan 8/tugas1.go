package main

import (
	"fmt"
)

// Fungsi untuk melakukan selection sort
func selectionSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		// Menemukan elemen terkecil di sisa array
		minIndex := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		// Menukar elemen yang ditemukan dengan elemen yang ada di posisi i
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
}

func main() {
	var n, m int
	// Membaca jumlah daerah
	fmt.Print("Masukkan jumlah daerah: ")
	fmt.Scan(&n)

	// Loop untuk membaca nomor rumah setiap daerah
	for i := 0; i < n; i++ {
		fmt.Print("Masukkan jumlah rumah di daerah ", i+1, ": ")
		fmt.Scan(&m)

		// Membaca nomor rumah dan menyimpannya dalam slice
		rumah := make([]int, m)
		fmt.Printf("Masukkan nomor rumah: ")
		for j := 0; j < m; j++ {
			fmt.Scan(&rumah[j])
		}

		// Mengurutkan nomor rumah menggunakan selection sort
		selectionSort(rumah)

		// Menampilkan hasil pengurutan
		fmt.Printf("Nomor rumah yang telah diurutkan: ")
		for _, nomor := range rumah {
			fmt.Print(nomor, " ")
		}
		fmt.Println()
	}
}
