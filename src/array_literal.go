package main

import "fmt"

func main() {
	// START OMIT
	roles := [3]string{ // HL
		"Viewer",
		"Editor",
		"Admin", // this final comma is mandatory // HL
	}
	fmt.Printf("%T\n", roles)
	// END OMIT
}
