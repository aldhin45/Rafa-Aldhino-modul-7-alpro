package main

import (
	"fmt"
	"math"
)

type ArrayData struct {
	elements []int
}

func (data *ArrayData) fillArray(size int) {
	data.elements = make([]int, size)
	fmt.Println("Masukkan elemen array:")
	for i := 0; i < size; i++ {
		fmt.Scan(&data.elements[i])
	}
}

func (data *ArrayData) display(filter func(int) bool) {
	for i, v := range data.elements {
		if filter(i) {
			fmt.Print(v, " ")
		}
	}
	fmt.Println()
}

func (data *ArrayData) deleteAtIndex(index int) {
	data.elements = append(data.elements[:index], data.elements[index+1:]...)
}

func (data *ArrayData) average() float64 {
	total := 0
	for _, v := range data.elements {
		total += v
	}
	return float64(total) / float64(len(data.elements))
}

func (data *ArrayData) stdDev() float64 {
	avg := data.average()
	var variance float64
	for _, v := range data.elements {
		variance += math.Pow(float64(v)-avg, 2)
	}
	return math.Sqrt(variance / float64(len(data.elements)))
}

func (data *ArrayData) frequency(target int) int {
	count := 0
	for _, v := range data.elements {
		if v == target {
			count++
		}
	}
	return count
}

func main() {
	var size, x, index, target int
	fmt.Println("Masukkan jumlah elemen array:")
	fmt.Scan(&size)

	data := ArrayData{}
	data.fillArray(size)

	fmt.Println("Seluruh elemen array:")
	data.display(func(i int) bool { return true })

	fmt.Println("Elemen dengan indeks ganjil:")
	data.display(func(i int) bool { return i%2 != 0 })

	fmt.Println("Elemen dengan indeks genap:")
	data.display(func(i int) bool { return i%2 == 0 })

	fmt.Println("Masukkan nilai x untuk indeks kelipatan:")
	fmt.Scan(&x)
	fmt.Println("Elemen dengan indeks kelipatan", x, ":")
	data.display(func(i int) bool { return i%x == 0 })

	fmt.Println("Masukkan indeks yang ingin dihapus:")
	fmt.Scan(&index)
	data.deleteAtIndex(index)
	fmt.Println("Array setelah penghapusan:", data.elements)

	fmt.Printf("Rata-rata elemen: %.2f\n", data.average())
	fmt.Printf("Standar deviasi elemen: %.2f\n", data.stdDev())

	fmt.Println("Masukkan angka untuk mencari frekuensinya:")
	fmt.Scan(&target)
	fmt.Printf("Frekuensi angka %d: %d\n", target, data.frequency(target))
}
