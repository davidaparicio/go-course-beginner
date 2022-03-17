package main

import (
	"fmt"
)

func main() {
	// START OMIT
	m := map[string]bool{"red": true, "blue": true, "green": true}
	const k = "yellow"

	// checking for membership
	if exists := m[k]; !exists {
		fmt.Printf("%q is not in the map.\n", k)
	}

	// adding an element
	m["white"] = true

	// deleting an element
	delete(m, "red")

	fmt.Println(m)
	// END OMIT
}
