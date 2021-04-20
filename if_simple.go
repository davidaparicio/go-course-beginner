package main

import (
	"fmt"
	"strings"
)

// START OMIT
func main() {
	const gandalfWisdom = "Fool of a Took!"
	if strings.Contains(strings.ToLower(gandalfWisdom), "fool") {
		fmt.Println(`Gandalf said "fool"`)
	}
}

// END OMIT
