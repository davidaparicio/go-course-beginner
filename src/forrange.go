package main

import "fmt"

func main() {
	// START OMIT
	numbers := []int{4, 8, 15, 16, 23, 42}
	for i, n := range numbers { // HL
		fmt.Printf("index %d: %d\n", i, n)
	}
	// END OMIT
}
