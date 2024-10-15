package main 

import (
	"fmt"
	"math"
)

func main(){
	var radius float64

	fmt.Print("Masukkan jari-jari bola :")
	fmt.Scan(&radius)

	volume := (4.0/3.0) * math.Pi * math.Pow(radius, 3)
	luasPermukaan := (4.0) * math.Pi * math.Pow(radius, 2)

	fmt.Printf ("Bola dengan jari-jari %.2f memiliki volume %.4f dan luas kulit %4f\n", radius, volume, luasPermukaan)
}