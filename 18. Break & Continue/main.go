package main

import "fmt"

// go run .\main.go

func main() {

	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}

		fmt.Println("perulangan ke-", i)
	}

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue //skip ke i+1
		}

		fmt.Println("perulangan ke-", i)
	}
}
