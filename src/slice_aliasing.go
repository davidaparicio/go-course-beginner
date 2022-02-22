package main

import "fmt"

// START OMIT
type User struct {
	name    string
	isAdmin bool
}

func main() {
	users := []User{
		User{name: "Julien"},
	}
	julien := &users[0]                     // HL
	users = append(users, User{name: "Ed"}) // HL
	julien.isAdmin = true
	fmt.Println(*julien)
	fmt.Println(users)
	fmt.Printf("%p\n", julien)
	fmt.Printf("%p\n", &users[0])
}

// END OMIT
