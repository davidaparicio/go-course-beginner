package main

import (
	"fmt"
)

// START OMIT
func main() {
	fmt.Println(MinAndMax(99, 1_000_000, -32))
}

func MinAndMax(first int, rest ...int) (min, max int) { // HL
	min, max = first, first
	for _, i := range rest {
		if i < min {
			min = i
		} // note: don't format your code like this
		if i > max {
			max = i
		}
	}
	return
}

// END OMIT
