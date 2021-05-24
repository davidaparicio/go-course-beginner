package main

import (
	"fmt"
	"sync"
)

// START OMIT
func main() {
	c := make(chan string)
	var wg sync.WaitGroup
	wg.Add(2)
	go produce(c, &wg)
	go consume(c, &wg)
	wg.Wait()
}

func produce(c chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	// <-c // would cause a compilation error // HL
	c <- "hello"
}

func consume(c <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	// c <- "hi" // would cause a compilation error // HL
	// close(c)  // would cause a compilation error // HL
	fmt.Println(<-c)
}

// END OMIT
