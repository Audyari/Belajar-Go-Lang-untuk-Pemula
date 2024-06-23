package main

import "fmt"

func main() {
	// Contoh Increment
	x := 5
	fmt.Println("Nilai awal x:", x)

	x++ // Increment dengan postfix
	fmt.Println("Nilai x setelah x++:", x)

	x++ // Increment dengan postfix
	fmt.Println("Nilai x setelah ++x:", x)

	// Contoh Decrement
	y := 10
	fmt.Println("Nilai awal y:", y)

	y-- // Decrement dengan postfix
	fmt.Println("Nilai y setelah y--:", y)

	y-- // Decrement dengan postfix
	fmt.Println("Nilai y setelah --y:", y)
}
