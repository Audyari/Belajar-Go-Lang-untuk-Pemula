package main

import "fmt"

// go run .\main.go

func factorialLoop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i
	}
	return result
}

func factorialRecursive(value int) int {
	if value == 0 {
		return 1
	}
	return value * factorialRecursive(value-1)
}

func main() {

	hasil := 10 * 9 * 8 * 7 * 6 * 5 * 4 * 3 * 2 * 1

	fmt.Println(factorialLoop(10))

	fmt.Println(factorialRecursive(10))

	fmt.Println(hasil)
}
