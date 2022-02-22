package main

import "fmt"

// START OMIT
func main() {
	lostNumbers := []int{4, 8, 15, 16, 23, 42}
	fmt.Println(&lostNumbers[0])
	printAddrOfFirstElemIfAny(lostNumbers)
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
