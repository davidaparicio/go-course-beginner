package main

import "fmt"

func main() {
	// START OMIT
	names := []string{"foo", "bar", "baz", "qux", "quux"}
	firstThree := names[0:3] // HL
	c := cap(firstThree)
	fmt.Printf("capacity of firstThree: %d\n", c)
	names2 := firstThree[:c] // HL
	fmt.Printf("names2: %v\n", names2)
	// END OMIT
}
