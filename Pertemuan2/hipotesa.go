package main

import "fmt"

func main() {
	var hipotenusa bool
	var a, b, c int

	fmt.Printf("Masukkan angka pertama :")
	fmt.Scanln(&a)
	fmt.Printf("Masukkan angka kedua :")
	fmt.Scanln(&b)
	fmt.Printf("Masukkan angka ketiga :")
	fmt.Scanln(&c)
	hipotenusa = (c * c) == (a*a + b*b)
	fmt.Printf("sisi c adalah hipotenusa segitiga a,b,c: ", hipotenusa)
}
