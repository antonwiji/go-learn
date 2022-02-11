package main

import "fmt"

func main() {
	var nama = "antonee"

	length := len(nama)

	switch {
	case length > 10:
		fmt.Println("Nama terlalu panjang")
	case length > 5:
		fmt.Println("nama sudah benar")
	default:
		fmt.Println("nama pendek sekali")
	}
}
