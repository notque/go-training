package main

import (
	"fmt"
)

func main() {
	a := varadic(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Printf("Output: %d\n", a)
}

func varadic(values ...int) (int) {
	var total int
	for i := range values {
		total = total + i
	}
	return total
}