package main

import "fmt"

// go run .\main.go

func main() {

	counter := 1

	for counter <= 10 {
		fmt.Println("counter = ", counter)
		counter++
	}

	for counter := 1; counter <= 10; counter++ {
		fmt.Println("counter = ", counter)
	}

	name := []string{"Audyari", "Wiyono", "budi", "wati", "caca", "dadi"}

	for i := 0; i < len(name); i++ {
		fmt.Println(name[i])
	}

	student := make(map[string]int)

	student["Audyari"] = 90
	student["Wiyono"] = 80
	student["budi"] = 70

	for key, value := range student {
		fmt.Println("nama = ", key, "nilai = ", value)
	}

}
