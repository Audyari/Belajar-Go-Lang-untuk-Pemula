package main

import "fmt"

// go run .\main.go

func NewMap(name string, age string, email string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name":  name,
			"age":   age,
			"email": email,
		}
	}
}

func contoh(name string) string {
	if name == "" {
		return "nama kosong"
	} else {
		return "nama ada"
	}
}

func main() {
	person := NewMap("Audyari", "21", "wiyono@gmail.com")
	fmt.Println(person)

	person = NewMap("Wiyono", "22", "wiyono@gmail.com")
	fmt.Println(person)

	person = NewMap("", "", "")
	fmt.Println(person)
}
