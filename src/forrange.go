package main

import "fmt"

func main() {
	// START OMIT
	const msg = "Hello, 世界"
	for i, r := range msg { // HL
		fmt.Printf("index %d: run %q\n", i, r)
	}
	// END OMIT
}
