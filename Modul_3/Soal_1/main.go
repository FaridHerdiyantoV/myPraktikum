package main
import "fmt"

func factorial(n int) int{
	if n == 0{
		return 1
	}
	return n * factorial(n-1)
}

func permutation(n, r int) int {
	return factorial(n) / factorial(n-r)
}

func combination(n, r int) int{
	return permutation(n, r) / factorial(r)
}

func main() {
	var a, b, c, d int 

	fmt.Scan(&a, &b, &c, &d)

	if a < c || b < d {
		fmt.Println("Error: a >= c dan b >= d harus terpenuhi!")
		return
	}

	fmt.Println(permutation(a, c), combination(a, c))
	fmt.Println(permutation(b, d), combination(b, d))
}
