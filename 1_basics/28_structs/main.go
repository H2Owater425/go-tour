package main

import "fmt"

type Vertex struct {
	X int64
	Y int64
}

func main() {
	fmt.Println(Vertex{1, 2})

	v := Vertex{1, 2}

	v.X = 4

	fmt.Println(v)

	p := &v

	p.X = 1e9

	fmt.Println(v)
}
