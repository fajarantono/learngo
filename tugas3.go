package main

import "fmt"

func main() {
	var buah = []string{"apel", "jeruk", "anggur", "melon", "pepaya"}

	fmt.Println("Jumlah Element :", len(buah))
	fmt.Println("Isi Element :", buah)

	var i = 0
	for i < 5 {
		fmt.Println("Element ke - : ", i, buah[i])

		i++
	}
}
