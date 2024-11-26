package main

import (
	"bufio" 
	"fmt"   
	"os"    
	"strconv"
	"strings" 
)

// Fungsi ini mengurutkan array secara ascending menggunakan metode insertion sort.
func insertionSort(arr []int) {
	for i := 1; i < len(arr); i++ { 
		key := arr[i]  
		j := i - 1

		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j] 
			j-- 
		}
		// Sisipkan key pada posisi yang benar
		arr[j+1] = key
	}
}

//
// Fungsi ini memeriksa apakah selisih antara elemen-elemen dalam array tetap sama.
func cekJarakTetap(arr []int) (bool, int) {
	if len(arr) < 2 { // Jika array memiliki kurang dari dua elemen, dianggap memiliki jarak tetap
		return true, 0
	}
	jarak := arr[1] - arr[0] // Menentukan jarak antara elemen pertama dan kedua
	// Periksa jarak antar elemen lainnya
	for i := 1; i < len(arr)-1; i++ {
		if arr[i+1]-arr[i] != jarak { // Jika jarak antara elemen tidak sama, return false
			return false, 0
		}
	}
	// Jika jarak antar elemen tetap sama, return true dan jaraknya
	return true, jarak
}

func main() {
	// Membuat scanner untuk membaca input dari standar input (keyboard)
	scanner := bufio.NewScanner(os.Stdin)

	// Meminta input dari pengguna
	fmt.Print("Masukkan data (diakhiri dengan bilangan negatif) : ")
	scanner.Scan() 
	input := scanner.Text() 

	// Memisahkan input berdasarkan spasi
	angkaStr := strings.Fields(input)
	var data []int // Menyiapkan slice untuk menampung data angka

	// Mengonversi string input menjadi integer dan menambahkannya ke dalam slice 'data'
	for i := 0; i < len(angkaStr); i++ {
		num, err := strconv.Atoi(angkaStr[i]) 
		if err != nil { 
			fmt.Println("Input tidak valid!") 
			return
		}
		if num < 0 { 
			break
		}
		// Menambahkan angka ke dalam slice data
		data = append(data, num)
	}

	// Mengurutkan data menggunakan insertion sort
	insertionSort(data)

	// Memeriksa apakah jarak antar elemen tetap
	isTetap, jarak := cekJarakTetap(data)

	// Menampilkan hasil urutan data
	fmt.Println("\nKeluaran : ")
	for i := 0; i < len(data); i++ {
		fmt.Printf("%d ", data[i]) 
	}
	fmt.Println() 

	// Menampilkan hasil pemeriksaan jarak antar elemen
	if isTetap {
		fmt.Printf("Data berjarak %d\n", jarak) 
	} else {
		fmt.Println("Data berjarak tidak tetap")
	}
}
