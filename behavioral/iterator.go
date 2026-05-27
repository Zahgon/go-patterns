package behavioral

// Iterator is an interface for an iterator.
type Iterator interface {

	// Index returns the index of the current iterator.
	Index() int

	// Value returns the current value of the iterator.
	Value() interface{}

	// HasNext returns whether another next element exists.
	HasNext() bool

	// Next increments the iterator to point to the next element.
	Next()
}

// ArrayIterator is an iterator which iterates over an array.
type ArrayIterator struct {
	array []interface{}
	index int
}

// Index returns the index of the current iterator.
func (i *ArrayIterator) Index() int {
	_ = "STUB: not implemented"

	// Value returns the current value of the iterator.
	return 0
}

func (i *ArrayIterator) Value() interface{} { _ = "STUB: not implemented"; return nil }

// HasNext returns whether another next element exists.
func (i *ArrayIterator) HasNext() bool { _ = "STUB: not implemented"; return false }

// Next increments the iterator to point to the next element.
func (i *ArrayIterator) Next() { _ = "STUB: not implemented"; return }
