package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "Anton"
	names[1] = "wijaya"
	names[2] = "Riski"
	fmt.Println(names)

	var ages = [3]int{
		13,
		15,
		14,
	}

	fmt.Println(ages)
}
