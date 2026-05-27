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

func main(){

	var n int
	fmt.Scan(&n)


	hasil := make([]string, n)

	for daerah := 0; daerah < n; daerah++{
		var m int
		fmt.Scan(&m)
		rumah := make([]int, m)
		for i := 0; i < m; i++{
			fmt.Scan(&rumah[i])
		}
		selectionSortAsc(rumah)

		baris := ""

		for i, val := range rumah{
			if i > 0{
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