package main

import "fmt"

func main() {
	// START OMIT
	var s []int
	for i := 0; i < 2048; i++ {
		fmt.Printf("len: %3d; cap: %3d\n", len(s), cap(s))
		s = append(s, i)
	}
	// END OMIT
}
