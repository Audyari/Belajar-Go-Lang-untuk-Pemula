package main

import "fmt"

// go run .\main.go

func main() {
	fmt.Println("Hello, World!")
	fmt.Println("Audy")
	fmt.Println("Audyari W")
	fmt.Println("Audyari Wiyono")

	fmt.Println("=================")
	fmt.Println(len("Audyari Wiyono"))

	fmt.Println("=================")
	// string di go itu array of byte, maka "a"[0] akan mengembalikan nilai byte dari karakter a
	fmt.Println("ab"[0])
	var a byte = "ab"[0]
	var b string = string(a)
	fmt.Println(b)
}
