package main

import (
	"fmt"
	"time"
)

type Ball struct {
	hits uint
}

// START OMIT
func main() {
	table := make(chan *Ball)
	done := make(chan struct{}, 1)
	defer close(done)
	go player(done, "ping", table)
	go player(done, "pong", table)
	table <- &Ball{}
	time.Sleep(1 * time.Second)
}

func player(done <-chan struct{}, name string, table chan *Ball) {
	for {
		select {
		case ball := <-table:
			ball.hits++
			fmt.Printf("%s %d\n", name, ball.hits)
			time.Sleep(100 * time.Millisecond)
			table <- ball
		case <-done:
			return
		}
	}
}

// END OMIT
