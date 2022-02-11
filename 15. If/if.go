package main

import "fmt"

func main() {
	nama := "anton wew"

	if nama == "anton" {
		fmt.Println("hallo Anton")
	} else if nama == "budi" {
		fmt.Println("hallo budi")
	} else {
		fmt.Println("nama kamu Siapa ?")
	}

	if panjang := len(nama); panjang > 5 {
		fmt.Println("nama", nama, "terlalu panjang")
	} else {
		fmt.Println("nama", nama, "sudah benar")
	}
}
