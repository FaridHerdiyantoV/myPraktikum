package main

import "fmt"

type arrKelinci [1000]float64

func main(){
	var n int
	var berat arrKelinci
	var max, min float64

	fmt.Print("Masukkan banyak kelinci (n):")
	fmt.Scan(&n)

	for i:= 0; i < n; i++{
		fmt.Printf("Berat kelinci ke-[%d]: ", i+1)
		fmt.Scan(&berat[i])
	}

	min = berat[0]
	max = berat[0]

	for j := 1; j < n; j++{
		if berat[j] < min{
			min = berat[j]
		}

		if berat[j] > max{
			max = berat[j]
		}
	}

	fmt.Printf("Nilai min dari semua berat kelinci adalah %.2f\n",min)
	fmt.Printf("Nilai max dari semua berat kelinci adalah %.2f\n",max)
}