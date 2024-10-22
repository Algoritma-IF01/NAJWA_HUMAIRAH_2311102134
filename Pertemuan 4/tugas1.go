package main

import (
	"fmt"
)

// Fungsi f(x) = x^2
func f(x int) int {
	return x * x
}

// Fungsi g(x) = x - 2
func g(x int) int {
	return x - 2
}

// Fungsi h(x) = x + 1
func h(x int) int {
	return x + 1
}

// Fungsi untuk menghitung f(g(h(x)))
func fogoh(x int) int {
	return f(g(h(x)))
}

// Fungsi untuk menghitung g(h(f(x)))
func gohof(x int) int {
	return g(h(f(x)))
}

// Fungsi untuk menghitung h(f(g(x)))
func hofog(x int) int {
	return h(f(g(x)))
}

func main() {
	var a, b, c int

	// Input nilai a, b, c
	fmt.Print("Masukkan nilai a, b, dan c : ")
	fmt.Scan(&a, &b, &c)

	// Tampilkan hasil
	fmt.Println("Fogoh : ", fogoh(a))
	fmt.Println("Gohof : ", gohof(b))
	fmt.Println("Hogof : ", hofog(c))
}
