package helper

import "fmt"

var version = "1.0.0" // huruf kecil gk bisa di akses
var Aplication = "golang"

func SayHello(name string) string {
	fmt.Println("Fungsi dari paket helper dipanggil!")
	return "Hello World " + name

}

func sayGoodbye(name string) string { // huruf keci gk bisa di akses
	return "Goodbye " + name
}
