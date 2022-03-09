package main

import (
	"fmt"
	"io"
)

func main() {
	// START OMIT
	var r io.Reader
	fmt.Println(r == nil)
	// END OMIT
}
