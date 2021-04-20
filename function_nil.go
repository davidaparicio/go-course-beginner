package main

import (
	"fmt"
	"strings"
)

// START OMIT
func main() {
	var greetings func(string) string
	fmt.Printf("%T\n", greetings)
	// fmt.Println(greetings == greetings) // try uncommenting this
	fmt.Println(greetings == nil)
	// greetings("Julien") // try uncommenting this

	greetings = strings.ToUpper
	fmt.Println(greetings("Julien!"))
}

// END OMIT
