package orderedset_test

import (
	"reflect"
	"testing"

	"github.com/lindell/go-ordered-set/orderedset"
)

func TestSimple(t *testing.T) {
	set := orderedset.New[string]()
	set.Add("A")
	set.Add("B")
	set.Add("C")
	set.AddFirst("E")
	set.Add("D")

	if deleted := set.Delete("B"); !deleted {
		t.Error("value should have been deleted")
	}

	if deleted := set.Delete("B"); deleted {
		t.Error("value should not have been deleted")
	}

	if !set.Contains("A") {
		t.Error("set should contain A")
	}

	if set.Contains("B") {
		t.Error("set should not contain deleted value B")
	}

	if set.Size() != 4 {
		t.Error("size is not correct")
	}

	checkValues(t, set, []string{"E", "A", "C", "D"})
}

func TestAddValuesAgain(t *testing.T) {
	set := orderedset.New[string]()
	set.Add("A")
	set.Add("B")
	set.Add("C")
	set.Add("D")
	set.Delete("A")
	set.Delete("C")
	set.Add("D")
	set.Add("C")
	set.Add("B")
	set.Add("A")
	set.AddFirst("A")

	checkValues(t, set, []string{"B", "D", "C", "A"})
}

func checkValues[T comparable](t *testing.T, set orderedset.OrderedSet[T], values []T) {
	t.Helper()

	// Make sure Values return correctly
	if !reflect.DeepEqual(set.Values(), values) {
		t.Errorf("values from Values() should be %v, was %v", values, set.Values())
	}

	// Verify the values returned by iterate by appending to a slice and then compare
	result := []T{}
	iterator := set.Iter()
	for val, ok := iterator.Next(); ok; val, ok = iterator.Next() {
		result = append(result, val)
	}
	if !reflect.DeepEqual(result, values) {
		t.Errorf("values from Values() should be %v, was %v", values, result)
	}

	// Verify the values returned by reverse iterate by appending to a slice and then compare
	reverseResult := []T{}
	reverseIterator := set.IterReverse()
	for val, ok := reverseIterator.Next(); ok; val, ok = reverseIterator.Next() {
		reverseResult = append(reverseResult, val)
	}
	if !reflect.DeepEqual(reverseResult, reverse(values)) {
		t.Errorf("values from Values() should be %v, was %v", reverse(values), reverseResult)
	}
}

func reverse[T any](s []T) []T {
	ret := make([]T, len(s))
	for i, j := 0, len(s)-1; i < len(s); i, j = i+1, j-1 {
		ret[i] = s[j]
	}
	return ret
}
