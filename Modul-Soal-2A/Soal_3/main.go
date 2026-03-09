package main

import "fmt"

func main() {

	var r int
	const phi = 3.1415926535

	fmt.Print("Masukkan jari-jari bola: ")
	fmt.Scan(&r)

	volumeBola := (4.0 / 3.0) * phi * float64(r*r*r)
	luasBola := 4 * phi * float64(r*r)

	fmt.Printf("Bola yang memiliki panjang jari-jari %d mempunyai Volume %.2f dan Luas %.2f", r, volumeBola, luasBola)
}