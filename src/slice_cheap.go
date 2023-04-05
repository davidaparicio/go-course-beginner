package main

import "fmt"

// START OMIT
func main() {
	numbers := []int{4, 8, 15, 16, 23, 42}
	fmt.Printf("addr in main:\t%p\n", &numbers[0])
	printAddrOfFirstElemIfAny(numbers) // HL

	numbers2 := []int{1, 2, 3}
	fmt.Printf("addr in main:\t%p\n", &numbers2[0])
	printAddrOfFirstElemIfAny(numbers2) // HL
}

func printAddrOfFirstElemIfAny(s []int) { // HL
	if len(s) < 1 {
		return
	}
	fmt.Printf("addr in function:\t%p\n", &s[0])
}

// END OMIT
