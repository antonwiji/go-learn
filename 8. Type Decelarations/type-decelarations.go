package main

import "fmt"

func main() {
	type nama string
	type nikah bool

	var anton nama = "anton wijaya"
	var status nikah = false

	fmt.Println(anton)
	if status == false {

		fmt.Println("belum Menikah")
	} else {
		fmt.Println("sudah Menikah")
	}

}
