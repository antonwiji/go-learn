package main

import "fmt"

func main() {
	for i := 1; i <= 5; i++ {
		fmt.Println("ini pengulangan yang ke-", i)
	}

	slice := []string{"anton", "wijaya", "kusuma"}

	for i, value := range slice {
		fmt.Println(i, "nama saya", value)
	}

	maping := make(map[string]string)

	maping["nama"] = "anton wijaya"
	maping["judul"] = "oke lah klw begitu"
	maping["nama"] = "gokil"
	maping["judul"] = "oke lahwewe"

	for key, veluem := range maping {
		fmt.Println(key, "adalah", veluem)
	}
}
