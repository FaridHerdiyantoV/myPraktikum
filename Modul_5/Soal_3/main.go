package main

import (
	"fmt"
)

func factor(N, pembagi int){
	if pembagi > N{
		return	
	}else{
		if N % pembagi == 0 {
			fmt.Print(pembagi, " ")
		}
		factor(N, pembagi + 1)
	}
}

func main(){
	var n int

	fmt.Print("Masukkan nilai N: ")
	fmt.Scan(&n)
	factor(n, 1)
}