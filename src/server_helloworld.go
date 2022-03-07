package main

import (
	"fmt"
	"log"
	"net/http"
)

// START OMIT
func helloHandler(w http.ResponseWriter, req *http.Request) { // HL
	fmt.Fprint(w, "Hello, world!\n")
}

func main() {
	http.HandleFunc("/hello", helloHandler)                   // HL
	if err := http.ListenAndServe(":8080", nil); err != nil { // HL
		log.Fatal(err)
	}
}

// END OMIT
