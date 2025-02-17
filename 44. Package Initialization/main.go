package main

import (
	"belajar/database"
	"fmt"
)

// go run .\main.go

func main() {
	fmt.Println(database.GetConnection())
}
