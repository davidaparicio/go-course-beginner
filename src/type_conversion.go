package main

import "fmt"

// START OMIT
type Celsius int
type Fahrenheit int

func main() {
	var freezingC Celsius = 0
	var freezingF Fahrenheit
	// freezingF = freezingC // compilation error
	freezingF = Fahrenheit(freezingC + 32)
	fmt.Println(freezingC, freezingF)
}
// END OMIT

