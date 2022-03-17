package main

import "fmt"

// START1 OMIT
type Climber interface {
	Climb(meters int) error // HL
}

type Mountaineer struct {
	Name string
}

func (m *Mountaineer) Climb(meters int) error { // HL
	fmt.Printf("%s climbs %d meters\n", m, meters)
	return nil
}

// END1 OMIT
func main() {
	// START2 OMIT
	var c Climber = &Mountaineer{Name: "Lynn Hill"} // HL
	fmt.Println(c)
	// END2 OMIT
}
