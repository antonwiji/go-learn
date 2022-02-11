package main

import "fmt"

func main() {

	var nama1 = "anton"
	var nama2 = "wijaya"
	var result bool = nama1 == nama2
	fmt.Println(result)

	a := 7
	b := 5
	var kecildari bool = a < b
	var lebihdari bool = a > b
	var samaldari bool = a == b
	var tidakdari bool = a != b

	fmt.Println(kecildari)
	fmt.Println(lebihdari)
	fmt.Println(samaldari)
	fmt.Println(tidakdari)
}
