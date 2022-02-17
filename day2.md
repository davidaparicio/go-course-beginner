## Arrays

---

### Arrays

* fixed-length, homogenous container
* elements are contiguous in memory
* array length is encoded in the type (e.g. `[3]int`)
* zero value: array full of element type's zero value

---

### Array literals
```go
words := [3]string{"foo", "bar", "baz"}
moreWords := [...]string { // size is optional
  "qux",
  "quux", // the last comma is mandatory
}
```

---

### Operations on array

* indexing  (same as for strings)
* array elements are variables: `&arr[i]` is legal
* get a slice from an array: `words[:]`
* ranging over an array
```go
for i, v := range arr {
  // ...
}
```

---

### Array are no reference types

* fully copied when passed as params
* arrays are seldom used directly

---

## Slices

---

### Slices

* fundamental, ubiquitous type in Go
* a resizable view inside an array
* a slice of `T` has type `[]T`

---

### Slices: under the hood

* a slice is composed of
  * a pointer to the underlying array
  * a length (`len`)
  * a capacity (`cap`)

---

### Capacity of a slice

> The array underlying a slice may extend past the end of the slice.
> The capacity is a measure of that extent: it is the sum of the length of the
> slice and the length of the array beyond the slice.

([source](https://golang.org/ref/spec#Slice_types))

---

### Slice is a reference type

* zero value: `nil`
* a function with a slice parameter gets a copy of it
* but the underlying array is not copied!

---

### Initializing a slice

```go
s := make([]string, l)    // l: length
s := make([]string, l, c) // c: capacity
```

---

### Slice literal

```go
words := []string{"foo", "bar", "baz"}
moreWords := []string {
  "qux",
  "quux", // the last comma is mandatory
}
```

---

### Slices are not comparable

* equality comparison disallowed by compiler
* a slice can only be compared to `nil`

---

### Indexing into a slice

* same syntax as for strings
* indexing into a `nil` slice: panic!
* slice elements are variables: `&s[i]` is legal

---

### Slicing operator

* `s1 := s[i:j]`
* if `j` exceeds `cap(s)`, panic!
* yields a different view into the same array

---

### Namechecker project: command-line argument

* modify the programme to pass one username as the 1st command-line argument

---

### Ranging over a slice

```go
for i, v := range s {
  // ...
}
```

---

### Namecheck project: use table-driven subtests

See https://go.dev/blog/subtests#table-driven-tests-using-subtests

---

### Appending an element to a slice

* use the built-in function `append`
* `s = append(s, elem)`
* built-in `append` copies elements to a larger underlying array if capacity has been reached
* https://play.golang.org/p/dXwZYnoB_Qj

---

### Appending another slice to a slice

* `s = append(s, s1...)`
* the `...` operator
  * explodes a slice
  * allows you to pass it where a variadic argument is expected

---

### Aliasing

* multiple slices may use the same underlying array
* careful: don't hold on to a pointer to a slice element while appending to it


---

## Interfaces

---

### Interfaces under the hood


---

### The empty interface

* `interface{}`
* satisfied by _any_ type
* useful in some cases (`fmt.Println`)
* but not very typesafe
* avoid, if possible

---

### `error` interface

```go
type error interface {
  Error() string
}
```
* for any type meant to represent an error value

---

### `io.Reader`

```go
type Reader interface {
  Read(p []byte) (n int, err error)
}
```
* source of `[]byte`
* reads the data in the supplied buffer, returns how many bytes were read, and any error that may have occurred
* satisfied by `os.Stdin`, `*os.File`, `*bytes.Buffer`.

---

### `io.Writer`

```go
type Writer interface {
  Write(p []byte) (n int, err error)
}
```
* sink of `[]byte`
* writes the data in the supplied buffer, returns how many bytes were written, and any error that may have occurred
* satisfied by `os.Stdout`, `os.Stderr`, `*os.File`, `*bytes.Buffer`, etc.

---

### Interfaces as contracts

* some interfaces specify implementors to follow certain additional rules
* example: [`io.Reader`](https://pkg.go.dev/io/#Reader)
* those rules are not enforcible by the compiler!
* the onus is on you to respect them when you implement an interface

---

### Gotcha: `nil` and interfaces

* An interface that holds a `nil` pointer to a concrete type is not itself `nil`!
* [Find the bug!](https://play.golang.org/p/8M8YjK9jjZw)
* see https://golang.org/doc/faq#nil_error
