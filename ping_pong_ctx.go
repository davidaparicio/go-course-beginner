package main

import (
	"context"
	"fmt"
	"time"
)

type Ball struct {
	hits uint
}

// START OMIT
func main() {
	table := make(chan *Ball)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go player(ctx, "ping", table)
	go player(ctx, "pong", table)
	table <- &Ball{}
	time.Sleep(1 * time.Second)
}

func player(ctx context.Context, name string, table chan *Ball) {
	for {
		select {
		case ball := <-table:
			ball.hits++
			fmt.Printf("%s %d\n", name, ball.hits)
			time.Sleep(100 * time.Millisecond)
			table <- ball
		case <-ctx.Done():
			return
		}
	}
}

// END OMIT
