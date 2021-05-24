package main

import (
	"fmt"
	"math"
)

// START OMIT
func main() {
	var x, n, limit float64 = 3, 3, 20
	if v := math.Pow(x, n); v < limit { // HL
		fmt.Println(v)
	}
	fmt.Println(limit)
}

// END OMIT
