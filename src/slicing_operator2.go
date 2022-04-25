package main

import "fmt"

func main() {
	// START OMIT
	names := []string{"foo", "bar", "baz", "qux", "quux"}
	fmt.Println(names[:4]) // HL
	fmt.Println(names[2:]) // HL
	// END OMIT
}
