package behavioral

// Memento stores the state of the Number.
type Memento struct {
	state int
}

// NewMemento creates a new memento.
func NewMemento(value int) *Memento { _ = "STUB: not implemented"; return nil }

// Number represents an integer which can be operated on.
type Number struct {
	value int
}

// NewNumber creates a new Number.
func NewNumber(value int) *Number { _ = "STUB: not implemented"; return nil }

// Dubble doubles the value of the number.
func (n *Number) Dubble() { _ = "STUB: not implemented"; return }

// Half halves the value of the number.
func (n *Number) Half() {
	_ = "STUB: not implemented"

	// Value returns the value of the number.
	return
}

func (n *Number) Value() int {
	_ = "STUB: not implemented"

	// CreateMemento creates a Memento with the current state of the number.
	return 0
}

func (n *Number) CreateMemento() *Memento { _ = "STUB: not implemented"; return nil }

// ReinstateMemento reinstates the value of the Number to the value of the memento.
func (n *Number) ReinstateMemento(memento *Memento) { _ = "STUB: not implemented"; return }
