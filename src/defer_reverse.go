package main

import "fmt"

// START OMIT
func main() {
	for i := 0; i < 3; i++ {
		defer fmt.Println(i) // HL
	}
}

//END OMIT
