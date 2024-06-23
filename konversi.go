package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Konversi dari int ke float64
	var x int = 42
	y := float64(x)
	fmt.Println("Int ke Float64:", y) // Output: Int ke Float64: 42.0

	// Konversi dari float64 ke int
	var a float64 = 3.14
	b := int(a)
	fmt.Println("Float64 ke Int:", b) // Output: Float64 ke Int: 3

	// Konversi dari string ke int
	var s string = "42"
	c, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("Terjadi error:", err)
	} else {
		fmt.Println("String ke Int:", c) // Output: String ke Int: 42
	}

	// Konversi dari int ke string
	var d int = 42
	e := strconv.Itoa(d)
	fmt.Println("Int ke String:", e) // Output: Int ke String: 42
}
