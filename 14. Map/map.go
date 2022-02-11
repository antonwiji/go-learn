package main

import "fmt"

func main() {
	// cara 1
	book := map[string]string{
		"judul":    "belajar golang",
		"author":   "Anton wiji",
		"category": "Progreming",
	}

	book["editor"] = "wiji geming"

	fmt.Println(book["judul"])

	// cara ke 2
	var person map[string]string = make(map[string]string)

	person["nama"] = "anton wijaya"
	person["alamat"] = "sekojo"

	fmt.Println(person)
	fmt.Println(person["nama"])
	fmt.Println(person["alamat"])

}
