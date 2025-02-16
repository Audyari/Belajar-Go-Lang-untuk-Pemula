package main

import "fmt"

// go run .\main.go

func logging() {
	fmt.Println("Sedang menulis log")
}

func runAplication() {
	defer logging()
	fmt.Println("Aplikasi selesai dijalankan")

}

func endApp() {
	fmt.Println("Aplikasi selesai")
	messages := recover()
	fmt.Println("terjadi eror : ", messages)
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("ERROR disebabkan panic")
	} else {
		fmt.Println("Aplikasi berjalan")
	}
}

func main() {
	runAplication()

	fmt.Println("---------------")

	runApp(true)
	fmt.Println("---------------")
	runApp(false)
}
