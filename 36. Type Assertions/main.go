package main

import "fmt"

// go run .\main.go

func random() interface{} {
	return "Audyari"
}

func main() {
	result := random()
	fmt.Println(result)

	resultString := result.(string)
	fmt.Println(resultString)

	// resultInt := result.(int)
	// fmt.Println(resultInt)

	switch result.(type) {
	case string:
		fmt.Println("result adalah string" + resultString)
	case int:
		fmt.Println("result adalah int")
	default:
		fmt.Println("result adalah tipe yang tidak dikenal")
	}

}
