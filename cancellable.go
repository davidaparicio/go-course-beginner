package main

import (
	"fmt"
	"time"
)

// START OMIT
func main() {
	quit := make(chan struct{})
	go printInts(quit)
	time.Sleep(1 * time.Second)
	close(quit) // HL
	// do more work
}

func printInts(quit <-chan struct{}) {
	var i int
	for {
		select {
		case <-quit: // HL
			return
		default:
			fmt.Println(i)
			i++
			time.Sleep(200 * time.Millisecond)
		}
	}
}

// END OMIT
