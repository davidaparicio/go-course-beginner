package main

import (
	"fmt"
	"sync"
)

func main() {
	printTenIntsConcurrently()
}

// START OMIT
func printTenIntsConcurrently() {
	var wg sync.WaitGroup
	const n = 10
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func(j int) { // HL
			defer wg.Done()
			fmt.Println(j) // HL
		}(i) // HL
	}
	wg.Wait()
}

// END OMIT
