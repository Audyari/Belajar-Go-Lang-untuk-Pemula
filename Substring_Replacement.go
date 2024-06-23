package main

import (
	"fmt"
	"strings"
)

func main() {
	pesan := "Saya suka bahasa pemrograman Go."
	pesanBaru := strings.Replace(pesan, "Go", "Rust", 1)
	fmt.Println(pesanBaru) // Output: Saya suka bahasa pemrograman Rust.
}
