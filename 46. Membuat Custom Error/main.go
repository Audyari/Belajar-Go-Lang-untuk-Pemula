package main

import "fmt"

// go run .\main.go

type validationErorr struct {
	message string
}

func (v *validationErorr) Error() string {
	return v.message
}

type notFoundError struct {
	message string
}

func (n *notFoundError) Error() string {
	return n.message
}

func SaveData(id string, data any) error {
	if id == "" {
		return &validationErorr{"ID Tidak Boleh Kosong"}
	}

	if id != "Audy" {
		return &notFoundError{"Data Tidak Ditemukan"}
	}

	return nil

}

func main() {
	err := SaveData("Audy", nil)

	// fmt.Println(err)

	// if err != nil {
	// 	if validationError, ok := err.(*validationErorr); ok {
	// 		fmt.Println("validation error : " + validationError.Error())
	// 	} else if notFoundError, ok := err.(*notFoundError); ok {
	// 		fmt.Println("not found error : " + notFoundError.Error())
	// 	} else {
	// 		fmt.Println("unknown error : " + err.Error())
	// 	}
	// }

	if err != nil {

		switch err.(type) {
		case *validationErorr:
			fmt.Println("validation error : " + err.Error())
		case *notFoundError:
			fmt.Println("not found error : " + err.Error())
		default:
			fmt.Println("unknown error : " + err.Error())
		}
	} else {
		fmt.Println("Data Berhasil Disimpan")
	}

}
