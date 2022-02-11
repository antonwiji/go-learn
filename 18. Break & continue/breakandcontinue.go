package main

import "fmt"

func main() {
	// for i := 0; i < 10; i++ {
	// 	if i == 6 {
	// 		break
	// 	}
	// 	fmt.Println("Pengulangan ke", i)
	// }
	for i := 0; i < 10; i++ {
		if i%2 == 1 {
			continue
		}
		fmt.Println("pengulangan ke-", i)
	}
}
