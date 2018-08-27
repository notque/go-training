package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := "10"
	b, err := strconv.Atoi(a)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
	fmt.Printf("Integer: %d", b)
}