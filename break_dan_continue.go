package main

import "fmt"

func main() {
	// Contoh penggunaan break
	fmt.Println("Contoh penggunaan break:")
	for i := 0; i < 5; i++ {
		if i == 3 {
			break
		}
		fmt.Println(i)
	}
	fmt.Println()

	// Contoh penggunaan continue
	fmt.Println("Contoh penggunaan continue:")
	for i := 0; i < 5; i++ {
		if i == 3 {
			continue
		}
		fmt.Println(i)
	}
	fmt.Println()

	// Contoh penggunaan break dan continue pada loop bersarang
	fmt.Println("Contoh penggunaan break dan continue pada loop bersarang:")
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 1 && j == 1 {
				break
			}
			if i == 2 && j == 1 {
				continue
			}
			fmt.Printf("i: %d, j: %d\n", i, j)
		}
	}
}
