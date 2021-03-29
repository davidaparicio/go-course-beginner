### Installation

* Go v1.16
* Visual Studio Code
* Go VSCode extension
* Git
* Wifi
* Slides: https://github.com/jub0bs/intro-to-golang

---

### What to expect from the course

* an invitation to Go
* emphasis on practice over theory
* interaction
* regular breaks

---

### Project: username checker

* checking a username on social networks
  * validity
  * availability
* business case
  * brand consistency
  * username squatting (😈)
* a forcing function for learning Go

---

### Go's history

* created at Google in 2007
* ...by [Rob Pike, Robert Griesemer, Ken Thompson](https://res.cloudinary.com/practicaldev/image/fetch/s--opTYcTMa--/c_limit%2Cf_auto%2Cfl_progressive%2Cq_auto%2Cw_880/https://cdn-images-1.medium.com/max/800/1%2A5O-YqwO94QUQ8OkNypcVdQ.png)
* open-sourced in 2009
* Go 1.0 released in 2012
* Go 1.16 (latest release)
* powers Docker, Kubernetes, Terraform, and many more

---

### What makes Golang special

* syntax reminiscent of C
* strong typing
* fast compilation
* garbage collector (fast & easy to tune)
* pointers, but no pointer arithmetic
* full Unicode support

---

### What makes Golang special (cont'd)

* multiple return values
* first-class functions
* object-oriented, but no inheritance
* built-in concurrency
* no operator overloading
* no generics (yet!)

---

### What makes Go's ecosystem special

* BSD licence
* an agenda of simplicity
* stability of the language and its standard library
* no centralised package repository
* rich standard library
* rich command-line interface (gofmt!)
* a welcoming community

---

### Go playground

* https://play.golang.org
* de facto online REPL
* automatic imports
* automatic formatting
* permalinks

---

### Go playground limitations

* mostly for a single main package
* no networking
* no command-line arguments
* only standard-library imports

---

# First steps in Go

---

### [Hello, World!](https://play.golang.org/p/7vin2BK8_A6)

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello, 世界")
}
```

---

### Hello, World!

* all source files start with a package clause
* `main` package: defines an executable program
* `main` function: program's entrypoint
* import-declaration syntax
* no semicolons needed to terminate statements
* tabs, not spaces (by convention)

---

### Namecheck project: getting started

* create a `namecheck` folder
* open it in VS Code
* create a `main.go` file in it
* write a "Hello World" programme

---

### Compiling and executing a Go program

```shell
$ go build main.go
$ ./main
```

---

### Executing a Go program in one step

```shell
$ go run main.go
```

roughly equivalent to


```shell
$ go build main.go
$ ./main
$ rm ./main
```

---

### Keywords

```txt
break       default       func      interface   select
case        defer         go        map         struct
chan        else          goto      package     switch
const       fallthrough   if        range       type
continue    for           import    return      var
```

---

### Predeclared names

* Constants
```txt
false true nil iota
```
* Built-in types
```txt
bool int uint rune float64 string ...
```
* Built-in functions
```txt
make    len     cap     new
append  copy    delete
complex real    imag
panic   recover
```

---

### Comments

* end-of-line comments:

```go
var foo = "foo" // this is a comment
```

* multiline comments (do not nest)

```go
/*
  this is
  a multiline comment
*/
```

---

## Packages

---

### Package path

* globally unique
* examples
  * `fmt`
  * `net/http`
  * `github.com/jub0bs/missilelauncher`

---

### Default package name

* the last component of the package path
* examples
  * `fmt`
  * `http`
  * `missilelauncher`

---

### Normal imports

```go
import "github.com/jub0bs/missilelauncher"

func main() {
  missilelauncher.NukeThemAll()
}
```

members of the package are accessible via the default package name

---

### Named imports

* allows you to use a different name for the package locally
* useful when the default package name is long
```go
import ml "github.com/jub0bs/missilelauncher"
```
* useful for resolving name collisions
```
import "math/rand"
import crand "crypto/rand"
```

---

### Blank imports

```go
import _ "github.com/lib/pq"
```

more on that later

---

### Import block

* possible to factor multiple import declarations

```go
import (
  "encoding/json"
  "fmt"
  "io"
)
```

---

### Elements of style: organising imports

```go
import (
  // standard-library packages

  // 3rd-party packages

  // your packages
)
```

---

### Package: exported and unexported

* no encapsulation at the type level
* only two access levels
  * exported (public)
  * unexported (package-private)
* the case of an identifier's initial matters
  * upper case: exported
  * lower case: unexported

---

# Variables

---

## Local variables must be used

* A variable declared within a function _must_ be used.
* Otherwise, compilation error!

---

## Variable declaration

```go
var <name> <type> = <expression>
```

* type is optional...
* initialising expression is optional...
* ...but not both!

---

### Variable blocks

possible to factor multiple `var` declarations

```go
var (
	legalRegexp   = "^[A-Za-z0-9]*$"
	illegalPattern = "--"
)
```

---

### Elements of style: variable naming

* strive for conciseness
* camel case (no underscores)
* the smaller the scope, the shorter the name

---

### Short variable declaration

```go
<name> := <expression>
```

* type is inferred from expression
* there must be new values on the left-hand side!
* unused declared variable => compilation error!
* use the _blank identifier_ (`_`) when needed
* shadowing (due to lexical scoping) can lead to bugs

---

### Initial value of a variable

* no uninitialised variable in Go!
* implicitly assigned the _zero value_ of its type
```go
var i int
fmt.Printf("value: %d\n", i) // "value: 0"
```

---

### `var` or `:=`

* `:=` only allowed for local variables
* favour `var` over `:=` when
  * the initial value must be the zero value
  * the initial value doesn't matter
  * the default type of the initialising expression won't do

---

### Namecheck project: 

* declare a `username` string variable
* initialize it with your GitHub/Twitter username
* use a function from the `fmt` package to print the message
```txt
Hello, <username>
```
* `fmt`'s documentation is available at https://golang.org/pkg/fmt/

---

## Pointers

* In Go, a variable is an _addressable value_.
* A pointer is a value that holds the address of a variable.

---

### Dereferencing a pointer

* `&`: address-of operator
* `*`: dereference operator
* https://play.golang.org/p/GUVuB8Ao6vU

---

### Pointer types

* `*T` denotes the type of a pointer to a variable of type `T`
* https://play.golang.org/p/zHvpWOITgVm

---

### Pointers are reference types

* zero value: `nil`
* Deferencing a `nil` pointer causes a _panic_ (unrecoverable failure).
* https://play.golang.org/p/hj95etHEX3j


---

# Strings

---

### UTF-8

* the most popular string encoding
* Go strings are encoded in UTF-8
* Go source files _must_ be UTF-8
* each code point takes 1 to 4 bytes
* a glyph/character can be composed of multiple code points

---

### Runes

* `rune` is a type that represents a Unicode code point
* interchangeable with `int32`
* To get the number of bytes in a given rune, use the `unicode/utf8` package.
* https://play.golang.org/p/G3LTF-HzfVb

---

### Rune literals

* delimited by single quotes: `'a'`, `'汉'`
* escape sequences:
  * `'\n'`
  * `'\uhhhh'`
  * `'\Uhhhhhhhh'`

---

### Strings

* immutable data structure
* under the hood
    * a pointer to an array of bytes
    * an int corresponding to the length
* zero value: empty string
* concatenation operator: `+`

---

### String literals

* delimited by double quotes
```go
s := "Hello World!"
```
* alternatively, backticks (_raw string_)
* useful for multiline strings and string literals containing `"`
```go
`json:"Name"`
```

---

### String manipulation

* `strings`
* `unicode/utf8`
* `strconv`
* `regexp`

---

### String length

```go
s := "Hello, World!"
len(s)                    // number of bytes
utf8.RuneCountInString(s) // number of runes
```

---

### Indexing a string

* 0-based indexing
* `s[i]` is the `i+1`th byte in string `s`
* if `i` out of bounds, _panic_!

---

### Substring operator

* structural sharing!
* syntax: `s[lo:hi]`
* semi-open interval: from `lo` to `hi` (exclusive)
* if _hi_ or _lo_ out of bounds or _hi < lo_, panic!

---

### Substring operator (cont'd)

`lo` and `hi` are optional
* `s[:hi]` is equivalent to `s[0:hi]`
* `s[lo:]` is equivalent to `s[lo:len(s)]`
* `s[:]` is equivalent to `s`

---

### Ranging over a string

* iterates over the _runes_ of a string
```
for i, r := range s {
 ...
}
```
* exercise: print each rune and its index

---

### Named types (user-defined types)

* `type` declaration
```go
type <name> <underlying_type>
```
* underlying type can be arbitrary
* default value: default value of underlying type
* _methods_ can be defined on named types!

---

### Conversion

* strong typing
* no implicit casting in Go
* (almost) all conversions must be explicit

---

### Constants

* values fixed at compile time
* `const <name> = <expression>`
* no function calls allowed on right-hand side
* allow optimisations by the compiler
* two categories of constants: typed and untyped


---

### Typed constants

```go
const Monday int = 1
```

---

### Untyped constants

* constants not committed to a particular type
* example: [`math` constants](https://golang.org/pkg/math/#pkg-constants)
* flexibility!
* numeric untyped constants: at least 256 bits of precision
* six _kinds_ of untyped constants

---

### Constant blocks

possible to factor multiple const declarations

```go
const (
  Un = 1
  Deux = 2
  Trois = 3
)
```

---

### `iota` constant generator

* syntactic sugar for creating numeric constants

```go
const (
  Un = iota + 1
  Deux
  Trois
)
```

* RHS expression is implicitly reused within the constant block
* `iota` evaluates to `0`, then `1`, then `2`, etc.

---

### Enum pattern

* no native enumeration types
* instead, named type based on `int`
* not very typesafe
* exercise: write a `Month` enum

---

### Enum pattern: example

```go
type Month int
const (
  _ Month = iota
  January
  February
  //...
  December
)
```


---

# Basic control-flow structures

---

### `if` (basic form)

```go
if <condition> {
  // ...
}

```
* condition is of boolean type
* no parentheses around condition
* braces are mandatory

---

### `if` (with optional declarations)

```go
if [short-var-declarations;] <condition> {
  // ...
}
```

* variables scoped to the `if` statement
* examples in Go playground

---

### `if`/`else`

```go
if <condition> {
  // ...
} else {
  // ...
}
```

* `else` is seldom used in Go
* avoid long `if`/`else` chains

---

### `switch`

```go
switch <expression> {
case <expression>:
  // ...
case <expression>:
  // ...
default:
  // ...
}
```

---

### `switch` (cont'd)

* cases are evaluated from top to bottom
* the first match is executed
* if none match, the (optional) default case is executed
* `default` can occur at any place among the cases

---

### `switch` cases with multiple values

```go
r := 'b'
switch r {
case 'a', 'e', 'i', 'o', 'u':
  fmt.Println("ASCII vowel")
default:
  fmt.Println("not an ASCII vowel")
}
```

---

### Tagless `switch`

```go
switch {
case <condition>:
  ...
case <condition>:
  ...
...
default:
  ...
}
```
* favour a tagless switch to `if`/`else` chains
* exercise: write a signum function

---

### `for` loop

```go
for <condition> {
  // ...
}
```

* no `while` keyword in Go!

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

# Functions

---

## Function declaration

```
func <name>(<param-list>) <result-list> {
  // ...
  return <result-values>
}
```

---

### Multiple results

```
func CountWords(path string) (int, error) {
  // ...
}
```

---

### Named results

```
func CountWords(path string) (count int, err error) {
  // ...
  return // bare return
}
```
* useful for
  * clarifying the roles of results
  * simplifying some implementations

---

### Type factorisation in param list

```
func first(a, b int) int {
  return a
}
```

---

### Variadic functions

* function that accepts any number of trailing arguments
```go
func max(first int, rest ...int) int {
  // ...
}
```
* variadic parameter `rest` is a slice of `int`s
* see also [`fmt.Println`](https://golang.org/pkg/fmt/#Println)

---

### Function call

* usual syntax (parens)
* all arguments must be specified
  * no default parameter value
  * no "named parameters"
* calling a `nil` function: panic!

---

### Parameter evaluation

* call-by-value evaluation
* "reference types"
  * built-in types that are glorified pointers
  * pointers, functions, slices, maps, channels

---

### Recursion

* a function can call itself
* exercise: write a "factorial" function
  ([solution](https://play.golang.org/p/yvRytLgk12U))

---

### Functions are first-class values

* a function can be
    * another function's parameter
    * another function's result

---

### Function literals

```go
f := func (i int) int {
  return i + 1
}
```

(no lambdas)

---

### Function result

* The result of a function can itself be a function.
* exercise: write a function that
  * takes a `name` parameter (of type `string`)
  * returns a function that prints "Hello, <name>"
* [solution](https://play.golang.org/p/zQ67huh73x-)

---

### Closures

* functions can capture variables in their environment
* example: [stateful function](https://play.golang.org/p/kcYOX2eGxWi)

---

### Project: set up directory structure

```txt
namecheck
└── main.go
```

---

### Project: validation (Twitter)

```go
func isLongEnough(username string) bool {
  // returns true if username is not empty
}

func isShortEnough(username string) bool {
  // returns true if username contains 15 chars or fewer
}

func containsNoIllegalPattern(username string) bool {
  // returns true if username does not contain "twitter"
  // (and all case variations, incl. "TwItTEr", etc.)
}
```

---

# Error handling & reporting

---

### Signaling a failure to the caller

* a function that can fail returns an additional value
* this value comes last, by convention
* only one possible cause: `bool`
* multiple possible causes: `error`
* non-nil error: something bad happened

---

### Error-check idiom

* convention: keep the happy path "on the left"
* (not nested in if-else statements)

---

### Error-check idiom

Bad:
```go
if err := fallibleFoo(); err == nil {
  // happy path :)
} else {
  // unhappy path :(
}
```

---

### Error-check idiom (cont'd)

Good:
```go
if err := fallibleFoo(); err != nil {
  // unhappy path :(
  // early return
}
// happy path :)
```

---

### Beyond `err != nil`

* to find out more about the nature of the error,
  assert on the error's behaviour
* more on that later (when we know about interfaces)

---

### Error handling

* always check for errors
* wrap low-level errors in higher-level ones
* few errors should be ignored

---

### `panic`

* similar to Java's `throw`
* signifies an _unrecoverable_ failure
* Do _not_ use `panic` for "expected" and recoverable failures, e.g.
  * a file couldn't be opened
  * a HTTP request failed

---

## Project: validation (cont'd)

```go
func containsOnlyLegalChars(username string) ??? {
  // returns true if username matches ^[0-9A-Z_a-z]*$
}
```
* Design question: what should the result list (`???`) be?

---

## Project: validation (cont'd)

```go
func IsValid(username string) bool {
  // returns true if all four predicates return true
}
```

---

## Project: create a `twitter` package

```txt
namecheck
├── main.go
└── twitter
    └── twitter.go
```

---

## Project: initialise a module

```shell
$ go mod init github.com/<your-GitHub-username>/namecheck
```

A `go.mod` file will be created:

```txt
namecheck
├── go.mod
├── main.go
└── twitter
    └── twitter.go
```

---

## Project: import twitter package in main

```go
package main

import "github.com/jub0bs/namecheck/twitter"

func main() {
  username := "jub0bs"
  fmt.Printf(
    "%s is valid on Twitter: %t\n,
    username,
    twitter.IsValid(username),
  )
}
```

---

# Testing

---

### `testing` package

* unit tests
* example tests (for documentation)
* micro-benchmarks

---

### Test files

* Test files for package `foo`
  * must have a `_test` suffix
  * must be placed in `foo` folder

```txt
foo
├── foo.go
├── foo_test.go
```

---

### White-box vs. black-box testing

determined by package clause
* white-box
  * `package twitter`
  * (favoured by the Go community)
* black-box
  * `package twitter_test`
  * (also useful for breaking dependency cycles)

---

### Test functions

* declared like normal functions in
* naming constraints
  * `Test` prefix
  * followed by an uppercase letter!
* take a single parameter of type `*testing.T`
* return no results

---

### Test functions: example

```go
func TestUsernameTooLong(t *testing.T) {
  username := "obviously_longer_than_15_chars"
  want := false
  got := IsValid(username)
  if got != want {
    t.Errorf(
      "IsValid(%s) = %t; want %t",
      username,
      got,
      want,
    )
  }
}
```

---

### Running tests of current package

```
$ go test
```

---

### Running tests of current package and subpackages

```
$ go test ./...
```

---

### Running specific tests

```
$ go test -run=<regexp-for-test-method-names>
```

---

### Project: validation (cont'd)

Write tests for `IsValid` in `twitter_test.go`

```txt
namecheck
├── go.mod
├── main.go
└── twitter
    ├── twitter.go
    └── twitter_test.go
```

---

### Code coverage
```
$ go test -coverprofile="coverage.out" ./...
$ go tool cover --html="coverage.out"
```

---

## Project: create a `github` package

* adapt it from package `twitter`
* different rules!

---

### Project: validation (GitHub)

```go
func isLongEnough(username string) bool {
  // returns true if username is not empty
}

func isShortEnough(username string) bool {
  // returns true if username contains 39 chars or fewer
}

func containsOnlyLegalChars(username string) ??? {
  // returns true if username matches ^[-0-9A-Za-z]*$
}
```

---

### Project: validation (GitHub) cont'd

```go
func containsNoIllegalPattern(username string) bool {
  // returns true if username does not contain "--"
}

func containsNoIllegalPrefix(username string) bool {
  // returns true if username does not start with "-"
}

func containsNoIllegalSuffix(username string) bool {
  // returns true if username does not end with "-"
}
```

---

## Project: validation (GitHub) (cont'd)

```go
func IsValid(username string) bool {
  // returns true if all predicates above return true
}
```
