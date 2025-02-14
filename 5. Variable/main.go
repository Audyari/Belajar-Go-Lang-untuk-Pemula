package main

import "fmt"

// go run .\main.go

func main() {
	var name string
	var age int

	name = "Audy"
	age = 20

	fmt.Println(name)
	fmt.Println(age)

	name = "Audyari"
	age = 21

	fmt.Println(name)
	fmt.Println(age)

	var namadepan string = "Audyari"
	var namabelakang string = "Wiyono"

	fmt.Println(namadepan)
	fmt.Println(namabelakang)

	fmt.Println("==================")

	namabaca := "Audyari Wiyono"

	fmt.Println(namabaca)

	fmt.Println("==================")

	var (
		contohnamadepan    string = "Audyari"
		contohnamabelakang string = "Wiyono"
	)

	fmt.Println(contohnamadepan)
	fmt.Println(contohnamabelakang)
}
