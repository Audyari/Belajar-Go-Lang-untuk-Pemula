package main

import "fmt"

func main() {
    // Map luar menyimpan data mahasiswa berdasarkan NIM
    students := map[string]map[string]interface{}{
        "2001001": {
            "name":  "John Doe",
            "grades": map[string]float64{
                "Algoritma":   90.5,
                "Pemrograman": 85.0,
                "Matematika":  92.0,
            },
        },
        "2001002": {
            "name":  "Jane Smith",
            "grades": map[string]float64{
                "Algoritma":   87.0,
                "Pemrograman": 91.5,
                "Matematika":  88.0,
            },
        },
        "2001003": {
            "name":  "Bob Johnson",
            "grades": map[string]float64{
                "Algoritma":   78.0,
                "Pemrograman": 83.5,
                "Matematika":  80.0,
            },
        },
    }

    // Mengakses data mahasiswa
    fmt.Println("Data Mahasiswa:")
    for nim, student := range students {
        fmt.Println("NIM:", nim)
        fmt.Println("Nama:", student["name"])
        fmt.Println("Nilai:")
        grades := student["grades"].(map[string]float64)
        for course, grade := range grades {
            fmt.Printf("- %s: %.2f\n", course, grade)
        }
        fmt.Println()
    }
}