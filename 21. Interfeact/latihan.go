package main

import "fmt"

type Akun struct {
	Name, Adress string
	Age          int
}

type AkunProfile interface {
	GetName() string
	GetAdress() string
	GetAge() int
}

func (akun Akun) GetName() string {
	return akun.Name
}

func (akun Akun) GetAdress() string {
	return akun.Adress
}

func (akun Akun) GetAge() int {
	return akun.Age
}

func sayHaloAkun(akunProfile AkunProfile) {
	fmt.Println("hello", akunProfile.GetName(), "Anda Beralamat", akunProfile.GetAdress(), "dan Umur Anda adalah", akunProfile.GetAge())
}

func main() {

	Anton := Akun{
		Name:   "Anton Wijaya",
		Adress: "jln. Printis Kemerdekaan",
		Age:    21,
	}

	Regita := Akun{
		Name:   "Regita Ismail",
		Adress: "jln Ramahkasih 6",
		Age:    20,
	}

	sayHaloAkun(Anton)
	sayHaloAkun(Regita)

}
