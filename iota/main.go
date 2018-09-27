package main

import (
	"fmt"
)

func main() {
	type Weekday int

	const (
		Sunday Weekday = iota // this sets it to 0 and increments automatically for remaining values
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
	)

	fmt.Println(Sunday)

	// setup some const values with iota where each successive const value is a power of 2
	type Flags uint

	const (
		a Flags = 1 << iota
		b
		c
		d
		e
	)

	fmt.Printf("%d, %d", a, d)
}
