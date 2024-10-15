package main

import "fmt"

func main() {

	fmt.Println("Masukkan 5 angka :")
	var a, b, c, d, e int
	fmt.Scan(&a, &b, &c, &d, &e)

	fmt.Println("Output dari 5 angka :")
	fmt.Printf("%c%c%c%c%c\n", a, b, c, d, e)

	fmt.Println("Masukkan 3 karakter :")
	var input string
	fmt.Scanln(&input)

	if len(input) == 3 {
		fmt.Println("Output dari 3 karakter :")
		fmt.Printf("%c%c%c\n", rune(input[0])+1, rune(input[1])+1, rune(input[2])+1)
	} else {
		fmt.Println("Input karakter tidak boleh lebih dari 3")
	}
}
