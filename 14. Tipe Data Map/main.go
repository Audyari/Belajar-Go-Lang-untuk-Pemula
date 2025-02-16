package main

import "fmt"

// go run .\main.go

func main() {

	person := map[string]string{
		"name":    "Audyari",
		"address": "Malang",
	}

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	book := make(map[string]string)
	book["title"] = "Belajar Go-Lang"
	book["author"] = "Audyari"
	book["year"] = "2024"

	delete(book, "year")

	fmt.Println(book)
	fmt.Println(book["title"])
	fmt.Println(book["author"])
	// fmt.Println(book["year"])

}
