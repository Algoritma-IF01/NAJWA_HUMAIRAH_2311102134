package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	var greetings = "selamat datang di dunia DAP"
	var a, b int

	fmt.Println(greetings)
	fmt.Printf("Masukkan angka pertama :")
	fmt.Scanln(&a)
	fmt.Printf("Masukkan angka kedua :")
	fmt.Scanln(&b)
	fmt.Printf("%v + %v = %v\n", a, b, a+b)
}
