package main

import "fmt"

func main() {
	
	var tahun int
	var isKabisat bool

	fmt.Print("Masukkan tahun: ")
	fmt.Scan(&tahun)

	if (tahun%400 == 0 || (tahun%4 == 0 && tahun%100 != 0)){
		isKabisat = true
		fmt.Println("Tahun", tahun, "adalah tahun kabisat", isKabisat)
	}else{
		isKabisat = false
		fmt.Println("Tahun", tahun, "bukan tahun kabisat", isKabisat)
	}
}