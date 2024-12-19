package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	for guess := 1.0; ; {
		state := guess
		guess -= (guess*guess - x) / (2 * guess)

		if float32(state) == float32(guess) {
			return guess
		}
	}
}

func main() {
	fmt.Println(Sqrt(math.Pow(999999, 2)))

	fmt.Println("printing")

	fmt.Println(math.Sqrt(math.Pow(999999, 2)))
}
