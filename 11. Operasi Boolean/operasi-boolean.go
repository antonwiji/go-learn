package main

import "fmt"

func main() {

	var slulus = 80
	var sabsen = 75

	var lulus = 80 >= slulus
	var absen = 60 >= sabsen

	var perbandingan = lulus == true && absen == true

	fmt.Println(lulus)
	fmt.Println(absen)
	fmt.Println(perbandingan, "selamat anda lulus")
}
