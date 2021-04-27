package main

import (
	"fmt"
)

func main() {
	// START OMIT
	m := map[string]string{"red": "#ff0000", "blue": "#0000ff", "green": "#00ff00", "empty": ""}

	v := m["red"] // HL
	fmt.Printf("red: %q\n", v)

	// fmt.Println(&m["blue"]) // causes a compilation error

	k := "empty"  // try another key here
	v, ok := m[k] // HL
	if ok {
		fmt.Printf("key %q is in the map and has value %q\n", k, v)
	}
	// END OMIT
}
