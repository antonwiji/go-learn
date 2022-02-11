package main

import "fmt"

type HasName interface {
	GetName() string
}

type Person struct {
	Name string
}

func sayHello(hasname HasName) {
	fmt.Println("hallo", hasname.GetName())
}

func (person Person) GetName() string {
	return person.Name
}

func main() {
	regita := Person{
		Name: "Regita Ismail",
	}

	sayHello(regita)

}
