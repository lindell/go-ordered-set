package orderedset

type Iterator[T comparable] struct {
	reversed bool
	current  *listElement[T]
}

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
