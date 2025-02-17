package main

import (
	"belajar/helper"
	"fmt"
)

// go run main.go
// go build .\main.go

func main() {

	fmt.Println("Program berjalan dengan baik!")

	result := helper.SayHello("Audyari")
	fmt.Println(result)
}
