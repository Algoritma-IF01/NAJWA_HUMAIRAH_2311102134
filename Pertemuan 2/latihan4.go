package main

import "fmt"

func main() {
	var celsius float64

	//Meminta input suhu dalam celcius
	fmt.Print("Masukkan temperatur celcius : ")
	fmt.Scanln(&celsius)

	//Menghitung suhu dalam berbagai satuan
	fahrenheit := (celsius * 9 / 5) + 32
	reamur := celsius * 4 / 5
	kelvin := (fahrenheit + 459.67) * 5 / 9

	fmt.Printf("Derajat Fahrenheit : %.2f\n", fahrenheit)
	fmt.Printf("Derajat Reamur : %.2f\n", reamur)
	fmt.Printf("Derajat Kelvin : %.2f\n", kelvin)
}
