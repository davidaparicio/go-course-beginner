package main

import (
	"fmt"
	"time"
)

// START OMIT
func main() {
	go printInts()
	time.Sleep(1 * time.Second)
	// do more work...
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
