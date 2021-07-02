package main

import (
	"fmt"
	"sync"
	"time"
)

// START OMIT
func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		printInts()
	}()
	wg.Wait()
}

func printInts() {
	var i int
	for {
		fmt.Println(i)
		i++
		time.Sleep(200 * time.Millisecond)
	}
}

// END OMIT
