package main

import "fmt"

// go run .\main.go

func main() {
	var a = 1
	var b = 2
	var c = a + b

	fmt.Println(c)

	var d = 3
	d += 2
	fmt.Println(d)

	var e = 10
	e++
	e++

	fmt.Println(e)
}
