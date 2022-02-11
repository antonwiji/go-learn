package main

import "fmt"

func main() {
	// cara 1
	var name string
	name = "Anton Wijaya"
	fmt.Println(name)

	// cara 2
	var teman = "Ilham Zahri"
	fmt.Println(teman)
	teman = "Riski Kurniawan"
	fmt.Println(teman)

	// cara 3
	negara := "indonesia"
	fmt.Println(negara)

	//Multiple Variabel
	var (
		nama = "wijaya anton"
		age  = 23
	)
	fmt.Println(nama)
	fmt.Println(age)
}
