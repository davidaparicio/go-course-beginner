package main

// START OMIT
import (
	"fmt"
	"time"
)

func grindCoffeeBeans() {
	fmt.Println("Grinding coffee beans...")
}

func frothMilk() {
	fmt.Println("Frothing milk...")
}

func main() {
	go grindCoffeeBeans()
	go frothMilk()
	time.Sleep(100 * time.Millisecond) // HL
}

// END OMIT
