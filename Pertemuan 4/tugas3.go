package main

import (
	"fmt"
	"strings"
)

// Fungsi untuk menghitung total soal yang diselesaikan dan total waktu yang dibutuhkan
func hitungSkor(nama string, skor *int, soal *int) {
	var waktu int
	// Variabel untuk menghitung soal yang diselesaikan dan skor
	*soal = 0
	*skor = 0

	// Membaca untuk menghitung waktu penyelesain dari peserta
	for i := 0; i < 8; i++ {
		fmt.Scan(&waktu)  // Membaca waktu untuk setiap soal

		// Cek apakah soal diselesaikan (waktu ≤ 300 menit)
		if waktu <= 300 { 
			*soal++        // Menambah jumlah soal yang diselesaikan
			*skor += waktu // Menambahkan waktu ke total skor
		}
	}
}

func main() {
	// Deklarasi variabel untuk menyimpan nama peserta, skor, dan soal
	var nama string
	var totalSkor, totalSoal int
	var pemenang string
	var pemenangSkor, pemenangSoal int = 301, 0 // Inisialisasi pemenang dengan waktu maksimal

	
	for {
		// Meminta input dari pengguna
		fmt.Println("Masukkan nama peserta dan 8 kali menyelesaikan soal (atau ketik 'selesai' untuk mengakhiri):")
		fmt.Scan(&nama)

		// Jika input adalah "Selesai", maka input berhenti
		if strings.ToLower(nama) == "selesai" {
			break
		}

		// Memanggil prosedur untuk menghitung skor
		hitungSkor(nama, &totalSkor, &totalSoal)

		// Tentukan pemenang berdasarkan jumlah soal yang diselesaikan dan total skor
		if totalSoal > pemenangSoal || (totalSoal == pemenangSoal && totalSkor < pemenangSkor) {
			pemenang = nama
			pemenangSoal = totalSoal
			pemenangSkor = totalSkor
		}
	}

	// Output nama pemenang, jumlah soal yang diselesaikan, dan total skor
	fmt.Printf("%s %d %d\n", pemenang, pemenangSoal, pemenangSkor)
}