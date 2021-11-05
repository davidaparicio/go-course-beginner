package main

import (
	"fmt"
	"sync"
)

// START OMIT
func main() {
	m := make(map[string]uint)
	var wg sync.WaitGroup
	const (
		concurrency = 100
		updates     = 100
		username    = "jub0bs"
	)
	wg.Add(concurrency)
	for i := 0; i < concurrency; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < updates; j++ {
				m["jub0bs"]++ // HL
			}
		}()
	}
	wg.Wait()
	fmt.Println("count:", m[username])
}

// END OMIT
