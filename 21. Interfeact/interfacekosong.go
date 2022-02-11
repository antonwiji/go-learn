package main

import "fmt"

func Ups(i int) interface{} {
	if i == 1 {
		return 1
	} else if i == 2 {
		return "dua"
	} else {
		return true
	}

}

func main() {
	ups := Ups(4)

	fmt.Println(ups)

}
