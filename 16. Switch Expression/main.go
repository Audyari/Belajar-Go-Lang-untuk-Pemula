package main

import "fmt"

// go run .\main.go

func main() {

	name := "Audyari"

	switch name {
	case "Audyari":
		fmt.Println("nama saya", name)
	default:
		fmt.Println("nama saya tidak dikenali")
	}

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("nama saya lebih dari 5 karakter", name, "yang memiliki panjang", length)
	default:
		fmt.Println("nama saya kurang dari 5 karakter", name, "yang memiliki panjang", length)
	}

	length := len(name)

	switch length > 5 {
	case true:
		fmt.Println("nama saya lebih dari 5 karakter", name, "yang memiliki panjang", length)
	default:
		fmt.Println("nama saya kurang dari 5 karakter", name, "yang memiliki panjang", length)
	}

}
