package main

import "fmt"

func main() {
	// START OMIT
	userIds := [3]int{42, 56, 70} // HL
	fmt.Printf("type name: %T\n", userIds)
	l := len(userIds) // HL
	fmt.Printf("length: %d\n", l)
	roles := [...]string{"Viewer", "Editor", "Admin"} // HL
	fmt.Println(roles == roles)                       // HL
	var funcs [0]func()
	fmt.Println(funcs)
	// fmt.Println(funcs == funcs) // compilation error  // HL
	// END OMIT
}
