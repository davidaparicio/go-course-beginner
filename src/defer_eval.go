package main

import (
	"fmt"
)

// START OMIT
func main() {
	var n int
	defer func(i int) { fmt.Println("defer with param:", i) }(n) // HL
	n++
	defer func() { fmt.Println("defer without param:", n) }() // HL
	n++
	fmt.Println("normal print", n)
}

// END OMIT
