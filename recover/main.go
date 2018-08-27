package main

import (
	"fmt"
)

func div(a, b int) (result int, err error) {
	defer func() {
	  if p := recover(); p != nil {
		err = fmt.Errorf("Internal Error: %d / %d = %v", a, b, p)
	  }
	}()
	// some division operation that returns result
	result = a / b
	return
  }

func main() {
	result, err := div(10, 0)
	if err != nil {
		fmt.Printf("The error was: %s\n", err)
	}
	fmt.Printf("The result was: %d\n", result)
}