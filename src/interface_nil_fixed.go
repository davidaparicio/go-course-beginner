package main

import (
	"fmt"
	"os"
)

type ErrLaunchFailure struct{}

func (*ErrLaunchFailure) Error() string {
	return "launcher: failure to launch"
}

// START OMIT
func launchMissile() error { // HL
	return nil
}

func main() {
	var err error
	if err = launchMissile(); err != nil { // HL
		fmt.Println("failed!")
		os.Exit(1)
	}
}

// END OMIT
