package main

import "fmt"

func insertionSortAsc(arr []int){
	for i := 1; i < len(arr); i++{
		temp := arr[i]
		j := i
		for j > 0 && temp < arr[j-1]{
			arr[j] = arr[j-1]
			j--
		}
		arr[j] = temp
	}
}

func main(){
	data := []int{}
	var x int 

	for{
		fmt.Scan(&x)
		if x < 0 {
			break
		}
		data = append(data, x)
	}

	if len(data) == 0{
		return
	}

	insertionSortAsc(data)

	for i, val := range data{
		if i > 0{
			fmt.Print(" ")
		}
		fmt.Print(val)
	}
	fmt.Println()

	jarak := data[1] - data[0]
	tetap := true

	for i := 2; i < len(data); i++{
		if data[i]-data[i-1] != jarak{
			tetap = false
			break
		}
	}

	if tetap{
		fmt.Printf("Data berjarak %d\n", jarak)
	}else{
		fmt.Println("Data berjarak tidak tetap")
	}
}