package main

import "fmt"

func main() {
	x := 10
	y := 5

	switch {
	case x > y:
		fmt.Println("x is greater than y")
	case x < y:
		fmt.Println("x is less than y")
	default:
		fmt.Println("x and y are equal")
	}
}
