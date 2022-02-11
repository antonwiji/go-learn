package main

import "fmt"

func sayHello(firstName string, lastName string, age int) {
	fmt.Println("Hallo", firstName, lastName, age)
}

func main() {
	sayHello("anton", "wijaya", 23)
}
