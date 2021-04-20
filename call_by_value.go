package main

import (
	"fmt"
)

// START OMIT
func main() {
	i := 41
	increment(i)
	fmt.Println("variable i at the end of 'main':", i)
}

func increment(i int) {
	i++
	fmt.Println("variable i at the end of 'increment':", i)
}

// END OMIT
