package main

import (
	"fmt"
	"math"
)

func main() {
	var berat1, berat2 float64

	for {

		fmt.Print("Masukan berat belanjaan di kedua kantong : ")
		fmt.Scan(&berat1, &berat2)

		if berat1 < 0 || berat2 < 0 || berat1+berat2 > 150 {
			fmt.Println("Proses selesai.")
			break
		}

		selisih := math.Abs(berat1 - berat2)
		oleng := selisih >= 9

		fmt.Printf("Sepeda motor Pak Andi akan oleng : %t\n", oleng)
	}
}
