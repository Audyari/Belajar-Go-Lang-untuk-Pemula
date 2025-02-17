package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {

	Eko := Man{
		Name: "Audyari",
	}

	fmt.Println(Eko)

	Eko.Married()

	fmt.Println(Eko)
}
