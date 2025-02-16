package main

import "fmt"

// go run .\main.go
type BlackList func(string) bool

func registerUser(name string, blackList BlackList) {
	if blackList(name) {
		fmt.Println("Maaf", name, "tidak bisa login")
		return
	}
	fmt.Println("Selamat datang", name)
}

func main() {

	blackList := func(name string) bool {
		return name == "anjing"
	}

	registerUser("anjing", blackList)
	registerUser("Audyari", blackList)

	registerUser("anjing", func(name string) bool {
		return name == "anjing"
	})

	registerUser("Audyari", func(name string) bool {
		return name == "anjing"
	})
}
