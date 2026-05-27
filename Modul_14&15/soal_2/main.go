package main

import "fmt"

func selectionSortAsc(arr []int){
	n := len(arr)
	for i := 0; i < n-1; i++{
		idxMin := i
		for j := i + 1; j < n; j++{
			if arr[j] < arr[idxMin]{
				idxMin = j
			}
		}
		arr[i], arr[idxMin] = arr[idxMin], arr[i]
	}
}

func selectionSortDesc(arr []int){
	n := len(arr)
	for i := 0; i < n-1; i++{
		idxMax := i
		for j := i + 1; j < n; j++{
			if arr[j] > arr[idxMax]{
				idxMax = j
			}
		}
		arr[i], arr[idxMax] = arr[idxMax], arr[i]
	}
}

func main(){
	var n int
	fmt.Scan(&n)

	hasil := make([]string, n)

	for daerah := 0; daerah < n; daerah++{
		var m int
		fmt.Scan(&m)

		ganjil := []int{}
		genap := []int{}

		for i := 0; i < m; i++{
			var x int
			fmt.Scan(&x)

			if x%2 == 1{
				ganjil = append(ganjil, x)
			}else{
				genap = append(genap, x)
			}
		}

		selectionSortAsc(ganjil)
		selectionSortDesc(genap)


		baris := ""

		for i, val := range ganjil{
			if i > 0 {
				baris += " "
			}
			baris += fmt.Sprint(val)
		}

		for i, val := range genap{
			if len(ganjil) > 0 || i > 0 {
				baris += " "
			}
			baris += fmt.Sprint(val)
		}

		hasil[daerah] = baris
	}

	fmt.Println("Output:")
	for _, baris := range hasil{
		fmt.Println(baris)
	}
}