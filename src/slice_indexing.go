package main

import "fmt"

func main() {
	// START OMIT
	roles := []string{"Viewer", "Editor", "Admin"}
	fmt.Println(&roles[0]) // HL
	fmt.Println(&roles[1])
	fmt.Println(&roles[2])
	// fmt.Println(roles[len(roles)]) // panic // HL
	// END OMIT
}
