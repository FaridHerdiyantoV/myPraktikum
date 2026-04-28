package main

import "fmt"


func main(){
	var x, y int
	var arrIkan [1000]float64
	var totalSeluruhBerat float64

	fmt.Scan(&x, &y)

	for i := 0; i < x; i++{
		fmt.Scan(&arrIkan[i])
		totalSeluruhBerat += arrIkan[i]
	}

	jumlahWadah := (x + y - 1) / y

	for i := 0; i < x; i+=y{
		var totalWadah float64 = 0

		for j := 0; j < y && (i+j) < x; j++{
			totalWadah += arrIkan[i+j]
		}
		fmt.Printf("%.2f ", totalWadah)
	}
	fmt.Println()

	fmt.Printf("%.2f\n", totalSeluruhBerat/float64(jumlahWadah))

}