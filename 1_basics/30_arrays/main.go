package main

import "fmt"

func main() {
	var a [2]string

	fmt.Printf("%p %p\n", &a[0], &a[1])

	fmt.Printf("%q %q\n", a[0], a[1])

	a[0] = "Hell"
	a[1] = "World"

	fmt.Println(a[0], a[1])

	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}

	fmt.Println(primes)
}
