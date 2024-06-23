package main

import "fmt"

func main() {
	// Penjumlahan
	x := 5 + 3
	fmt.Println("Hasil penjumlahan:", x) // Output: Hasil penjumlahan: 8
	fmt.Println("===========================")

	// Pengurangan
	y := 10 - 4
	fmt.Println("Hasil pengurangan:", y) // Output: Hasil pengurangan: 6
	fmt.Println("===========================")

	// Perkalian
	z := 7 * 6
	fmt.Println("Hasil perkalian:", z) // Output: Hasil perkalian: 42
	fmt.Println("===========================")

	// Pembagian
	a := 15 / 3
	fmt.Println("Hasil pembagian:", a) // Output: Hasil pembagian: 5
	fmt.Println("===========================")

	// Modulo
	b := 17 % 5
	fmt.Println("Sisa bagi:", b) // Output: Sisa bagi: 2
	fmt.Println("===========================")

	// Operator Unary
	c := 4
	c++                                // c = 5
	fmt.Println("Hasil increment:", c) // Output: Hasil increment: 5

	d := 7
	d--                                // d = 6
	fmt.Println("Hasil decrement:", d) // Output: Hasil decrement: 6
	fmt.Println("===========================")

	// Operator Gabungan
	x = 10
	x += 5                         // x = 15
	x -= 3                         // x = 12
	x *= 2                         // x = 24
	x /= 4                         // x = 6
	x %= 3                         // x = 0
	fmt.Println("Hasil akhir:", x) // Output: Hasil akhir: 0
	fmt.Println("===========================")
}
