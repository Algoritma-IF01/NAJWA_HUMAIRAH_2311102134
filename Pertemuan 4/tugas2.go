package main

import (
	"fmt"
	"math"
)

// Fungsi untuk menghitung jarak antara dua titik
func jarak(a, b, c, d int) float64 {
	return math.Sqrt(float64((a-c)*(a-c) + (b-d)*(b-d)))
}

// Fungsi untuk menentukan apakah titik (x, y) berada di dalam lingkaran dengan pusat (cx, cy) dan radius r
func dalamLingkaran(x, y, cx, cy, r int) bool {
	return jarak(x, y, cx, cy) <= float64(r)
}

func main() {
	var x1, y1, r1, x2, y2, r2, px, py int

	// Input koordinat dan radius lingkaran 1
	fmt.Print("Masukkan x1, y1, dan radius lingkaran 1 : ")
	fmt.Scan(&x1, &y1, &r1)

	// Input koordinat dan radius lingkaran 2
	fmt.Print("Masukkan x2, y2, dan radius lingkaran 2 : ")
	fmt.Scan(&x2, &y2, &r2)

	// Input titik sembarang
	fmt.Print("Masukkan koordinat titik : ")
	fmt.Scan(&px, &py)

	// Cek apakah titik berada di dalam lingkaran 1 atau lingkaran 2
	didalam1 := dalamLingkaran(px, py, x1, y1, r1)
	didalam2 := dalamLingkaran(px, py, x2, y2, r2)

	// Output hasil pengecekan
	if didalam1 && didalam2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if didalam1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if didalam2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
