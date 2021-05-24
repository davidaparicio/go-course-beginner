package main

import (
	"fmt"
)

// START OMIT
func main() {
	enToFr := map[string]string{}
	add(enToFr, "one", "un")
	fmt.Printf("at the end of 'main': %v\n", enToFr)
}

func add(m map[string]string, k string, v string) {
	m[k] = v
	fmt.Printf("at the end of 'add': %v\n", m)
}

// END OMIT
