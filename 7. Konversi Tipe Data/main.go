package main

import "fmt"

// go run .\main.go
func main() {

	var nilai32 int32 = 32768
	var nilai64 int64 = int64(nilai32)
	var nilai16 int16 = int16(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai16)

	var name = "Audyari"
	var data byte = name[0]
	var conversi string = string(name[0])

	fmt.Println(name)
	fmt.Println(data)
	fmt.Println(conversi)
}
