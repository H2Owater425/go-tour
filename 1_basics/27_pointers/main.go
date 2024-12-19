package main

import "fmt"

func main() {
	i, j := 42, 2701

	// Pointer to i
	p := &i

	// Read i through pointer
	fmt.Println(*p)

	// Set i through pointer
	*p = 21

	// Check new i
	fmt.Println(i)

	// Pointer to j
	p = &j

	// Devide j through pointer
	*p /= 37

	// Check new j
	fmt.Println(j)
}
