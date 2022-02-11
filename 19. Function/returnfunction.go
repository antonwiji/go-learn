package main

import "fmt"

func sayHello(name string) string {
	if name == "" {
		return "halo bro"
	} else {
		return "hallo" + name
	}
}
func main() {

	result := sayHello(" anton")
	fmt.Println(result)
	fmt.Println(sayHello(""))

}
