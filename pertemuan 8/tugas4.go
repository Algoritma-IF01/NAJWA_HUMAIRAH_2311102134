package main

import "fmt"

// Konstanta nMax untuk batas maksimal ID buku
const nMax = 7919

// struct Buku
// Struktur untuk menyimpan data buku
type Buku struct {
	ID        int    // ID Buku
	Judul     string // Judul Buku
	Penulis   string // Penulis Buku
	Penerbit  string // Penerbit Buku
	Eksemplar int    // Jumlah Eksemplar Buku
	Tahun     int    // Tahun Terbit Buku
	Rating    int    // Rating Buku
}

// struct DaftarBuku
// Struktur untuk menyimpan daftar buku beserta jumlah buku yang ada
type DaftarBuku struct {
	Pustaka  []Buku // Slice yang menyimpan daftar buku
	nPustaka int    // Jumlah buku dalam pustaka
}

// fungsi utk DaftarkanBuku
// Fungsi ini digunakan untuk mendata buku-buku baru ke dalam daftar pustaka
func DaftarkanBuku(pustaka *DaftarBuku, n int) {
	for i := 0; i < n; i++ {
		var buku Buku
		// Meminta input untuk setiap buku
		fmt.Printf("Masukkan data buku ke-%d (ID Judul Penulis Penerbit Eksemplar Tahun Rating):\n", i+1)
		fmt.Scan(&buku.ID, &buku.Judul, &buku.Penulis, &buku.Penerbit, &buku.Eksemplar, &buku.Tahun, &buku.Rating)
		// Menambahkan buku ke daftar pustaka
		pustaka.Pustaka = append(pustaka.Pustaka, buku)
	}
	pustaka.nPustaka = n // Menyimpan jumlah pustaka
}

// fungsi utk CetakTerfavorit
// Fungsi ini digunakan untuk menampilkan buku dengan rating tertinggi
func CetakTerfavorit(pustaka DaftarBuku) {
	if pustaka.nPustaka == 0 {
		// Menangani jika tidak ada buku di pustaka
		fmt.Println("Tidak ada buku di perpustakaan.")
		return
	}
	// Menyimpan buku dengan rating tertinggi
	terfavorit := pustaka.Pustaka[0]
	for i := 1; i < len(pustaka.Pustaka); i++ {
		buku := pustaka.Pustaka[i]
		if buku.Rating > terfavorit.Rating {
			// Update jika ada buku dengan rating lebih tinggi
			terfavorit = buku
		}
	}
	// Menampilkan buku terfavorit
	fmt.Printf("Buku terfavorit: %s oleh %s (%s, %d) - Rating: %d\n", terfavorit.Judul, terfavorit.Penulis, terfavorit.Penerbit, terfavorit.Tahun, terfavorit.Rating)
}

// fungsi utk UrutBuku
// Fungsi ini digunakan untuk mengurutkan buku berdasarkan rating menggunakan metode Insertion Sort
func UrutBuku(pustaka *DaftarBuku) {
	// Menggunakan Insertion Sort untuk mengurutkan berdasarkan rating
	for i := 1; i < len(pustaka.Pustaka); i++ {
		key := pustaka.Pustaka[i]
		j := i - 1
		// Geser elemen yang memiliki rating lebih kecil ke kanan
		for j >= 0 && pustaka.Pustaka[j].Rating < key.Rating {
			pustaka.Pustaka[j+1] = pustaka.Pustaka[j]
			j--
		}
		// Tempatkan key (buku) ke posisi yang sesuai
		pustaka.Pustaka[j+1] = key
	}
}

// fungsi utk Cetak5Terbaru
// Fungsi ini digunakan untuk menampilkan 5 buku dengan rating tertinggi
func Cetak5Terbaru(pustaka DaftarBuku) {
	fmt.Println("5 Buku dengan rating tertinggi:")
	// Menampilkan 5 buku pertama (dengan rating tertinggi)
	for i := 0; i < 5 && i < pustaka.nPustaka; i++ {
		buku := pustaka.Pustaka[i]
		fmt.Printf("%s oleh %s (%s, %d) - Rating: %d\n", buku.Judul, buku.Penulis, buku.Penerbit, buku.Tahun, buku.Rating)
	}
}

// fungsi utk mencari buku berdasarkan rating
// Fungsi ini digunakan untuk mencari buku berdasarkan rating tertentu menggunakan Binary Search
func CariBuku(pustaka DaftarBuku, rating int) {
	// Menggunakan Binary Search untuk mencari buku berdasarkan rating
	low, high := 0, len(pustaka.Pustaka)-1
	for low <= high {
		mid := (low + high) / 2
		if pustaka.Pustaka[mid].Rating == rating {
			// Jika rating ditemukan, tampilkan buku tersebut
			buku := pustaka.Pustaka[mid]
			fmt.Printf("Buku ditemukan : %s oleh %s (%s, %d) - Rating : %d, Eksemplar : %d\n", buku.Judul, buku.Penulis, buku.Penerbit, buku.Tahun, buku.Rating, buku.Eksemplar)
			return
		}
		if pustaka.Pustaka[mid].Rating < rating {
			// Jika rating lebih tinggi, cari di bagian kanan
			high = mid - 1
		} else {
			// Jika rating lebih rendah, cari di bagian kiri
			low = mid + 1
		}
	}
	// Jika rating tidak ditemukan
	fmt.Println("Tidak ada buku dengan rating seperti itu.")
}

func main() {
	var pustaka DaftarBuku // Membuat pustaka (daftar buku) kosong

	var n int
	fmt.Print("Masukkan jumlah buku : ")
	fmt.Scan(&n)

	// Daftarkan buku-buku
	DaftarkanBuku(&pustaka, n)

	// Cetak buku terfavorit
	CetakTerfavorit(pustaka)

	// Urutkan buku berdasarkan rating
	UrutBuku(&pustaka)

	// Cetak 5 buku dengan rating tertinggi
	Cetak5Terbaru(pustaka)

	// Cari buku berdasarkan rating
	var rating int
	fmt.Print("Masukkan rating yang ingin dicari : ")
	fmt.Scan(&rating)
	CariBuku(pustaka, rating)
}
