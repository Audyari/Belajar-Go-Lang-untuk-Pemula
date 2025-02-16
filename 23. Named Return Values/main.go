package main

// go run .\main.go

import "fmt"

func getComplateName() (firstName string, middleName string, lastName string) {
	firstName = "Audyari"
	middleName = "Wiyono"
	lastName = "Wiyonoterakhir"
	return firstName, middleName, lastName
}

func main() {
	firstName, middleName, lastName := getComplateName()
	fmt.Println(firstName, middleName, lastName)
}
