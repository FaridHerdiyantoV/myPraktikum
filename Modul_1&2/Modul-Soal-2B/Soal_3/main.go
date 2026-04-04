package main

import "fmt"

func main(){

    for{
        var kantong1, kantong2 float64
        fmt.Print("Masukkan berat kedua kantong: ")
        fmt.Scan(&kantong1, &kantong2)

        if kantong1 + kantong2 > 150 || kantong1 < 0 || kantong2 < 0{
            break   
        }

        selisih := kantong1 - kantong2
        if selisih < 0 {
            selisih = -selisih
        }

        if selisih >= 9 {
            fmt.Println("Sepeda motor pak Andi akan oleng: true")
        }else{
            fmt.Println("Sepeda motor pak Andi akan oleng: false")
        }

    }
    fmt.Println("Proses selesai.")
}
