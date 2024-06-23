package main

import "fmt"

func main() {
	// Deklarasi dan inisialisasi array
	numbers := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Array angka:", numbers)

	// Mengakses elemen array
	fmt.Println("Elemen pertama:", numbers[0])
	fmt.Println("Elemen ketiga:", numbers[2])

	// Mengubah nilai elemen array
	numbers[1] = 10
	fmt.Println("Array angka setelah diubah:", numbers)

	// Menghitung panjang array
	fmt.Println("Panjang array:", len(numbers))

	// Iterasi menggunakan for loop
	fmt.Println("Iterasi menggunakan for loop:")
	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}

	// Iterasi menggunakan for-range
	fmt.Println("\nIterasi menggunakan for-range:")
	for index, value := range numbers {
		fmt.Printf("Index: %d, Nilai: %d\n", index, value)
	}

	// Array multidimensi
	matrix := [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println("\nArray multidimensi (Matrix):")
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Print(matrix[i][j], " ")
		}
		fmt.Println()
	}
}
