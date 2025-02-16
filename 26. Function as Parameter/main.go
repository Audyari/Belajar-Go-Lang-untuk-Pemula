package main

import "fmt"

// go run .\main.go

type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {
	fmt.Println("Hello", filter(name))
}

func spamFilter(name string) string {
	if name == "anjing" {
		return "..."
	} else {
		return name
	}
}

func main() {

	sayHelloWithFilter("Audyari", spamFilter)

	filter := spamFilter
	sayHelloWithFilter("anjing", filter)

	sayHelloWithFilter("Wiyono", filter)
}
