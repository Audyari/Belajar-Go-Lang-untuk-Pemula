package main

import "fmt"

func main() {
    // Deklarasi dan inisialisasi map
    myMap := make(map[string]int)

    // Menambahkan entri ke map
    myMap["apple"] = 5
    myMap["banana"] = 3
    myMap["orange"] = 7

    // Mengubah entri dalam map
    myMap["apple"] = 10

    // Menghapus entri dari map
    delete(myMap, "banana")

    // Mengakses nilai dalam map
    value, exists := myMap["apple"]
    if exists {
        fmt.Println("Value of 'apple':", value)
    } else {
        fmt.Println("Key 'apple' not found")
    }

    // Iterasi atas map
    fmt.Println("\nIterating over the map:")
    for key, value := range myMap {
        fmt.Printf("Key: %s, Value: %d\n", key, value)
    }

    // Pengecekan keberadaan kunci dalam map
    _, exists = myMap["banana"]
    if exists {
        fmt.Println("\nKey 'banana' exists")
    } else {
        fmt.Println("\nKey 'banana' not found")
    }

    // Inisialisasi map dengan data awal
    anotherMap := map[string]int{
        "apple":  5,
        "banana": 3,
        "orange": 7,
    }
    fmt.Println("\nAnother map:", anotherMap)

    // Menggunakan map sebagai parameter fungsi
    printMap(anotherMap)
}

func printMap(m map[string]int) {
    fmt.Println("\nPrinting map from function:")
    for key, value := range m {
        fmt.Printf("Key: %s, Value: %d\n", key, value)
    }
}