package main

import "fmt"

func main(){
	
	var (
		a, b, c string
		temp string
	)

	fmt.Print("Masukkan input string: ")
	fmt.Scan(&a)
	fmt.Print("Masukkan input string: ")
	fmt.Scan(&b)
	fmt.Print("Masukkan input string: ")
	fmt.Scan(&c)

	fmt.Println("Output awal = " + a + " " + b + " " + c)
	temp = a
	a = b
	b = c
	c = temp
	fmt.Println("Output akhir = " + a + " " + b + " " + c)
}