package main

import "fmt"

func main() {
	// Bentuk Dasar
	fmt.Println("Bentuk Dasar:")
	for i := 0; i < 5; i++ {
		fmt.Println(i) // Output: 0 1 2 3 4
	}
	fmt.Println()

	// Bentuk Singkat
	fmt.Println("Bentuk Singkat:")
	x := 10
	for x > 0 {
		fmt.Println(x) // Output: 10 9 8 7 6 5 4 3 2 1
		x--
	}
	fmt.Println()

	// Bentuk for-each
	fmt.Println("Bentuk for-each:")
	fruits := []string{"apple", "banana", "cherry"}
	for i, fruit := range fruits {
		fmt.Println(i, fruit)
		// Output:
		// 0 apple
		// 1 banana
		// 2 cherry
	}
	fmt.Println()

	// Bentuk Tak Terbatas
	fmt.Println("Bentuk Tak Terbatas:")
	count := 0
	for {
		fmt.Println("Looping...", count)
		count++
		if count > 2 {
			break // Menghentikan loop
		}
	}
}
