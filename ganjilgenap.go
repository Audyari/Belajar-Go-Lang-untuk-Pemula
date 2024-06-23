package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Menerima input dari pengguna
	fmt.Print("Masukkan sebuah bilangan: ")
	var input string
	fmt.Scanln(&input)

	// Mengubah input menjadi integer
	num, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Input tidak valid:", err)
		return
	}

	// Memeriksa apakah bilangan ganjil atau genap
	if num%2 == 0 {
		fmt.Println(num, "adalah bilangan genap")
	} else {
		fmt.Println(num, "adalah bilangan ganjil")
	}

	// Menampilkan bilangan ganjil dalam rentang 1 sampai input
	fmt.Println("\nBilangan Ganjil:")
	for i := 1; i <= num; i++ {
		if i%2 != 0 {
			fmt.Println(i)
		}
	}

	// Menampilkan bilangan genap dalam rentang 1 sampai input
	fmt.Println("\nBilangan Genap:")
	for j := 1; j <= num; j++ {
		if j%2 == 0 {
			fmt.Println(j)
		}
	}
}
