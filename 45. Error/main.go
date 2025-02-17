package main

import (
	"errors"
	"fmt"
)

// go run .\main.go

func Pembagian(nilai int, divisor int) (int, error) {
	if divisor == 0 {
		return 0, errors.New("divisor tidak boleh 0")
	}
	return nilai / divisor, nil
}

func main() {
	hasil, err := Pembagian(0, 10)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Hasil", hasil)
	}
}
