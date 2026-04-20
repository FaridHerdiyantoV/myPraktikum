package main

import (
	"fmt"
	"math"
)

func main() {
	var arr [100]int
	var n int

	fmt.Print("Masukkan jumlah elemen N: ")
	fmt.Scan(&n)

	fmt.Println("Masukkan elemen-elemen array:")
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	fmt.Println("\n--- OUTPUT ---")

	fmt.Print("a. Keseluruhan isi array: ")
	for i := 0; i < n; i++ {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	fmt.Print("b. Elemen indeks ganjil: ")
	for i := 0; i < n; i++ {
		if i%2 != 0 {
			fmt.Print(arr[i], " ")
		}
	}
	fmt.Println()

	fmt.Print("c. Elemen indeks genap: ")
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			fmt.Print(arr[i], " ")
		}
	}
	fmt.Println()

	var x int
	fmt.Print("d. Masukkan nilai x untuk kelipatan indeks: ")
	fmt.Scan(&x)
	fmt.Print("   Elemen indeks kelipatan ", x, ": ")
	for i := 0; i < n; i++ {
		if i%x == 0 {
			fmt.Print(arr[i], " ")
		}
	}
	fmt.Println()

	var idx int
	fmt.Print("e. Masukkan indeks yang ingin dihapus: ")
	fmt.Scan(&idx)
	
	for i := idx; i < n-1; i++ {
		arr[i] = arr[i+1]
	}
	n-- 
	
	fmt.Print("   Isi array setelah dihapus: ")
	for i := 0; i < n; i++ {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	var total float64 = 0
	for i := 0; i < n; i++ {
		total += float64(arr[i])
	}
	rataRata := total / float64(n)
	fmt.Printf("f. Rata-rata: %.2f\n", rataRata)

	var totalVarian float64 = 0
	for i := 0; i < n; i++ {
		selisih := float64(arr[i]) - rataRata
		totalVarian += selisih * selisih 
	}
	varian := totalVarian / float64(n)
	standarDeviasi := math.Sqrt(varian)
	fmt.Printf("g. Standar deviasi: %.2f\n", standarDeviasi)

	var cari, frekuensi int
	fmt.Print("h. Masukkan bilangan yang dicari frekuensinya: ")
	fmt.Scan(&cari)
	for i := 0; i < n; i++ {
		if arr[i] == cari {
			frekuensi++
		}
	}
	fmt.Printf("Frekuensi bilangan %d adalah %d kali\n", cari, frekuensi)
}