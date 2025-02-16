package main

import "fmt"

// go run .\main.go

func Ups() interface{} {
	return "Interface Kosong"
}

func main() {
	kosong := Ups()
	fmt.Println(kosong)
}
