package main

import "fmt"

func main() {
	// START OMIT
	a := []int{1, 2}
	b := a
	fmt.Println(a, b)

	b = append(b, 3) // side effect: decouples b from a! // HL
	fmt.Println(a, b)

	// a proof: some change to b's 1st element isn't reflected in a
	b[0] = 0
	fmt.Println(a, b)
	// END OMIT
}
