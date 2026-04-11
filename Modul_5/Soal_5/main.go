	package main

	import (
		"fmt"
	)

	func odd(bilangan, n int){
		if bilangan > n{
			return 
		}else{
			if bilangan % 2 != 0 {
				fmt.Print(bilangan, " ")
			}
			odd(bilangan + 2, n)
		}
	}

	func main(){
		var n int 

		fmt.Print("Masukkan angka (N): ")
		fmt.Scan(&n)
		odd(1, n)
	}