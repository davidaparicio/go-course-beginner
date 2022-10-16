package main

import "fmt"

// START OMIT
type User struct {
	name    string
	isAdmin bool
	manager *User
}

func main() {
	var u User
	fmt.Printf("%+v\n", u)
}

// END OMIT
