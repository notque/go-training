package main

import (
	"fmt"
)

func main() {
	var m map[string]string
	m = make(map[string]string)

	m["Nathan"] = "Arizona"
	m["Bill"] = "Denmark"

	fmt.Println("The value:", m["Nathan"])
	elem, ok := m["Natha"]
	if !ok {
		fmt.Printf("Element: %s does not exist\n", elem)
	}

	delete(m, "Bill")

	newmap := map[string]int{}
	fmt.Println(newmap)
}