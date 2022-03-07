package main

import "fmt"

func main() {
	// START1 OMIT
	roles := []string{"Viewer", "Editor", "Admin"}
	fmt.Println(&roles[0]) // HL
	fmt.Println(&roles[1])
	fmt.Println(&roles[2])
	// fmt.Println(roles[len(roles)]) // panic // HL
	// END1 OMIT
	// START2 OMIT
	roles = nil
	// fmt.Println(roles[0]) // panic // HL
	// END2 OMIT
}
