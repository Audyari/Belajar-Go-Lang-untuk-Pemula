package main

import "fmt"

// go run .\main.go

func sumALL(numbers ...int) int {
	total := 0
	for _, value := range numbers {
		total += value
	}
	return total
}

func main() {

	result := sumALL(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(result)

	number := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	result = sumALL(number...)
	fmt.Println(result)

}
