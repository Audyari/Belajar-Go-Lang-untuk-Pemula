package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func changeAddressToIndonesia(address *Address) {
	address.City = "Jakarta"
	address.Province = "Jawa Barat"
	address.Country = "Indonesia"

}

func main() {
	address1 := Address{"Bogor", "bekasi", "jawa barat"}

	changeAddressToIndonesia(&address1)

	fmt.Println(address1)
}
