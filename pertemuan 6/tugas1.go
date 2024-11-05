package main

import (
	"fmt"
	"math"
)


func tampilArray(arr []int) {
	fmt.Println("Isi array: ", arr)
}


func indeksGanjil(arr []int) {
	fmt.Print("Elemen dengan indeks ganjil: ")
	for i := 1; i < len(arr); i += 2 {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()
}


func indeksGenap(arr []int) {
	fmt.Print("Elemen dengan indeks genap: ")
	for i := 0; i < len(arr); i += 2 {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()
}


func kelipatan(arr []int, x int) {
	fmt.Printf("Elemen pada indeks kelipatan %d: ", x)
	for i := x; i < len(arr); i += x {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()
}


func hapusElemen(arr []int, idx int) []int {
	return append(arr[:idx], arr[idx+1:]...)
}

func hitungRataRata(arr []int) float64 {
	total := 0
	for _, nilai := range arr {
		total += nilai
	}
	return float64(total) / float64(len(arr))
}


func standarDeviasi(arr []int, rata float64) float64 {
	rataRata := hitungRataRata(arr)
	var total float64
	for _, nilai := range arr {
		total += math.Pow(float64(nilai)-rataRata, 2)
	}
	return math.Sqrt(total / float64(len(arr)))
}


func frekuensi(arr []int, bilangan int) int {
	jumlah := 0
	for _, nilai := range arr {
		if nilai == bilangan {
			jumlah++
		}
	}
	return jumlah
}

func main() {
	var n, x, deleteIndex, bilangan int


	fmt.Print("Masukkan jumlah elemen array (N): ")
	fmt.Scan(&n)

	
	array := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan nilai untuk elemen ke-%d: ", i)
		fmt.Scan(&array[i])
	}


	for {
		fmt.Println("\nPilih Opsi:")
		fmt.Println("1. Tampilkan keseluruhan isi array")
		fmt.Println("2. Tampilkan elemen array dengan indeks ganjil")
		fmt.Println("3. Tampilkan elemen array dengan indeks genap")
		fmt.Println("4. Tampilkan elemen array dengan indeks kelipatan x")
		fmt.Println("5. Hapus elemen array pada indeks tertentu")
		fmt.Println("6. Tampilkan rata-rata dari elemen array")
		fmt.Println("7. Tampilkan standar deviasi elemen array")
		fmt.Println("8. Tampilkan frekuensi dari suatu bilangan dalam array")
		fmt.Println("9. Keluar")

		var pilihan int
		fmt.Print("Masukkan pilihan: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			tampilArray(array)

		case 2:
			indeksGanjil(array)

		case 3:
			indeksGenap(array)

		case 4:
			fmt.Print("Masukkan nilai x untuk kelipatan indeks: ")
			fmt.Scan(&x)
			kelipatan(array, x)

		case 5:
			fmt.Print("Masukkan indeks yang ingin dihapus: ")
			fmt.Scan(&deleteIndex)
			if deleteIndex >= 0 && deleteIndex < len(array) {
				array = hapusElemen(array, deleteIndex)
				tampilArray(array)
			} else {
				fmt.Println("Indeks tidak valid!")
			}

		case 6:
			rataRata := hitungRataRata(array)
			fmt.Printf("Rata-rata elemen array: %.2f\n", rataRata)

		case 7:
			rataRata := hitungRataRata(array)
			standarDeviasi := standarDeviasi(array, rataRata)
			fmt.Printf("Standar deviasi elemen array: %.2f\n", standarDeviasi)

		case 8:
			fmt.Print("Masukkan nilai untuk mencari frekuensi: ")
			fmt.Scan(&bilangan)
			frekuensi := frekuensi(array, bilangan)
			fmt.Printf("Frekuensi %d dalam array: %d kali\n", bilangan, frekuensi)

		case 9:
			fmt.Println("Keluar dari program.")
			return

		default:
			fmt.Println("Pilihan tidak valid, silahkan coba lagi!")
		}
	}
}