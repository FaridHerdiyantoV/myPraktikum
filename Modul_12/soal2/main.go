package main

import "fmt"

const NMAX = 100000
var data [NMAX]int

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	
	isiArray(n)
	
	pos := posisi(n, k)
	
	if pos == -1 {
		fmt.Println("TIDAK ADA")
	} else {
		fmt.Println(pos)
	}
}

func isiArray(n int) {
	for i := 0; i < n; i++ {
		fmt.Scan(&data[i])
	}
}

func posisi(n, k int) int {
	var kr int = 0
	var kn int = n - 1
	var found int = -1
	var med int

	for kr <= kn && found == -1 {
		med = (kr + kn) / 2
		if k > data[med] {
			kr = med + 1 
		} else if k < data[med] {
			kn = med - 1 
		} else {
			found = med 
		}
	}
	
	return found
}