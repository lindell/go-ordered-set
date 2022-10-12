package orderedset

// OrderedSet is a set where the order of elements does kept.
type OrderedSet[T comparable] struct {
	lookup map[T]*listElement[T]

	first *listElement[T]
	last  *listElement[T]
}

// New creates a new ordered set.
func New[T comparable]() *OrderedSet[T] {
	first := &listElement[T]{}
	last := &listElement[T]{}
	first.next = last
	last.prev = first

	return &OrderedSet[T]{
		lookup: map[T]*listElement[T]{},

		first: first,
		last:  last,
	}
}

// Add a new value to the set (at the end).
// If the item already exist, the placement of the item is kept.
func (os *OrderedSet[T]) Add(val T) {
	if _, exist := os.lookup[val]; exist {
		return
	}

	element := &listElement[T]{
		value: val,
	}
	os.last.insertBefore(element)
	os.lookup[val] = element
}

// AddFirst adds the items at the beginning.
func (os *OrderedSet[T]) AddFirst(val T) {
	if _, exist := os.lookup[val]; exist {
		return
	}

	element := &listElement[T]{
		value: val,
	}
	os.first.insertAfter(element)
	os.lookup[val] = element
}

// Delete an item from the set.
func (os *OrderedSet[T]) Delete(val T) (deleted bool) {
	element, ok := os.lookup[val]
	if !ok {
		return false
	}

	delete(os.lookup, val)
	element.remove()

	return true
}

// Contains checks if the set contains the value.
func (os *OrderedSet[T]) Contains(val T) bool {
	_, exist := os.lookup[val]

	return exist
}

// Size returns the size of the set-.
func (os *OrderedSet[T]) Size() int {
	return len(os.lookup)
}

// Iter returns an iterator that can be used to iterate over all values.
// Consistent iteration is not guaranteed if the set is changes during iteration.
func (os *OrderedSet[T]) Iter() *Iterator[T] {
	return &Iterator[T]{
		current: os.first.next,
	}
}

// IterReverse returns an iterator that can be used to iterate over all values from the end to the beginning.
// Consistent iteration is not guaranteed if the set is changes during iteration.
func (os *OrderedSet[T]) IterReverse() *Iterator[T] {
	return &Iterator[T]{
		reversed: true,
		current:  os.last.prev,
	}
}

// Values returns all values of the set in the correct order.
func (os *OrderedSet[T]) Values() []T {
	values := make([]T, len(os.lookup))
	i := 0
	for current := os.first.next; current.next != nil; current = current.next {
		values[i] = current.value
		i++
	}

	return values
}
