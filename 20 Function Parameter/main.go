package main

import "fmt"

// go run .\main.go

func sayHello(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}

func main() {
	sayHello("Audyari", "Wiyono")
}
