package main 

import "fmt"

const nMax = 7919

type Buku struct{
	id, judul, penulis, penerbit string
	jumlahEksemplar, tahunTerbit, ratingBuku int 
}

type daftarBuku [nMax]Buku

func DaftarkanBuku(pustaka *daftarBuku, n *int){
	fmt.Scan(n)
	for i := 0; i < *n; i++{
		fmt.Scan(&pustaka[i].id, &pustaka[i].judul, &pustaka[i].penulis, &pustaka[i].penerbit, 
			&pustaka[i].jumlahEksemplar, &pustaka[i].tahunTerbit, &pustaka[i].ratingBuku)
	}
}

func CetakTerfavorit(pustaka daftarBuku, n int){
	if n == 0{
		return
	}

	maxIdx := 0

	for i := 1; i < n; i++{
		if pustaka[i].ratingBuku > pustaka[maxIdx].ratingBuku{
			maxIdx = i
		}
	}
	buku := pustaka[maxIdx]
	fmt.Println("--- Buku Terfavorit ---")
	fmt.Printf("Judul: %s\nPenulis: %s\nPenerbit: %s\nTahun Terbit: %d\n", buku.judul, buku.penulis, buku.penerbit, buku.tahunTerbit)
	fmt.Println()
}

func UrutBuku(pustaka *daftarBuku, n int){
	for i := 1; i < n; i++{
		temp := pustaka[i]
		j := i

		for j > 0 && temp.ratingBuku > pustaka[j-1].ratingBuku{
			pustaka[j] = pustaka[j-1]
			j--
		}
		pustaka[j] = temp
	}
}

func CetakTBerbaru(pustaka daftarBuku, n int){
	limit := 5
	if n < 5{
		limit = n
	}

	fmt.Println("--- 5 Buku Rating Tertinggi ---")
	for i := 0; i < limit; i++{
		fmt.Printf("%d. %s (Rating: %d)\n", i+1, pustaka[i].judul, pustaka[i].ratingBuku)
	}
	fmt.Println()
}

func CariBuku(pustaka daftarBuku, n int, r int){

	left, right := 0, n-1
	found := -1

	for left <= right {
		mid := (left + right) / 2
		if pustaka[mid].ratingBuku == r{
			found = mid
			break
		}else if pustaka[mid].ratingBuku > r{
			left = mid + 1
		}else{
			right = mid - 1
		}
	}

	fmt.Printf("--- Hasil Pencarian Buku dengan Rating %d ---\n", r)
	if found == -1{
		fmt.Println("Tidak ada buku dengan rating seperti yang diminta")
	}else{
		buku := pustaka[found]
		fmt.Printf("Judul: %s\nPenulis: %s\nPenerbit: %s\nTahun Terbit: %d\nJumlah Eksemplar: %d\nRating: %d\n", buku.judul, buku.penulis, buku.penerbit, buku.tahunTerbit, buku.jumlahEksemplar, buku.ratingBuku)
	}
}

func main(){
	var pustaka daftarBuku
	var n int

	DaftarkanBuku(&pustaka, &n)
	fmt.Println()
	CetakTerfavorit(pustaka, n)
	UrutBuku(&pustaka, n)
	CetakTBerbaru(pustaka, n)

	var cariRating int
	fmt.Print("Masukkan rating buku yang ingin dicari: ")
	fmt.Scan(&cariRating)
	CariBuku(pustaka, n, cariRating)
}