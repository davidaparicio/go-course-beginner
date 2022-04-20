package main

import "fmt"

func main() {
	// START OMIT
	var p *int
	fmt.Println(*p) // panic
	// END OMIT
}
