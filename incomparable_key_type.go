package main

// START OMIT
func main() {
	var (
		_ map[int]string
		// _ map[[]int]string // causes a compilation error
	)
}

// END OMIT
