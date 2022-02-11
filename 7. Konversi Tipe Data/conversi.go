package main

import "fmt"

func main() {
	// convert int
	var nilai1 int32 = 3000
	var nilai2 int64 = int64(nilai1)
	var nilai3 int8 = int8(nilai1)

	fmt.Println(nilai1)
	fmt.Println(nilai2)
	fmt.Println(nilai3)

	// convert bynaery to string
	var nama = "anton wijaya"
	enama := nama[1]
	getnama := string(enama)

	fmt.Println(nama)
	fmt.Println(getnama)
}
