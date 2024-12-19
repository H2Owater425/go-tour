package main

import "fmt"

func main() {
	var s []int

	printSlice(s)

	// Works on nil slice
	s = append(s, 0)

	printSlice(s)

	// Grows as needed
	s = append(s, 1)

	printSlice(s)

	// Can add many elements
	s = append(s, 2, 3, 4)

	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v %p\n", len(s), cap(s), s, &s)
}
