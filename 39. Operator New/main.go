package main

import "fmt"

// go run .\main.go

type Address struct {
	City, Province, Country string
}

func main() {
	alamat1 := new(Address)
	alamat2 := alamat1

	alamat2.City = "Jakarta"

	fmt.Println(*alamat1)
	fmt.Println(*alamat2)
}
