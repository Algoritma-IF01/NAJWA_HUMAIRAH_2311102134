package main

import "fmt"

func main() {
    // Loop untuk meminta input dua bilangan
    for j := 0; j < 2; j++ {
        var bilangan int // Variabel untuk menyimpan input bilangan

        // Meminta input bilangan
        fmt.Print("Bilangan : ")
        fmt.Scan(&bilangan)

        // Mencetak faktor dari bilangan
        fmt.Print("Faktor : ")
        for i := 1; i <= bilangan; i++ {
            if bilangan%i == 0 { // Mengecek apakah 'i' adalah faktor dari 'bilangan'
                fmt.Printf("%d ", i)
            }
        }
        fmt.Println() // Baris baru setelah mencetak faktor

        // Menghitung jumlah faktor untuk menentukan apakah bilangan prima
        var jumlahFaktor int = 0 // Menghitung jumlah faktor
        for i := 1; i <= bilangan; i++ {
            if bilangan%i == 0 {
                jumlahFaktor++ // Menambahkan 1 untuk setiap faktor yang ditemukan
            }
        }

        // Mengecek apakah bilangan adalah bilangan prima
        if jumlahFaktor == 2 { // Bilangan prima hanya memiliki 2 faktor (1 dan bilangan itu sendiri)
            fmt.Println("Prima : true")
        } else {
            fmt.Println("Prima : false")
        }

        fmt.Println() // Menampilkan baris baru untuk pemisah antara input
    }
}
