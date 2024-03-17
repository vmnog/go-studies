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

# Loop Iterations
* Go only has only "for", it does not have while,until, etc...

```go
package iteration

const repeatCount = 5

func Repeat(character string) string {
  var repeated string

  for i:= 0; i < repeatCount; i++ {
    repeated += character
  }

  return repeated
}
```

OBS:
When you declare a function in a folder it automatically creates a package with the folder names
Then, if you create another file with the package import, you have access to all functions declared in this folder
** It does not apply for subfolders, if you want to use functions from other folders you need to import the module 


You can use range to get the array size in for loop

"For in Range" Example
```go
func Sum(numbers [5]int) int {
  sum := 0
  for _, number := range numbers {
    sum += number
  }
  return sum
}
```

To run tests with coverage analysis:
```bash
➜  sum git:(main) ✗ go test -cover
PASS
coverage: 100.0% of statements
ok  	hello/sum	0.215s
```

```go
package main

import (
	"reflect"
	"testing"
)

func TestSumAll(t *testing.T) {
  got := SumAll([]int{1, 2}, []int{0, 9})
  want := []int{3, 9}

  // It's important to note that reflect.DeepEqual
  // is not "type safe" - the code will compile even
  // if you did something a bit silly. 
  if !reflect.DeepEqual(got, want) {
    t.Errorf("got %v, want %v", got, want)
  }
}

```

improving the function...

```go
// We need a new function called SumAll
// which will take a varying number of slices,
// returning a new slice containing the totals
// for each slice passed in
func SumAll(numbersToSum ...[]int) []int {
  lengthOfNumbers := len(numbersToSum)
  sums := make([]int, lengthOfNumbers)
  for i, numbers := range numbersToSum {
    sums[i] = Sum(numbers)
    fmt.Printf("numbers: %d, sums: %d\n", numbers, sums)
  }
  return sums
}

// output
// numbers: [1 2], sums: [3 0]
// numbers: [0 9], sums: [3 9]
// PASS
// ok  	hello/sum	0.514s
```

Refactoring this SumAll  function

```go
package main

func SumAll(numbersToSum ...[]int) []int {
  var sums []int
  for _, numbers := range numbersToSum {
    sums = append(sums, Sum(numbers))
  }
  return sums
}
```

Two ways of getting the tail (last item) from array

```go
package main

func SumAllTails(tailsToSum... []int) []int {
  var sums []int 
  for _, numbers := range tailsToSum {
    // sums = append(sums, Sum(numbers[len(numbers) - 1]))
    // Slices can be sliced! The syntax is slice[low:high]
    sums = append(sums, Sum(numbers[1:]))
  }
  return sums
}
```

Creating testing helper functions to avoid code repeat
```go
func TestSumAllTails(t *testing.T) {
      // Helper func
      checkSums := func (t testing.TB, got, want []int) {
        t.Helper()
        if !reflect.DeepEqual(got, want) {
                t.Errorf("got %v want %v", got, want)
        }
      }

	t.Run("make the sums of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}
        // Helper func
        checkSums(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
        // Helper func
        checkSums(t, got, want)
		want := []int{0, 9}
	})
}

```
