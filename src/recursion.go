package main

import "fmt"

// START OMIT
func factorial(n uint) uint {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1) // HL
}

func main() {
	const n = 5
	fmt.Println(factorial(n))
}

// END OMIT
