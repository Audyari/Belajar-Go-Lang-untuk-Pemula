package main

import "fmt"

// go run .\main.go

func getFullName() (string, string) {
	return "Audyari", "Wiyono"
}

func main() {

	firstName, lastName := getFullName()
	fmt.Println(firstName, lastName)

	data1, _ := getFullName()
	fmt.Println(data1)

}
