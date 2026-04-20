package main

import (
	"fmt"
)

const NMAX int = 127

type tabel [NMAX]rune

func isiArray(t *tabel, n *int) {
	var char string
	*n = 0
	for *n < NMAX {
		fmt.Scan(&char)
		if char == "." {
			break
		}
		(*t)[*n] = rune(char[0])
		*n++
	}
}

func cetakArray(t tabel, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%c ", t[i])
	}
	fmt.Println()
}

func balikanArray(t *tabel, n int) {
	for i := 0; i < n/2; i++ {
		temp := (*t)[i]
		(*t)[i] = (*t)[n-1-i]
		(*t)[n-1-i] = temp
	}
}

func palindrom(t tabel, n int) bool {
	
	var aslinya tabel = t
	balikanArray(&t, n)
	
	for i := 0; i < n; i++ {
		if aslinya[i] != t[i] {
			return false 
		}
	}
	return true 
}

func main() {
	var tab tabel
	var m int

	fmt.Print("Teks\t\t: ")
	isiArray(&tab, &m)

	hasilPalindrom := palindrom(tab, m)
	fmt.Printf("Palindrom\t? %t\n", hasilPalindrom)
}