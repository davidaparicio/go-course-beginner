package main

import (
	"fmt"
)

func main() {
	// START OMIT
	var nothing struct{}
	m := map[string]struct{}{"red": nothing, "blue": nothing, "green": nothing}
	const k = "yellow"

	// checking for membership
	if _, ok := m[k]; !ok {
		fmt.Printf("%q is not in the map.\n", k)
	}

	// adding an element
	m["white"] = nothing

	// deleting an element
	delete(m, "red")

	fmt.Println(m)
	// END OMIT
}
