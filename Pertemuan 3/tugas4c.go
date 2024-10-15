package main

import (
	"fmt" 
)

func main() {
	var nam float64 // Variabel untuk menyimpan nilai akhir mata kuliah
	var nmk string  // Variabel untuk menyimpan nilai mata kuliah

	// Menerima input nilai akhir mata kuliah dari pengguna
	fmt.Print("Nilai akhir mata kuliah: ")
	fmt.Scan(&nam) // Mengambil input dari pengguna dan menyimpannya dalam variabel nam

	// Menghitung NMK berdasarkan nilai NAM
	if nam >= 80 {
		nmk = "A" 
	} else if nam > 72.5{
		nmk = "AB"
	} else if nam >= 65 {
		nmk = "B" 
	} else if nam >= 57.5 {
		nmk = "BC" 
	} else if nam >= 50 {
		nmk = "C" 
	} else if nam >= 40 {
		nmk = "D" 
	} else {
		nmk = "E" 
	}

	// Menampilkan hasil
	fmt.Println("Nilai mata kuliah: ", nmk)
}
