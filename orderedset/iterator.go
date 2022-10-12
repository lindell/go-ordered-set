// Package orderedset implements an ordered set, i.e. a set that also keeps track of
// the order in which values were inserted.
//
// All operations are constant-time (except getting a values that are O(n)).

package orderedset

// Iterator is used to iterate over an OrderedSet.
type Iterator[T comparable] struct {
	reversed bool
	current  *listElement[T]
}

// Next returns the next value in the iteration if there is one,
// and reports whether the returned value is valid.
// Once Next returns ok==false, the iteration is over,
// and all subsequent calls will return ok==false.
func (i *Iterator[T]) Next() (val T, ok bool) {
	if i.reversed {
		current := i.current
		if current.prev == nil {
			var ret T
			return ret, false
		}
		i.current = current.prev

		return current.value, true
	}

	current := i.current
	if current.next == nil {
		var ret T
		return ret, false
	}
	i.current = current.next

	return current.value, true
}
