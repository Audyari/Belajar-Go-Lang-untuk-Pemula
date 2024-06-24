package main

import "fmt"

func main() {
    // Deklarasi dan inisialisasi slice
    mySlice := []int{1, 2, 3, 4, 5}
    fmt.Println("Slice awal:", mySlice)

    // Akses elemen slice
    fmt.Println("Elemen ke-3:", mySlice[2])

    // Pemotongan slice
    subSlice := mySlice[1:4]
    fmt.Println("Sub-slice:", subSlice)

    // Penambahan elemen baru
    mySlice = append(mySlice, 6, 7)
    fmt.Println("Slice setelah ditambah:", mySlice)

    // Menampilkan panjang dan kapasitas slice
    fmt.Println("Panjang slice:", len(mySlice))
    fmt.Println("Kapasitas slice:", cap(mySlice))
}