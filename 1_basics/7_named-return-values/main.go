package main

import "fmt"

// return 변수와 타입 선언을 동시에
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x

	return
}

func main() {
	fmt.Println(split(17))
}
