package main

import "fmt"

func sayHello() (string, string) {
	return "anton", "wijaya"
}
func main() {
	firsName, lastName := sayHello()

	fmt.Println(firsName, lastName)

}
