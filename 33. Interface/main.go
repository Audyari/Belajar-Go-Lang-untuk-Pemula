package main

import "fmt"

// go run .\main.go

type HasName interface {
	GetName() string
	GetAge() int
}

func sayHello(hasName HasName) {
	fmt.Println("Hello", hasName.GetName(), "your age is", hasName.GetAge())
}

type Person struct {
	Name string
	Age  int
}

func (person Person) GetName() string {
	return person.Name
}

func (person Person) GetAge() int {
	return person.Age
}

func main() {

	person := Person{
		Name: "Audyari",
		Age:  21,
	}

	sayHello(person)

}
