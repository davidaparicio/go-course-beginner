package main

import (
	"fmt"
)

func main() {
	// START OMIT
	m := map[string]bool{"red": true, "blue": true, "green": true}
	const k = "yellow"
	if ok := m[k]; !ok {
		fmt.Printf("%q is not in the map.\n", k)
	}
	// END OMIT
}
