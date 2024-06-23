package main

import "fmt"

func main() {
	fruit := "apple"

	switch fruit {
	case "apple", "banana", "orange":
		fmt.Println("This is a fruit.")
	case "carrot", "potato":
		fmt.Println("This is a vegetable.")
	default:
		fmt.Println("I don't know what this is.")
	}
}
