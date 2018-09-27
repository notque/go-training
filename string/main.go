package main

import (
	"fmt"
	"strings"
)

// count the number of times a pattern appears in text.
func main() {
	stuff := "Do you know that feeling when The Onion is eerily right? This is one of those times. People might not be physically shuddering at your docs, but there’s a good chance they are doing it mentally. I struggled with the idea that people aren’t going to read what I write, unless I present it in the most easily digestible way. Heck, this could even happen for this blog post. "
	if strings.Contains(stuff, "Onion") {
		fmt.Printf("It does\n")
	}
	for index, s := range strings.Fields(stuff) {
		if index == 1 {
			fmt.Printf("String contains:\n")
		}
		count := wordcount(stuff, s)
		fmt.Printf("%s %d times\n", s, count)
	}

	fmt.Printf(stuff[0:6])
}

func wordcount(source, word string) (count int) {
	count = strings.Count(source, word)
	return
}
