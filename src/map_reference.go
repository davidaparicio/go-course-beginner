package main

import "fmt"

// START OMIT
func main() {
	var m map[int]string
	fmt.Println(m == nil)
	// fmt.Println(m == m) // causes a compilation error
}

// END OMIT
