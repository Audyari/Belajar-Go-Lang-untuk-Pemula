package main

import (
	"fmt"
	"strings"
)

func main() {
	kalimat := "Kucing cepat mengejar tikus."
	kata := strings.Split(kalimat, " ")
	fmt.Println(kata) // Output: [Kucing cepat mengejar tikus.]

	kalimatGabung := strings.Join(kata, ", ")
	fmt.Println(kalimatGabung) // Output: Kucing, cepat, mengejar, tikus.
}
