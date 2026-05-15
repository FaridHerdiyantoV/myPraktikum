package main

import "fmt"

func main() {
	var input int
	var totalSuara, suaraSah int
	var counts [21]int 

	for {
		fmt.Scan(&input)
		if input == 0 { 
			break
		}
		
		totalSuara++
		
		if input >= 1 && input <= 20 {
			suaraSah++
			counts[input]++
		}
	}

	fmt.Printf("Suara masuk: %d\n", totalSuara)
	fmt.Printf("Suara sah: %d\n", suaraSah)
	
	for i := 1; i <= 20; i++ {
		if counts[i] > 0 {
			fmt.Printf("%d: %d\n", i, counts[i])
		}
	}
}