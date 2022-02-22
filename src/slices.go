package main

import "fmt"

func main() {
	// START OMIT
	roles := []string{ // HL
		"Viewer",
		"Editor",
		"Admin", // this final comma is mandatory // HL
	}
	fmt.Printf("type name: %T\n", roles)
	s := make([]string, 10) // HL
	fmt.Println(s)
	// END OMIT
}
