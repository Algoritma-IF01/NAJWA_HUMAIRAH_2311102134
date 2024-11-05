package main

import (
	"fmt"
)

// Fungsi untuk menghitung biaya sewa
func hitungBiaya(durasiJam int, durasiMenit int, isMember bool) int {
	// Variabel tarif per jam
	tarifMember := 3500
	tarifNonMember := 5000
	
	// Menghitung total jam berdasarkan input durasi
	totalJam := durasiJam
	if durasiMenit >= 10 {
		totalJam += 1
	} else if durasiJam == 0 && durasiMenit > 0 {
		totalJam = 1
	}
	
	// Menentukan tarif berdasarkan status member
	tarif := 0
	if isMember {
		tarif = tarifMember
	} else {
		tarif = tarifNonMember
	}

	// Menghitung biaya total sebelum diskon
	biaya := totalJam * tarif

	// Berikan diskon 10% jika durasi lebih dari 3 jam
	if totalJam > 3 {
		biaya = biaya * 90 / 100
	}

	return biaya
}

func main() {
	// Meminta input durasi jam, durasi menit, dan status membership dari pengguna
	var durasiJam, durasiMenit int
	var isMember bool

	fmt.Print("Masukkan durasi jam : ")
	fmt.Scan(&durasiJam)
	fmt.Print("Masukkan durasi menit : ")
	fmt.Scan(&durasiMenit)
	fmt.Print("Apakah Anda member (true/false) : ")
	fmt.Scan(&isMember)

	// Menghitung biaya sewa
	biaya := hitungBiaya(durasiJam, durasiMenit, isMember)

	// Menampilkan biaya sewa
	fmt.Printf("Biaya sewa adalah: Rp %d\n", biaya)
}