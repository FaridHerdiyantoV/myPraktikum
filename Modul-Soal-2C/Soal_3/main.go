package main

import "fmt"

func main() {
	
	var b int
	fmt.Print("Bilangan: ")
	fmt.Scan(&b)

	jumlahFaktor := 0
	fmt.Print("Faktor: ")

	for f := 1; f <= b; f++ {
		if b%f == 0 {
			fmt.Print(f, " ")
			jumlahFaktor++
		}
	}
	fmt.Println()

	if jumlahFaktor == 2 {
		fmt.Println("Prima: true")
	} else {
		fmt.Println("Prima: false")
	}
}