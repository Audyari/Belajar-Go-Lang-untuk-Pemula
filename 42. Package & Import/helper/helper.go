package helper

import "fmt"

func SayHello(name string) string {
	fmt.Println("Fungsi dari paket helper dipanggil!")
	return "Hello World " + name
}
