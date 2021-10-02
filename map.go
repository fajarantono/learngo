package main

import "fmt"

func main() {
	var harga_makanan = map[string]int{"ayam_goreng": 20000, "sate_ayam": 15000}
	fmt.Println("sate ayam ", harga_makanan["sate_ayam"])
}
