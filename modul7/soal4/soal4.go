package main

import (
	"fmt"
)

const MAX = 127

type Table struct {
	arr [MAX]rune
	n   int
}

func isiArray(tab *Table) {
	var input string
	fmt.Print("Masukkan teks : ")
	fmt.Scan(&input)

	tab.n = len(input)
	for i, char := range input {
		tab.arr[i] = char
	}
}

func cetakArray(tab Table) {
	for i := 0; i < tab.n; i++ {
		fmt.Print(string(tab.arr[i]))
	}
	fmt.Println()
}

func balikkanArray(tab *Table) {
	for i, j := 0, tab.n-1; i < j; i, j = i+1, j-1 {
		tab.arr[i], tab.arr[j] = tab.arr[j], tab.arr[i]
	}
}

func isPalindrome(tab Table) bool {
	for i := 0; i < tab.n/2; i++ {
		if tab.arr[i] != tab.arr[tab.n-1-i] {
			return false
		}
	}
	return true
}

func main() {
	var tabel Table

	isiArray(&tabel)

	fmt.Print("Array asli: ")
	cetakArray(tabel)

	balikkanArray(&tabel)
	fmt.Print("Array setelah dibalik: ")
	cetakArray(tabel)

	if isPalindrome(tabel) {
		fmt.Println("Array membentuk palindrome!")
	} else {
		fmt.Println("Array tidak membentuk palindrome.")
	}
}
