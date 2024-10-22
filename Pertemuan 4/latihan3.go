package main

import (
	"fmt"
)

// Fungsi rekursif untuk menghitung deret fibonaci
func fibonaci(n int)int{
	if n == 0 {
		return 0
	} else if n == 1{
		return 1
	} else {
		return fibonaci(n-1) + fibonaci(n-2)
	}
}

func main(){
	// Menampilkan deret fibonaci hingga suku ke-10
	fmt.Println("Deret fibonaci hingga suku ke-10 : ")
	for i := 0; i <= 10; i++ {
		fmt.Printf("Fibonaci(%d) = %d\n", i, fibonaci(i))
	}
}