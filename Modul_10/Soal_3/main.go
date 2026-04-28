package main

import "fmt"

type arrBalita [100]float64

func hitungMinMax(arrBerat arrBalita, n int, bMin, bMax *float64){

	*bMin = arrBerat[0]
	*bMax = arrBerat[0]

	for i := 1; i < n; i++{
		if arrBerat[i] < *bMin{
			*bMin = arrBerat[i]
		}
		if arrBerat[i] > *bMax{
			*bMax = arrBerat[i]
		}
	}
}

func rata2(arrBerat arrBalita, n int)float64{
	var totalBerat float64 = 0

	for i := 0; i < n; i++{
		totalBerat += arrBerat[i]
	}
	return totalBerat / float64(n)
}

func main(){
	var n int
	var dataBalita arrBalita
	var min, max float64

	fmt.Print("Masukkan jumlah balita (n): ")
	fmt.Scan(&n)

	for i := 0; i < n; i++{
		fmt.Printf("Masukkan berat balita ke-%d: ", i+1)
		fmt.Scan(&dataBalita[i])
	}

	hitungMinMax(dataBalita, n, &min, &max)
	average := rata2(dataBalita, n)

	fmt.Printf("Berat balita minimum adalah %.2f kg\n", min)
	fmt.Printf("Berat balita maximum adalah %.2f kg\n", max)
	fmt.Printf("Rata-rata berat balita adalah %.2f kg\n", average)
}