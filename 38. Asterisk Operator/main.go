package main

import "fmt"

// go run .\main.go

type Address struct {
	City, Province, Country string
}

func main() {
	address1 := Address{"Bogor", "Jawa Barat", "Indonesia"}
	address2 := &address1

	address2.City = "Jakarta"

	*address2 = Address{"Cimahi", "Jawa Barat", "Indonesia"}

	fmt.Println(address1)
	fmt.Println(*address2)
}
