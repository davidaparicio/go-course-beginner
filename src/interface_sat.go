package main

import "fmt"

// START OMIT
type Climber interface {
	Climb(int) error
}

type Mountaineer struct {
	Name string
}

func (m *Mountaineer) Climb(height int) error {
	fmt.Printf("%s climbs %d meters\n", m, height)
	return nil
}

func main() {
	var c Climber = &Mountaineer{Name: "Lynn Hill"} // HL
	fmt.Println(c)
}

// END OMIT
