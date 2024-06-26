package main

import "fmt"

func main() {
	n := 10
	fib := fibonacci(n)
	fmt.Println(fib)
}

func fibonacci(n int) []int {
	fib := make([]int, n)
	fib[0] = 0
	fib[1] = 1
	for i := 2; i < n; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}
	return fib
}
