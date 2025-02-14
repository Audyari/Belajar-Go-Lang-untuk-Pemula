package main

import "fmt"

// go run .\main.go
func main() {

	type NoKtp string

	var KtpAudi NoKtp = "1234567890"

	fmt.Println(KtpAudi)

}
