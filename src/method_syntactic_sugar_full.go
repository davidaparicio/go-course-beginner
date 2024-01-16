package main

// START OMIT
type T int

func (t T) Val() {}

func (t *T) Ptr() {}

func main() {
	var v T = 42
	p := &v // of type *T

	v.Val()
	p.Ptr()
	p.Val() // syntactic sugar: the compiler implicitly derefences p
	v.Ptr() // syntactic sugar: the compiler implicitly takes the address of variable f
	// T(42).Ptr() // compilation error! T(42) is not an addressable value
}

// END OMIT
