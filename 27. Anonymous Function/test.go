package main

import "fmt"

// go run .\test.go
type BlackList1 func(string) bool

func registerUser1(name string, blackList BlackList1) {
	if blackList(name) {
		fmt.Println("Maaf", name, "tidak bisa login")
		return
	}
	fmt.Println("Selamat datang", name)
}

func blackListUser(name string) bool {

	listOfBlacklist := []string{"anjing", "bebek", "kucing"}

	for _, blacklist := range listOfBlacklist {
		if name == blacklist {
			return true
		}
	}

	return false
}

// func main() {

// 	registerUser1("Audyari", blackListUser)
// 	registerUser1("anjing", blackListUser)
// 	registerUser1("kucing", blackListUser)
// 	registerUser1("bebek", blackListUser)
// }
