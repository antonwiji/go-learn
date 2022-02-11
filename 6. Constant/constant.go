package main

import "fmt"

func main() {
	// cara 1
	const name string = "Anton wijaya"
	const age = 13

	fmt.Println(name)
	fmt.Println(age)

	// Multiple Constant
	const (
		nama = "wijaya anton"
		umur = 17
	)

	fmt.Println(nama)
	fmt.Println(umur)
}
