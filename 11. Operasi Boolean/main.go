package main

import "fmt"

// go run .\main.go

func main() {

	var nilaiakhir int = 98
	var nilaiabsen int = 81

	var lulusnilaiakhir bool = nilaiakhir > 80
	var lulusnilaiabsen bool = nilaiabsen > 80

	var lulus bool = lulusnilaiakhir && lulusnilaiabsen

	fmt.Println(lulus)

	if lulus {
		fmt.Println("Selamat Anda Lulus")
	} else {
		fmt.Println("Maaf Anda Tidak Lulus")
	}

}
