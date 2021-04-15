package main

import (
	"fmt"
	"sync"
)

// START OMIT
func main() {
	var visits uint
	var wg sync.WaitGroup
	const concurrency = 50
	const n = 1000
	wg.Add(concurrency)
	for i := 0; i < concurrency; i++ {
		go func() {
			for j := 0; j < n; j++ {
				visits++ // HL
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("visits:", visits)
}

// END OMIT
