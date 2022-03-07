package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		printInts()
	}()
	wg.Wait()
}

// START OMIT
func printInts() {
	for i := 0; ; i++ {
		fmt.Println(i)
		time.Sleep(200 * time.Millisecond)
	}
}

// END OMIT
