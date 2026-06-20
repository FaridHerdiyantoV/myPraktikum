package main

import (
	"fmt"

	"mesin-abstrak/domino"
	"mesin-abstrak/karakter"
)

func main() {
	fmt.Println("===================================")
	fmt.Println("      PENGUJIAN MESIN DOMINO       ")
	fmt.Println("===================================")

	tumpukan := domino.InisialisasiDomino()
	fmt.Printf("Jumlah kartu awal: %d\n", tumpukan.Sisa)

	domino.KocokKartu(&tumpukan)
	fmt.Println("Kartu telah dikocok!")

	kartuKu := domino.AmbilKartu(&tumpukan)
	fmt.Printf("Kartu yang diambil: (%d, %d), Balak: %t\n", kartuKu.Sisi1, kartuKu.Sisi2, kartuKu.Balak)
	fmt.Printf("Sisa kartu di tumpukan: %d\n", tumpukan.Sisa)

	fmt.Printf("Total nilai kartuKu: %d\n", domino.NilaiKartu(kartuKu))

	suitTarget := 3
	hasilSuit := domino.GambarKartu(kartuKu, suitTarget)
	if hasilSuit != -1 {
		fmt.Printf("Kartu cocok dengan suit %d, sisi lainnya adalah %d\n", suitTarget, hasilSuit)
	} else {
		fmt.Printf("Kartu tidak memiliki suit %d\n", suitTarget)
	}

	fmt.Println("\n--- Menggali Kartu ---")
	kartuCocok := domino.GaliKartu(&tumpukan, kartuKu)
	if kartuCocok.Sisi1 != -1 {
		fmt.Printf("Berhasil menggali kartu yang cocok: (%d, %d)\n", kartuCocok.Sisi1, kartuCocok.Sisi2)
	} else {
		fmt.Println("Tumpukan habis, tidak ada kartu yang cocok.")
	}

	fmt.Println("\n--- Mengecek Sepasang Kartu ---")
	kartuA := domino.Domino{Sisi1: 6, Sisi2: 6, Balak: true}
	kartuB := domino.Domino{Sisi1: 0, Sisi2: 0, Balak: true}
	fmt.Printf("Apakah kartu (6,6) dan (0,0) bernilai total 12? %t\n", domino.SepasangKartu(kartuA, kartuB))

	fmt.Println("\n\n===================================")
	fmt.Println("     PENGUJIAN MESIN KARAKTER      ")
	fmt.Println("===================================")

	inputPita := "ALAM LESTARI DAMAI BERSERI."
	fmt.Printf("Pita karakter yang dibaca: \"%s\"\n\n", inputPita)
	fmt.Println("Hasil Analisis:")

	karakter.ProsesMesinKarakter(inputPita)
}
