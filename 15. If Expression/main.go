package main

import "fmt"

// go run .\main.go

func main() {

	name := "Audyari"

	if name == "Audyari" {
		fmt.Println("nama saya", name)
	} else if name == "wiyono" {
		fmt.Println("nama saya ", name)
	} else {
		fmt.Println("nama saya tidak dikenali")
	}

	if length := len(name); length > 5 {
		fmt.Println("nama saya lebih dari 5 karakter", name, "yang memiliki panjang", length)
	} else {
		fmt.Println("nama saya kurang dari 5 karakter", name, "yang memiliki panjang", length)
	}
}
