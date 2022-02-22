package main

import "fmt"

func main() {
	// START OMIT
	roles := [...]string{"Viewer", "Editor", "Admin"}
	for i := range roles { // HL
		fmt.Println(&roles[i])
	}
	fmt.Println()
	for _, role := range roles { // HL
		fmt.Println(role)
	}
	fmt.Println()
	for i, role := range roles { // HL
		fmt.Printf("%d: %q, %t\n", i, role, &role == &roles[i])
	}
	// END OMIT
}
