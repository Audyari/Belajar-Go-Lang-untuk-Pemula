package main

import "fmt"

func main() {
    // Slicing
    mySlice := []int{1, 2, 3, 4, 5}
    subSlice := mySlice[1:4]
    fmt.Println("Slicing:", subSlice)

    // Append
    mySlice = append(mySlice, 6, 7)
    fmt.Println("Append:", mySlice)

    // Copy
    src := []int{1, 2, 3, 4, 5}
    dst := make([]int, len(src))
    copy(dst, src)
    fmt.Println("Copy:", dst)

    // Insert
    mySlice = []int{1, 2, 4, 5}
    mySlice = append(mySlice[:3], append([]int{3}, mySlice[3:]...)...)
    fmt.Println("Insert:", mySlice)

    // Remove
    mySlice = []int{1, 2, 3, 4, 5}
    mySlice = append(mySlice[:2], mySlice[3:]...)
    fmt.Println("Remove:", mySlice)

    // Resize
    mySlice = []int{1, 2, 3}
    mySlice = mySlice[:2]
    fmt.Println("Resize (shrink):", mySlice)
    mySlice = mySlice[1:]
    fmt.Println("Resize (expand):", mySlice)

    // Concatenate
    slice1 := []int{1, 2, 3}
    slice2 := []int{4, 5, 6}
    combined := append(slice1, slice2...)
    fmt.Println("Concatenate:", combined)

    // Iterate
    mySlice = []int{1, 2, 3, 4, 5}
    for i, v := range mySlice {
        fmt.Printf("Index: %d, Value: %d\n", i, v)
    }
}