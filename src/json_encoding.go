package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// START1 OMIT
type User struct {
	Name    string `json:"name"`     // HL
	IsAdmin bool   `json:"is_admin"` // HL
}

// END1 OMIT

func getAdmin(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	// START2 OMIT
	admin := User{Name: "admin", IsAdmin: true} // HL
	w.Header().Set("Content-Type", "application/json")
	enc := json.NewEncoder(w)                 // HL
	if err := enc.Encode(admin); err != nil { // HL
		w.WriteHeader(http.StatusInternalServerError)
	}
	// END2 OMIT
}

func createAdmin(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	if r.Header.Get("Content-Type") != "application/json" {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		return
	}
	// START3 OMIT
	var data struct { // HL
		Name string `json:"name"`
	}
	dec := json.NewDecoder(r.Body)            // HL
	if err := dec.Decode(&data); err != nil { // HL
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// omitted: instantiate admin User and save it to the database
	w.WriteHeader(http.StatusCreated)
	// END3 OMIT
}

func main() {
	http.HandleFunc("/admin", getAdmin)
	http.HandleFunc("/admin/new", createAdmin)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
