package main

import "fmt"

// go run .\main.go

func getGoodbye(name string) string {
	return "Goodbye " + name
}

func main() {

	goodbye := getGoodbye
	fmt.Println(goodbye("Audyari"))

}
