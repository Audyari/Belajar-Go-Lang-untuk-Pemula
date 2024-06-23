package main

import "fmt"

func main() {
	// Bilangan Bulat (Integers)
	var usia int = 30
	var jumlah uint8 = 255
	fmt.Println("Usia:", usia)
	fmt.Println("Jumlah:", jumlah)
	fmt.Println("=================")

	// Bilangan Floating-Point
	var pi float64 = 3.14159
	var jarak float32 = 5.2
	fmt.Println("Pi:", pi)
	fmt.Println("Jarak:", jarak)
	fmt.Println("=================")

	// Boolean
	var adalahMahasiswa bool = true
	var sudahMenikah bool = false
	fmt.Println("Apakah Mahasiswa:", adalahMahasiswa)
	fmt.Println("Sudah Menikah:", sudahMenikah)
	fmt.Println("==============================")

	// String
	var nama string = "Audyari W"
	var pesan string = "Halo, Dunia!"
	fmt.Println("Nama:", nama)
	fmt.Println("Pesan:", pesan)
	fmt.Println("==============================")

	// Kompleks
	var kompleks1 complex64 = 2 + 3i
	var kompleks2 complex128 = 4 + 5i
	fmt.Println("Kompleks1:", kompleks1)
	fmt.Println("Kompleks2:", kompleks2)
	fmt.Println("==============================")

	// Array
	var angka [5]int = [5]int{1, 2, 3, 4, 5}
	var warna [3]string = [3]string{"merah", "hijau", "biru"}
	fmt.Println("Angka:", angka)
	fmt.Println("Warna:", warna)
	fmt.Println("==============================")

	// Slice
	var buah []string = []string{"apel", "pisang", "jeruk"}
	var data []int = []int{10, 20, 30, 40, 50}
	fmt.Println("Buah:", buah)
	fmt.Println("Data:", data)
	fmt.Println("==============================")

	// Map
	var orang map[string]interface{} = map[string]interface{}{
		"nama":      "Audyari W",
		"usia":      30,
		"adalahVIP": true,
	}
	fmt.Println("Orang:", orang)
	fmt.Println("==============================")

	// Struct
	type Orang struct {
		Nama string
		Usia int
	}

	var john Orang = Orang{
		Nama: "Audyari W",
		Usia: 30,
	}
	fmt.Println("Audy:", john)
	fmt.Println("==============================")
}
