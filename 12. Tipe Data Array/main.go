package main

import "fmt"

// go run .\main.go

func main() {

	var nama [2]string
	nama[0] = "Audyari"
	nama[1] = "Wiyono"

	fmt.Println(nama)
	fmt.Println(nama[0])
	fmt.Println(nama[1])

	var value = [3]int{
		1, 2, 3,
	}

	fmt.Println(value)

	var data = [...]int{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
	}

	fmt.Println(data)
	fmt.Println(len(data))
	fmt.Println(data[0])
	data[0] = 100
	fmt.Println(data)

}
