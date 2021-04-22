package main

import (
	"fmt"
)

// START OMIT
func main() {
	const r = 'Z'
	if 'A' <= r && r <= 'Z' || 'a' <= r && r <= 'z' { // HL
		fmt.Printf("%q is alphabetical ASCII\n", r)
	}
}

// END OMIT
