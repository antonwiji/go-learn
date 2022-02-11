package main

import "fmt"

type Akun struct {
	UserName, Name string
	Id             int
}

func (akun Akun) sayHi(name string) {
	fmt.Println("Hallo", akun.Name, "My Name", name, "Anda Login Dengan Username", akun.UserName)
}
func main() {
	Anton := Akun{
		UserName: "leraku12",
		Name:     "Anton Wijaya",
		Id:       23,
	}
	Anton.sayHi("Riski")

	Regita := Akun{
		UserName: "regita12",
		Name:     "Regita Ismail",
	}

	Regita.sayHi("Anton")

}
