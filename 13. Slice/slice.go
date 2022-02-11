package main

import "fmt"

func main() {
	var hari = [...]string{
		"senin", "selasa", "rabu", "kamis", "jumat", "sabtu", "minggu",
	}

	slice1 := hari[3:]
	fmt.Println(hari)
	fmt.Println(slice1)
	example := append(slice1, "ganteng nian")
	fmt.Println(hari)
	fmt.Println(example)
}
