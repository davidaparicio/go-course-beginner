package main

import "fmt"

// START OMIT
func inspect(i interface{}) {
	switch v := i.(type) { // HL
	case int:
		fmt.Printf("This is an int: %d\n", v)
	case string:
		fmt.Printf("This is a string: %q\n", v)
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func main() {
	inspect(1_000_000)
}

// END OMIT
