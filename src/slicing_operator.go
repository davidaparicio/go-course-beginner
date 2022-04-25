package main

import "fmt"

func main() {
	// START OMIT
	names := []string{"foo", "bar", "baz", "qux", "quux"}
	fewerNames := names[2:4] // HL
	fmt.Printf("fewerNames has type %T and value %v\n", fewerNames, fewerNames)

	// names and fewerNames use the same backing array
	fmt.Println(&fewerNames[0] == &names[2] && &fewerNames[1] == &names[3]) // HL
	// END OMIT
}
