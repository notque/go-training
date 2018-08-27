package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	b := make([]int, 0, 15)

	fmt.Println(a)
	fmt.Println(b)
}