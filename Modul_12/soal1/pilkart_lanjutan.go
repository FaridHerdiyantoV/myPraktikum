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

	counts[0] = -1 
	ketua := 0
	wakil := 0

	for i := 1; i <= 20; i++ {
		if counts[i] > counts[ketua] {
			wakil = ketua
			ketua = i
		} else if counts[i] > counts[wakil] {
			wakil = i
		}
	}

	if ketua != 0 {
		fmt.Printf("Ketua RT: %d\n", ketua)
	}
	if wakil != 0 {
		fmt.Printf("Wakil ketua: %d\n", wakil)
	}
}