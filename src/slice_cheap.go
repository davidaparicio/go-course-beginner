package main

import "fmt"

// START OMIT
func main() {
	numbers := []int{4, 8, 15, 16, 23, 42}
	fmt.Println(&numbers[0])
	printAddrOfFirstElemIfAny(numbers)
	fmt.Println()
	smallerSlice := []int{1, 2, 3}
	printAddrOfFirstElemIfAny(smallerSlice)
}

func printAddrOfFirstElemIfAny(s []int) {
	if len(s) < 1 {
		return
	}
	fmt.Println(&s[0])
}

// END OMIT
