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

func median(data []int) int{
	n := len(data)
	if n == 0{
		return 0
	}

	if n%2 == 1{
		return data[n/2]
	}

	return (data[n/2-1] + data[n/2]) / 2
}

func main(){
	data := []int{}

	var x int

	for{
		fmt.Scan(&x)

		if x == -5313{
			break
		}

		if x == 0{
			insertionSortAsc(data)
			fmt.Println(median(data))
		}else{
			data = append(data, x)
		}
	}
}