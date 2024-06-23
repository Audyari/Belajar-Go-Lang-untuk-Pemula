package main

import "fmt"

func main() {
	x := 10
	y := 5

	// Equal (==)
	fmt.Println(x == y) // Output: false

	// Not Equal (!=)
	fmt.Println(x != y) // Output: true

	// Less Than (<)
	fmt.Println(x < y) // Output: false

	// Less Than or Equal To (<=)
	fmt.Println(x <= y) // Output: false

	// Greater Than (>)
	fmt.Println(x > y) // Output: true

	// Greater Than or Equal To (>=)
	fmt.Println(x >= y) // Output: true

	// Contoh penggunaan dalam if-else
	if x > y {
		fmt.Println("x lebih besar dari y")
	} else {
		fmt.Println("x tidak lebih besar dari y")
	}

	// Contoh penggunaan dalam for loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
}
