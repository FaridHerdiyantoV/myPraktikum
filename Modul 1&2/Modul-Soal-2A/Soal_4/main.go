package main

import "fmt"

func main() {

	var celsius float64

	fmt.Print("Masukkan suhu (derajat Celcius): ")
	fmt.Scan(&celsius)

	fahrenheit := (celsius * 9 / 5) + 32
    reamur := celsius * 4 / 5
    kelvin := (fahrenheit + 459.67) * 5 / 9

    fmt.Printf("Derajat Reamur: %.0f\n", reamur)
    fmt.Printf("Derajat Fahrenheit: %.0f\n", fahrenheit)
    fmt.Printf("Derajat Kelvin: %.0f\n", kelvin)

}