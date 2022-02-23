package main

import (
	"io"
	"log"
	"net/http"
)

// START OMIT
func main() {
	http.HandleFunc("/hello", helloHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func helloHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello, world!\n")
}
// END OMIT
