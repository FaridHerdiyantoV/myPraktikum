package main

import (
	"fmt"
)

func hitungSkor(soal, skor *int){
	*soal = 0
	*skor = 0

	for i := 0; i < 8; i++{
		var waktu int
		fmt.Scan(&waktu)

		if waktu < 301{
			*soal = *soal + 1
			*skor = *skor + waktu
		}
	}
}

func main(){

	var nama string
	fmt.Scan(&nama)

	var pemenang string
	var soalTerbaik, skorTerbaik int

	pemenang = nama
	hitungSkor(&soalTerbaik, &skorTerbaik)
	fmt.Scan(&nama)

	for nama != "Selesai"{
		var soal, skor int
		hitungSkor(&soal, &skor)

		if soal > soalTerbaik || (soal == soalTerbaik && skor < skorTerbaik){
			pemenang = nama
			soalTerbaik = soal
			skorTerbaik = skor
		}
		fmt.Scan(&nama)
	}

	fmt.Println(pemenang, soalTerbaik, skorTerbaik)
}