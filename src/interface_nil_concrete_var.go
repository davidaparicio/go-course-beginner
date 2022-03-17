package main

import (
	"fmt"
	"os"
)

// START OMIT
type ErrLaunchFailure struct{}

func (*ErrLaunchFailure) Error() string {
	return "launcher: failure to launch"
}

func launchMissile() error { // HL
	var err *ErrLaunchFailure
	return err
}

func main() {
	var err error
	if err = launchMissile(); err != nil { // HL
		fmt.Println("failed!")
		os.Exit(1)
	}
}

// END OMIT
