package main

import "fmt"

func main() {
	var A, B string
	var skorA, skorB int
	var pemenang []string
	pertandingan := 1
	pertandinganSelesai := false

	
	fmt.Print("Klub A: ")
	fmt.Scanln(&A)
	fmt.Print("Klub B: ")
	fmt.Scanln(&B)


	for {
		fmt.Printf("Pertandingan %d: ", pertandingan)
		fmt.Scan(&skorA, &skorB)


		if skorA < 0 || skorB < 0 {
			fmt.Println("Pertandingan selesai")
			break
		}

	
		if skorA > skorB {
			pemenang = append(pemenang, A)
			fmt.Printf("Hasil %d %s\n", pertandingan, A)
		} else if skorB > skorA {
			pemenang = append(pemenang, B)
			fmt.Printf("Hasil %d %s\n", pertandingan, B)
		} else {
			pemenang = append(pemenang, "Draw")
			fmt.Printf("Hasil %d Draw\n", pertandingan)
		}

		pertandingan++
	}

	
	fmt.Println("\nDaftar Klub Pemenang:")
	for i, hasil := range pemenang {
		fmt.Printf("Pertandingan %d: %s\n", i+1, hasil)
	}

	// Menampilkan pesan akhir jika pertandingan dihentikan karena skor negatif
	if pertandinganSelesai {
		fmt.Println("Pertandingan selesai.")
	}
}