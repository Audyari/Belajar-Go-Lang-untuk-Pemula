package main

import "fmt"

// go run .\main.go

func getHello(name string) string {
	return "Hello, " + name + "!"
}

func main() {
	result := getHello("Audyari")
	fmt.Println(result)

}
