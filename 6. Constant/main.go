package main

import "fmt"

// go run .\main.go

func main() {

	const namadepan string = "Audyari"
	const namabelakang string = "Wiyono"

	fmt.Println(namadepan, namabelakang)

	const (
		firstname = "Audyari"
		lastname  = "Wiyono"
	)

	fmt.Println(firstname, lastname)
}
