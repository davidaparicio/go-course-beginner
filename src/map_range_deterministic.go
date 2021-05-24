package main

import (
	"fmt"
	"sort"
)

func main() {
	// START OMIT
	m := map[string]string{
		"red":   "#ff0000",
		"blue":  "#0000ff",
		"green": "#00ff00",
	}
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Printf("%s: %q\n", k, m[k])
	}
	// END OMIT
}
