package main

import (
	"fmt"
)

func main() {
	var klubA, klubB string
	var pemenang [100]string
	var skorA, skorB int
	var totalPertandingan int

	fmt.Print("Klub A : ")
	fmt.Scan(&klubA)
	fmt.Print("Klub B : ")
	fmt.Scan(&klubB)

	for i := 1; ; i++ {
		fmt.Printf("Pertandingan %d : ", i)
		fmt.Scan(&skorA, &skorB)

		if skorA < 0 || skorB < 0 {
			break
		}

		if skorA > skorB {
			pemenang[totalPertandingan] = klubA
		} else if skorA < skorB {
			pemenang[totalPertandingan] = klubB
		} else {
			pemenang[totalPertandingan] = "Draw"
		}
		
		totalPertandingan++
	}

	fmt.Println() 
	for i := 0; i < totalPertandingan; i++ {
		fmt.Printf("Hasil %d : %s\n", i+1, pemenang[i])
	}
	fmt.Println("Pertandingan selesai")
}