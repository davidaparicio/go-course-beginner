package main

import "fmt"

func main() {
	// START OMIT
	names := []string{"foo", "bar", "baz", "qux", "quux"}
	s := names[2:4] // HL
	fmt.Printf("%T\n", s)
	fmt.Println(s)
	fmt.Println(names[:4]) // HL
	fmt.Println(names[2:]) // HL
	// END OMIT
}
