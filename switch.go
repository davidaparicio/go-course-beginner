package main

import "fmt"

// START OMIT
func main() {
	r := 'b'
	switch r {
	case 'a', 'e', 'i', 'o', 'u':
		fmt.Println("ASCII vowel")
	case 'y':
		fmt.Println("ASCII vowel (en France !)")
	default:
		fmt.Println("not an ASCII vowel")
	}
}

// END OMIT
