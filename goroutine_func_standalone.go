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
		go printInt(i, &wg) // HL
	}
	wg.Wait()
}

func printInt(i int, wg *sync.WaitGroup) { // HL
	defer wg.Done()
	fmt.Println(i)
}

// END OMIT
