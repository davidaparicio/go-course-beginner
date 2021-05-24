package main

import (
	"fmt"
)

// START OMIT
func makeIncrement() func() int {
	var i int // HL
	return func() int {
		i++ // HL
		return i
	}
}

func main() {
	incr := makeIncrement()
	fmt.Println(incr())
	fmt.Println(incr())
}

// END OMIT
