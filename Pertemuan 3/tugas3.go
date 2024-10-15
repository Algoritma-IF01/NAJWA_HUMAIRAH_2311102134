package main

import (
	"fmt"
)

// Fungsi untuk menghitung biaya pengiriman berdasarkan berat parsel
func hitungBiaya(berat int) (int, int, int) {
	// Menghitung berat dalam kg dan sisa gram
	beratKg := berat / 1000
	sisaGram := berat % 1000

	// Biaya dasar dihitung dengan mengalikan berat dalam kg dengan 10.000
	biayaDasar := beratKg * 10000
	biayaTambahan := 0

	// Jika berat total lebih dari 10 kg, sisa gram diabaikan (biaya tambahan = 0)
	if beratKg >= 10 {
		biayaTambahan = 0
	} else if sisaGram >= 500 {
		// Jika sisa berat >= 500 gram, biaya tambahan 5 per gram
		biayaTambahan = sisaGram * 5
	} else {
		// Jika sisa berat < 500 gram, biaya tambahan 15 per gram
		biayaTambahan = sisaGram * 15
	}

	// Total biaya adalah penjumlahan dari biaya dasar dan biaya tambahan
	totalBiaya := biayaDasar + biayaTambahan
	return biayaDasar, biayaTambahan, totalBiaya
}

func main() {
	for {
		var berat int // Deklarasi variabel untuk menyimpan berat parsel dari input

		// Menerima input berat parsel dari pengguna dalam satuan gram
		fmt.Print("Masukkan berat parsel (gram): ")
		fmt.Scan(&berat)

		// Memanggil fungsi hitungBiaya untuk menghitung biaya pengiriman
		biayaDasar, biayaTambahan, totalBiaya := hitungBiaya(berat)

		// Menghitung berat dalam kg dan gram untuk ditampilkan
		beratKg := berat / 1000
		sisaGram := berat % 1000

		// Menampilkan detail berat dan rincian biaya kepada pengguna
		fmt.Printf("Berat parsel (gram): %d\n", berat)
		fmt.Printf("Detail berat: %d kg + %d gram\n", beratKg, sisaGram)
		fmt.Printf("Detail biaya: Rp. %d + Rp. %d\n", biayaDasar, biayaTambahan)
		fmt.Printf("Total biaya: Rp. %d\n\n", totalBiaya)
	}
}
