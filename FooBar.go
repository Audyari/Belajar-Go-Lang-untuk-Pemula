package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FooBar")
		} else if i%3 == 0 {
			fmt.Println("Foo")
		} else if i%5 == 0 {
			fmt.Println("Bar")
		} else {
			fmt.Println(i)
		}
	}
}
