### Namechecker project: command-line arguments

* modify the programme to pass usernames as command-line arguments
* store the valid and invalid ones in separate collections
* ... we need to learn about loops and slices first!

---

## `for` loops

---

### for loop with a condition

* no `while` keyword in Go!
```go
for <condition> {
  // ...
}
```

---

### `for` infinite loop

```go
for {
  // ...
}
```

---

### C-style `for` loop

```go
for <init>; <condition>; <post> {
  // ...
}
```

---

### `range`-based `for` loop

```go
for <variables> := range <data-structure> {
  // ...
}
```

* `range` works with strings, arrays, slices, maps, channels.

---

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

### Namechecker project: command-line arguments

* modify the programme to pass usernames as command-line arguments
* store the valid and invalid ones in separate collections

---

### Namecheck project: availability

* study https://golang.org/pkg/net/http/#pkg-overview
* implement the following function in the `github` package
```go
func IsAvailable(username string) (bool, error)
```
* simply send a GET to `https://github.com/<username>`
  * 404 => available
  * otherwise, unavailable

---

### Mechanics of the `defer` keyword

* delays a function call to when the enclosing function terminates
* ... whether normally or not (panic)
* the arguments of the deferred function call are evaluated at the `defer` statement!
* `defer` statements are executed in reverse order in which they appear in the enclosing function

---

### Cleaning up resources with `defer`

```go
func IsAvailable(username string) (bool, error) {
  resp, err := http.Get("https://twitter.com/" + username)
  if err != nil {
    return false, errors.New("Unknown availability")
  }
  if resp.StatusCode != http.StatusNotFound {
    resp.Body.Close() // <--- reclaim resource no longer needed
    return false, nil
  }
  resp.Body.Close() // <--- needed here too!
  return true, nil
}
```

---

### Cleaning up resources with `defer`

```go
func IsAvailable(username string) (bool, error) {
  resp, err := http.Get("https://twitter.com/" + username)
  if err != nil {
    return false, errors.New("Unknown availability")
  }
  defer resp.Body.Close() // <--- once and for all!
  if resp.StatusCode != http.StatusNotFound {
    return false, nil
  }
  return true, nil
}
```

---

### Gotcha: `defer` in a loop

* Can you spot the problem?

```go
func process(paths []string) {
  for _, path := range paths {
    file, err := os.Open(path)
    if err != nil {
      // deal with error
    }
    defer file.Close()
    // do something interesting with file...
  }
}
```

---

### Gotcha: `defer` in a loop

```go
func process(paths []string) {
  for _, path := range paths {
    if err := processOne(path); err != nil {
      // deal with error
    }
  }
}
func processOne(path string) error {
    file, err := os.Open(path)
    if err != nil {
      // deal with error
    }
    defer file.Close()
    // do something interesting with file...
}
```

---

### Gotcha: errors in deferred function calls

* Deferred statements return no result.
* What if the function call returns an error?

```go
func foo() error {
  // ...
  defer badIfThisFails() // but error is lost!
  // ...
  return err
}
```

---

### Gotcha: errors in deferred function calls

* workaround: use a named return and a func literal

```go
func foo() (err error) {
  // ...
  defer func() {
    if err1 := gobadIfThisFails(); err1 != nil {
      err = err1
  }()
  // ...
  return
}
```

* see https://www.joeshaw.org/dont-defer-close-on-writable-files/

---

### Namecheck project: write tests for `IsAvailable`

* any issue?

---

## Structs

Live demo in Playground

---

## Namecheck project: custom Twitter & GitHub types

* in package `twitter`, declare a `Twitter` struct type
* in package `github`, declare a `GitHub` struct type

---

# Methods

Live demo in Playground

---

### Namecheck project: declare methods on custom types

* turn `IsValid` into a method
* turn `IsAvailable` into a method
* declare a `String() string` method
* adjust your tests

---

# Interfaces

---

### Interfaces

* abstract type
* describes a type that has certain methods
* an interface is composed of
  * the name of a concrete type
  * a pointer to a value of that type
* THE mechanism for polymorphism in Go

---

### Interfaces are reference types

* zero value: `nil`
* pointers to interfaces? rarely (if ever) useful

---

### Interface satisfaction

* simply declare the methods required by the interface on the concrete type
* verified at compile time
* implicit!
* compare to Java, PHP, C#, etc.

```
// Java
class PokerHand implements Comparable<? super PokerHand>

// PHP
class Template implements iTemplate

// C#
class BoomerangCollection : IEnumerable
```
---

### Interfaces as contracts

* some interfaces ask implementors more than satisfying a signature
* example: [`io.Reader`](https://golang.org/pkg/io/#Reader)
* not enforcible by the compiler!
* write tests to check that the implementor is lawful!

---

### Keep interfaces small

* favour one-method interfaces
* easier to satisfy
* many examples in the standard library

---

### Naming of single-method interfaces

* method + "er"

```go
type Climber interface {
    Climb(int)
}
```

---

### Interface composition

```go
type ReadWriteCloser interface {
  Reader
  Writer
  Closer
}
```
* small interfaces are convenient for composition!

---

## Notable interface types

---

### The empty interface

* `interface{}`
* satisfied by _any_ type
* useful in some cases (`fmt.Println`)
* but not very typesafe
* avoid, if possible

---

### `fmt.Stringer`

```go
type Stringer interface {
  String() string
}
```
* governs how `fmt` functions print values of a concrete type

---

## Namecheck project: satisfy `fmt.Stringer`

* Do `twitter.Twitter` and `github.GitHub` satisfy it?
* If not, make it so!

---

### `error` interface

```go
type error interface {
  Error() string
}
```
* for any type meant to represent an error value

---

## Namecheck project: define a custom error type

```go
type ErrUnknownAvailability struct {
    Username string
    Platform string
}
```

* make it satisfy the `error` interface
* use it in `twitter.IsAvailable`

---

### Gotcha: `nil` and interfaces

* An interface that holds a `nil` pointer to a concrete type is not itself `nil`!
* Example with `ErrUnknownAvailability`
* see https://golang.org/doc/faq#nil_error

---

### `sort.Interface`

```go
type Interface interface {
    Len() int
    Less(i, j int) bool
    Swap(i, j int)
}
```

---

### `io.Reader`

```go
type Reader interface {
  Read(p []byte) (n int, err error)
}
```
* source of `[]byte`
* satisfied by `*os.File`, and many others

---

### `io.Writer`

```go
type Writer interface {
  Write(p []byte) (n int, err error)
}
```
* sink of `[]byte`
* satisfied by `*os.File`, and many others

---

### Namecheck project: define `Validator` interface

```go
type Validator interface {
  IsValid(username string) bool
}
```

* Do `twitter.Twitter` and `github.GitHub` satisfy it?
* If not, make it so!


---

### Namecheck project: define `Availabler` interface

```go
type Availabler interface {
  IsAvailable(username string) (bool, error)
}
```

* Do `twitter.Twitter` and `github.GitHub` satisfy it?
* If not, make it so!

---

## Namecheck project: define `Checker` interface

* ...by composition of
  * a `fmt.Stringer`
  * a `Validator`
  * an `Availabler`
* Do `twitter.Twitter` and `github.GitHub` satisfy it?

---

### Favour interface parameters

* if possible, take interface types as function parameters, rather than concrete types

```go
type Tree struct {...}
func (t *Tree) Save(f *os.File) error
```

---

### Favour interface parameters (cont'd)

* if possible, take interface types as function parameters, rather than concrete types
* more flexible
* easier to test!

```go
type Tree struct {...}
func (t *Tree) Save(w io.Writer) error
```

---

### Require no more, promise no less

* be wary of broad interfaces
* only ask for the required behaviour from interface parameters

```go
func (t *Tree) Save(rw io.ReadWriter) error // good?
```

---

### Require no more, promise no less (cont'd)

* only ask for the required behaviour
* similar to DDD principle of "keeping ports small"

```go
func (t *Tree) Save(w io.Writer) error // better!
```

---

## Test doubles

* Interfaces are great for mocking, stubbing, faking, etc.
* You can define an interface for a concrete type you don't own!

---

## Namecheck project: test doubles

* Define a `Client` interface that `&http.Client` satisfies
* create a `stub` package
* In it, declare
```go
type clientFunc func(url string) (*http.Response, error)
```
* Make `clientFunc` satisfy your `Client` interface

---

## Namecheck project: test doubles (cont'd)

* Define the following functions in `stub`
```go
ClientWiththError(err error) namecheck.Client
ClientWithStatusCodeAndBody(code int, body string) namecheck.Client
```
* Add a field of type `Client` to your `Twitter` and `GitHub` struct types
* Adjust your `main` and your tests for `IsAvailable`
