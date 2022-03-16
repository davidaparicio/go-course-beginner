package main

import (
	"fmt"
	"log"
	"net/http"
)

// START OMIT
func handleHello(w http.ResponseWriter, req *http.Request) { // HL
	fmt.Fprint(w, "Hello, world!\n")
}

func main() {
	http.HandleFunc("/hello", handleHello)                    // HL
	if err := http.ListenAndServe(":8080", nil); err != nil { // HL
		log.Fatal(err)
	}
}

// END OMIT
