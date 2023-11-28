package main

import (
	"fmt"
	"log"
	"net/http"
)

// START OMIT
func main() {
	http.HandleFunc("GET /hello", handleHello)                // HL
	if err := http.ListenAndServe(":8080", nil); err != nil { // HL
		log.Fatal(err)
	}
}

func handleHello(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(w, "Hello, world!\n")
}

// END OMIT
