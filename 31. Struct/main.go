package main

import "fmt"

// go run .\main.go

type customer struct {
	Name, Address string
	Age           int
}

func main() {
	var audyari customer

	fmt.Println(audyari)

	audyari.Name = "Audyari"
	audyari.Address = "Jl. Raya No. 1"
	audyari.Age = 20

	fmt.Println(audyari)

	joko := customer{
		Name:    "Joko",
		Address: "Jl. Raya No. 2",
		Age:     30,
	}

	fmt.Println(joko)

	budi := customer{"Budi", "Jl. Raya No. 3", 40}

	fmt.Println(budi)
}
