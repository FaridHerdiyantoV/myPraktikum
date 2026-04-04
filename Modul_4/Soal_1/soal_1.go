package main

import (
	"fmt"
)

func factorial(n int, hasil *int){
	*hasil = 1
	for i := 1; i <= n; i++{
		*hasil = *hasil * i
	}
}

func permutation(n, r int, hasil *int){
	var factN, factNR int
	factorial(n, &factN)
	factorial(n-r, &factNR)
	*hasil = factN / factNR
}

func combination(n, r int, hasil *int){
	var factN, factR, factNR int
	factorial(n, &factN)
	factorial(r, &factR)
	factorial(n-r, &factNR)
	*hasil = factN / (factR * factNR)
}

func main(){

	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)

	if a >= c && b >= d{
		var perm1, perm2, comb1, comb2 int
		permutation(a, c, &perm1)
		permutation(b, d, &perm2)
		combination(a, c, &comb1)
		combination(b, d, &comb2)

		fmt.Println(perm1, comb1)
		fmt.Println(perm2, comb2)
	}
}