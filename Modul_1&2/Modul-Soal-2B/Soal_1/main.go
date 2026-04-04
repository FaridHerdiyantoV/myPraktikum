package main

import (
	"fmt"
)

func main(){

    var berhasil bool = true

    for i := 1; i <= 5; i++{
        var gelas1, gelas2, gelas3, gelas4 string
        fmt.Printf("Percobaan ke-%d: ", i)
        fmt.Scanln(&gelas1, &gelas2, &gelas3, &gelas4)

        if gelas1 != "merah" || gelas2 != "kuning" || gelas3 != "hijau" || gelas4 != "ungu"{
            berhasil = false
        }
    }

    if berhasil {
        fmt.Println("Berhasil: true")
    } else{
        fmt.Println("Berhasil: false")
    }
}
