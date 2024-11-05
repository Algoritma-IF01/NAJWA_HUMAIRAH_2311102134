package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Batas maksimum elemen array
const NMAX int = 127 // batas maksimum elemen dalam array `t`

// tipe data tabel untuk menyimpan karakter
type tabel [NMAX]rune 

// Mengisi array `t` dengan karakter dari input
func isiArray(t *tabel, n *int) {
	var input string
	fmt.Print("Teks : ")
	scanner := bufio.NewScanner(os.Stdin) // Membuat scanner untuk membaca input dari pengguna
	if scanner.Scan() {
		input = scanner.Text() // Membaca input sebagai satu baris teks
	}

	input = strings.ToUpper(input)        // Mengubah teks menjadi huruf kapital 
	input = strings.TrimSuffix(input, " .") 

	for i, char := range input {
		if i >= NMAX {
			break // Menghentikan loop jika jumlah karakter melebihi batas NMAX
		}
		t[i] = char // Menyimpan karakter ke array `t`
		*n++        
	}
}

// Mencetak isi array `t` sebanyak `n` karakter
func cetakArray(t tabel, n int) {
	for i := 0; i < n; i++ {
		fmt.Print(string(t[i])) // Mengonversi rune ke string dan mencetak karakter satu per satu
	}
	fmt.Println() 
}

// Membalik urutan elemen dalam array `t`
func balikanArray(t *tabel, n int) {
	for i := 0; i < n/2; i++ {
		t[i], t[n-i-1] = t[n-i-1], t[i] // Menukar elemen dari awal dan akhir array
	}
}

// Memeriksa apakah array `t` adalah palindrom
func palindrome(t tabel, n int) bool {
	for i := 0; i < n/2; i++ {
		if t[i] != t[n-i-1] {
			return false // Jika elemen dari depan dan belakang tidak sama, maka bukan palindrom
		}
	}
	return true // Jika semua elemen cocok, maka array adalah palindrom
}

func main() {
	var tab tabel 
	var n int     

	isiArray(&tab, &n)              // Mengisi array dengan input dari pengguna
	fmt.Print("Reverse teks : ")    
	balikanArray(&tab, n)           // Membalikkan urutan karakter dalam array
	cetakArray(tab, n)              // Mencetak array yang telah dibalik

	balikanArray(&tab, n)           // Membalikkan kembali array ke urutan semula
	if palindrome(tab, n) {         // Memeriksa apakah array adalah palindrom
		fmt.Println("Palindrom ? true") 
	} else {
		fmt.Println("Palindrom ? false") 
	}
}
