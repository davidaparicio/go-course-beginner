package main

import "fmt"

func main() {
	// START OMIT
	var s []int
	for i := 0; i < 1024; i++ {
		fmt.Printf("len: %4d; cap: %4d\n", len(s), cap(s))
		s = append(s, i) // HL
	}
	// END OMIT
}
