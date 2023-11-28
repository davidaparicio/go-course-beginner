package main

import "fmt"

func main() {
	// START OMIT
	var count int
	go func() {
		count++
	}()
	if count == 0 {
		fmt.Println(count)
	}
	// END OMIT
}
