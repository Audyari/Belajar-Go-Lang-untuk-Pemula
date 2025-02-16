package main

import "fmt"

// go run .\main.go

func main() {

	counter := 0

	increase := func() {
		fmt.Println("counter = ", counter)
		counter++

	}

	increase()
	increase()
	increase()

	fmt.Println("counter = ", counter)
}
