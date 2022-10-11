package orderedset

type listElement[T comparable] struct {
	next, prev *listElement[T]

	value T
}

// insertAfter inserts the new element after the current one.
func (elem *listElement[T]) insertAfter(newElem *listElement[T]) {
	next := elem.next
	elem.next = newElem
	next.prev = newElem
	newElem.next = next
	newElem.prev = elem
}

// insertBefore inserts the new element before the current one.
func (elem *listElement[T]) insertBefore(newElem *listElement[T]) {
	prev := elem.prev
	elem.prev = newElem
	prev.next = newElem
	newElem.next = elem
	newElem.prev = prev
}

// remove the element from the list, don't use the element after this method is called.
func (elem *listElement[T]) remove() {
	elem.prev.next = elem.next
	elem.next.prev = elem.prev
}
