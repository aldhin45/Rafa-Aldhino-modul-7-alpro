package main

import "fmt"

const nMax int = 1000

type Data struct {
	arr1, arr2, hasilOR, hasilAND, hasilXOR [nMax]int
	n                                       int
}

func isiArray(arr *[nMax]int, n int) {
	for i := 0; i < n; i++ {
		for {
			var x int
			fmt.Scan(&x)
			if x == 0 || x == 1 {
				arr[i] = x
				break
			}
			fmt.Println("Masukkan hanya 0 atau 1:")
		}
	}
}

func hitungOperasi(data *Data) {
	for i := 0; i < data.n; i++ {
		data.hasilOR[i] = data.arr1[i] | data.arr2[i]
		data.hasilAND[i] = data.arr1[i] & data.arr2[i]
		data.hasilXOR[i] = data.arr1[i] ^ data.arr2[i]
	}
}

func cetakHasil(nama string, hasil [nMax]int, n int) {
	total := 0
	fmt.Printf("Hasil %s: ", nama)
	for i := 0; i < n; i++ {
		fmt.Print(hasil[i])
		total += hasil[i]
	}
	fmt.Printf("\nTotal %s: %d\n\n", nama, total)
}

func main() {
	var data Data

	fmt.Println("Masukkan jumlah elemen (n):")
	fmt.Scan(&data.n)
	if data.n < 1 || data.n > nMax {
		fmt.Printf("Jumlah elemen harus antara 1 dan %d\n", nMax)
		return
	}

	fmt.Println("Masukkan elemen untuk array 1:")
	isiArray(&data.arr1, data.n)
	fmt.Println("Masukkan elemen untuk array 2:")
	isiArray(&data.arr2, data.n)

	hitungOperasi(&data)
	cetakHasil("OR", data.hasilOR, data.n)
	cetakHasil("AND", data.hasilAND, data.n)
	cetakHasil("XOR", data.hasilXOR, data.n)
}
