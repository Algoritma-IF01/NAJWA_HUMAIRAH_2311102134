package main

import (
	"fmt"
)

func main() {
	// Masukkan 5 data integer
	var a, b, c, d, e int
	fmt.Scanf("%d %d %d %d %d", &a, &b, &c, &d, &e)

	// Masukkan 3 karakter
	var chars [3]rune
	fmt.Scanf(" %c %c %c", &chars[0], &chars[1], &chars[2])

	// Cetak hasil pertama: 5 karakter hasil konversi dari integer
	fmt.Printf("%c %c %c %c %c\n", a, b, c, d, e)

	// Cetak hasil kedua: 3 karakter yang sudah dimasukkan
	fmt.Printf("%c %c %c\n", chars[0], chars[1], chars[2])
}
