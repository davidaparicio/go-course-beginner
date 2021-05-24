package main

import (
	"fmt"
)

func main() {
	// START OMIT
	m := map[string]string{"red": "#ff0000", "blue": "#0000ff", "green": "#00ff00"}
	fmt.Println(m)

	m["white"] = "#ffffff" // HL
	fmt.Println(m)

	m = nil
	// m["white"] = "#ffffff" // causes a compilation error // HL
	// END OMIT
}
