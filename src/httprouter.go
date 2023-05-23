package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()                                   // HL
	router.GET("/hello", hello)                                  // HL
	if err := http.ListenAndServe(":8080", router); err != nil { // HL
		log.Fatal(err)
	}
}

func hello(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) { // HL
	fmt.Fprint(w, "Hello, World!")
}
