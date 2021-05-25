package main

import "fmt"

// START OMIT
var i = -42

func main() {
	s := "Hello, 世界"
	fmt.Println("length:", len(s))

	fmt.Print("bytes: ")
	for i := 0; i < len(s); i++ {
		fmt.Print(s[i], " ")
	}
	fmt.Println()

	// indexing a string out of bounds causes a panic
	// fmt.Println(s[i])
	// fmt.Println(s[len(s)])
}

// END OMIT
