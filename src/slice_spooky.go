package main

import "fmt"

func main() {
	// START OMIT
	numbers := []int{4, 8, 15, 16, 23, 42}
	first3 := numbers[:3]
	first3[0] = 99
	fmt.Println(first3)
	fmt.Println(numbers)
	// END OMIT
}
