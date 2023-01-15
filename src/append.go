package main

import "fmt"

func main() {
	// START OMIT
	names := []string{"foo"}
	names = append(names, "bar", "baz") // HL
	moreNames := []string{"qux", "quux"}
	fmt.Println(names)
	names = append(names, moreNames...) // HL
	fmt.Println(names)
	fmt.Println(moreNames)
	// END OMIT
}
