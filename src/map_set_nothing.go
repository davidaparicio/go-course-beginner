package main

import (
	"fmt"
)

func main() {
	// START OMIT
	var nothing struct{}
	m := map[string]struct{}{"red": nothing, "blue": nothing, "green": nothing}
	const k = "yellow"
	if _, ok := m[k]; !ok {
		fmt.Printf("%q is not in the map.\n", k)
	}
	// END OMIT
}
