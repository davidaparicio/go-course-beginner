package main

import (
	"fmt"
	"net/http"
	"os"
)

// START OMIT
func main() {
	if _, err := http.Get("https://gouglé.com/search?q=gandalf"); err != nil {
		fmt.Fprintf(os.Stderr, "failure to query Gouglé about Gandalf")
		return
	}
	fmt.Println("successfully queried Gouglé about Gandalf")
}

// END OMIT
