package main

import (
	"fmt"
)

func printStar(total int){
	if total == 0{
		fmt.Println()
	}else{
		fmt.Print("*")
		printStar(total - 1)
	}
}

func starPattern(line, N int){
	if line > N {
		return 
	}else{
		printStar(line)
		starPattern(line + 1, N)
	}
}

func main(){
	var n int

	fmt.Print("Masukkan N: ")
	fmt.Scan(&n)
	starPattern(1, n)
}