# Hello World

Go has a helper CLI
```go
$ go help mod 
$ go help mod init
```

```go
// hello.go
package main

import "fmt"

func main() {
	fmt.Println("Hello, world")
}

```
## Run
```go
$ go run hello
```

## Testing
```go
// hello_test.go
package main

import "testing"

func TestHello(t *testing.T) {
    got := Hello()
    want := "Hello, world"

    if got != want {
        t.Errorf("got %q want %q", got, want)
    }
}

```

## Create a module
```go
$ go mod init hello
```

## Run Test
```go
$ go test
```
About testing in go:
- It needs to be in a file with a name like xxx_test.go
- The test function must start with the word Test
- The test function takes one argument only t *testing.T
- In order to use the *testing.T type, you need to import "testing", like we did with fmt in the other file

## Run docs
```go
$ godoc -http :8000
```
If you go to localhost:8000/pkg you will see all the packages installed on your system.
If you don't have godoc command, then maybe you are using the newer version of Go (1.14 or later) which is no longer including godoc. You can manually install it with go install golang.org/x/tools/cmd/godoc@latest.

### Using TDD
- Let's go over the cycle again
- Write a test
- Make the compiler pass
- Run the test, see that it fails and check the error message is meaningful
- Write enough code to make the test pass
- Refactor

Creating dictonaries / hashmaps -> Go = const
```go
const (
	french  = "French"
	spanish = "Spanish"

	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
)
```

### Returning a value from a function


Here we are receiving language as a parameter
And returning prefix as the output of this function
```go
func greetingPrefix(language string) (prefix string) {
```

### Switch case
```go
  switch language {
    case french:
      // Here we are assigning the prefix that didnt need exist before
      prefix = frenchHelloPrefix
    case spanish:
      prefix = spanishHelloPrefix
    default:
      prefix = englishHelloPrefix
  }
```

# Integers

```go
package integers

// Add takes two integers and returns the sum of them.
func Add(x, y int) int {
  return x + y
}
```

To print the integer we can do it like this:
```go
package integers

import "testing"

func TestAdder(t *testing.T) {
  sum := Add(2, 2)
  expected := 4

  if sum != expected {
    t.Errorf("expected '%d' but got '%d'", expected, sum)
  }
}
```
