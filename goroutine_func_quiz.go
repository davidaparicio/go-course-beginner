package main

import (
	"fmt"
	"sync"
)

// START OMIT
func main() {
	printTenIntsConcurrently()
}

func printTenIntsConcurrently() {
	var wg sync.WaitGroup
	const n = 10
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			fmt.Println(i)
		}()
	}
	wg.Wait()
}

// END OMIT
