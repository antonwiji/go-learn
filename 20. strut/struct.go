package main

import "fmt"

type Costumer struct {
	Name, Adress string
	Age          int
}

func main() {
	var Anton Costumer
	Anton.Name = "Anton Wijaya"
	Anton.Adress = "Palembang"
	Anton.Age = 21

	fmt.Println(Anton)

	// struct leteratur

	Riski := Costumer{
		Name:   "Riski Kurniawan",
		Adress: "Kuto",
		Age:    20,
	}

	fmt.Println(Riski)

	Regita := Costumer{"Regita Ismail", "Ramah kasih 6", 20}

	fmt.Println(Regita)
}
