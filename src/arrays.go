package main

import "fmt"

func main() {
	// START OMIT
	userIds := [3]int{42, 56, 70} // HL
	fmt.Printf("type name: %T\n", userIds)
	l := len(userIds) // HL
	fmt.Printf("length: %d\n", l)
	roles := [...]string{ // HL
		"Viewer",
		"Editor",
		"Admin", // this final comma is mandatory // HL
	}
	fmt.Println(roles)
	// END OMIT
}
