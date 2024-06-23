package main

import "fmt"

func main() {
	// Contoh 1: if dasar
	age := 18
	if age >= 18 {
		fmt.Println("Anda sudah dewasa")
	}

	// Contoh 2: if-else
	score := 85
	if score >= 90 {
		fmt.Println("Anda mendapat nilai A")
	} else {
		fmt.Println("Nilai Anda di bawah 90")
	}

	// Contoh 3: if-else if-else
	temperature := 25
	if temperature < 0 {
		fmt.Println("Cuaca sangat dingin")
	} else if temperature < 20 {
		fmt.Println("Cuaca dingin")
	} else if temperature < 30 {
		fmt.Println("Cuaca nyaman")
	} else {
		fmt.Println("Cuaca panas")
	}

	// Contoh 4: inisialisasi variabel dalam if
	if num := 7; num > 0 {
		fmt.Println("Nilai", num, "adalah positif")
	} else {
		fmt.Println("Nilai", num, "adalah negatif atau nol")
	}

	// Contoh 5: kombinasi dengan operator logika
	x, y := 5, 10
	if x > 0 && y > 0 {
		fmt.Println("Kedua nilai positif")
	}

	if x == 5 || y == 5 {
		fmt.Println("Salah satu nilai sama dengan 5")
	}

	// Contoh 6: switch
	day := 3
	switch day {
	case 1:
		fmt.Println("Senin")
	case 2:
		fmt.Println("Selasa")
	case 3:
		fmt.Println("Rabu")
	default:
		fmt.Println("Hari yang tidak valid")
	}
}
