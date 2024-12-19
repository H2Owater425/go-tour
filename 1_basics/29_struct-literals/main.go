package main

import "fmt"

type Vertex struct {
	X int64
	Y int64
}

var (
	// Vertex
	v1 = Vertex{1, 2}

	// implicit Y:0
	v2 = Vertex{X: 1}

	// X:0 and Y:0
	v3 = Vertex{}

	// *Vertex
	p = &Vertex{1, 2}
)

func main() {
	fmt.Println(v1, p, v2, v3)
}
