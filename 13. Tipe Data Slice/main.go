package main

import "fmt"

// go run .\main.go

func main() {
	name := [...]string{"Audyari", "Wiyono", "budi", "wati", "caca", "dadi"}
	silce := name[4:6]
	fmt.Println(name)
	fmt.Println(silce)

	fmt.Println(silce[0])
	silce[0] = "keren"
	fmt.Println(name)

	fmt.Println("==============")

	days := []string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}
	fmt.Println(days)

	dayslice1 := days[5:]
	dayslice1[0] = "Sabtu Baru"
	dayslice1[1] = "Minggu Baru"

	fmt.Println(dayslice1)
	fmt.Println(days)

	fmt.Println("==============")

	dayslice2 := append(dayslice1, "Bukan Hari Libur")

	fmt.Println(dayslice2)
	fmt.Println(days)

	fmt.Println("==============")

	newSlice := make([]string, 2, 5)
	newSlice[0] = "Senin"
	newSlice[1] = "Selasa"

	appendSlice := append(newSlice, "Rabu")

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))
	fmt.Println(appendSlice)
	fmt.Println(len(appendSlice))
	fmt.Println(cap(appendSlice))

	fmt.Println("==============")

	fromslice := name[:]

	toslice := make([]string, len(fromslice), cap(fromslice))

	copy(toslice, fromslice)

	fmt.Println(fromslice)
	fmt.Println(toslice)

	fmt.Println("==============")

	iniArray := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	inisclice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println(iniArray)
	fmt.Println(inisclice)

}
