package main

import "fmt"

func main() {
	// START OMIT
	names := []string{"foo", "bar", "baz"}
	red := names[:2]
	names = red[:cap(red)] // HL
	fmt.Println(names)
	// END OMIT
}
