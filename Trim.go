package main

import (
	"fmt"
	"strings"
)

func main() {
	input := "   Halo, Dunia!   "
	dipangkas := strings.Trim(input, " ")
	fmt.Println(dipangkas) // Output: Halo, Dunia!
}
