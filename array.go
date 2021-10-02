package main

import "fmt"

func main() {
	var buah = [4]string{"apel", "jeruk", "anggur", "melon"}
	buah[1] = "mangga"

	fmt.Println("jumlah element :", len(buah))
	fmt.Println("Isi element :", buah)
}
