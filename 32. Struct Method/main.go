package main

import "fmt"

// go run .\main.go

type customer struct {
	Name, Address string
	Age           int
}

func (customer customer) sayHello(name string) {
	fmt.Println("Hello", name, "my name is", customer.Name)
}

func main() {

	audy := customer{
		Name:    "Audy",
		Address: "Jl. Raya No. 1",
		Age:     20,
	}

	audy.sayHello("Budi")
}
