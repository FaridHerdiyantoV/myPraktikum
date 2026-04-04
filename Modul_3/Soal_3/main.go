	package main

	import (
		"fmt"
		"math"
	)

	func jarak(a, b, c, d float64) float64{
		return math.Sqrt(math.Pow(a-c, 2) + math.Pow(b-d, 2))
	}

	func diDalam(cx, cy, r, x, y float64) bool{
		return jarak(cx, cy, x, y) <= r
	}

	func main(){
		var cx1, cy1, r1 float64
		var cx2, cy2, r2 float64
		var x, y float64
		fmt.Scan(&cx1, &cy1, &r1)
		fmt.Scan(&cx2, &cy2, &r2)
		fmt.Scan(&x, &y)

		if diDalam(cx1, cy1, r1, x, y) && diDalam(cx2, cy2, r2, x, y) {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
		} else if diDalam(cx1, cy1, r1, x, y) {
		fmt.Println("Titik di dalam lingkaran 1")
		} else if diDalam(cx2, cy2, r2, x, y) {
		fmt.Println("Titik di dalam lingkaran 2")
		} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
		}
	}
