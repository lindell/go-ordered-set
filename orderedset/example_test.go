package orderedset_test

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
