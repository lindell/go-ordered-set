[![Go Reference](https://pkg.go.dev/badge/github.com/lindell/go-ordered-set/orderedset.svg)](https://pkg.go.dev/github.com/lindell/go-ordered-set/orderedset)
[![Coverage Status](https://coveralls.io/repos/github/lindell/go-ordered-set/badge.svg)](https://coveralls.io/github/lindell/go-ordered-set)


# go-ordered-set
Insertion ordered set for Go using generics.

All operations are running in constant time (`O(1)`).

[The full documentation is available on pkg.go.dev.](https://pkg.go.dev/github.com/lindell/go-ordered-set/orderedset)

# Install

```sh
go get github.com/lindell/go-ordered-set
```

Import:
```go
import "github.com/lindell/go-ordered-set/orderedset"
```

# Example:

```go
import (
	"fmt"

	"github.com/lindell/go-ordered-set/orderedset"
)

func Example() {
	set := orderedset.New[string]()
	set.Add("A")
	set.Add("B")
	set.Add("C")
	set.Add("D")
	set.AddFirst("E")

	deleted := set.Delete("B")
	fmt.Printf("B deleted: %v\n", deleted)

	fmt.Printf("Set contains A: %v\n", set.Contains("A"))
	fmt.Printf("Set contains B: %v\n", set.Contains("B"))

	fmt.Printf("Size: %v\n", set.Size())

	fmt.Printf("Values: %v\n", set.Values())

	fmt.Println("Iterated Values:")
	it := set.Iter()
	for val, ok := it.Next(); ok; val, ok = it.Next() {
		fmt.Printf(" %v\n", val)
	}

	// Output:
	// B deleted: true
	// Set contains A: true
	// Set contains B: false
	// Size: 4
	// Values: [E A C D]
	// Iterated Values:
	//  E
	//  A
	//  C
	//  D
}
```
