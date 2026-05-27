package behavioral

// Machine defines a machine which can be swwitched on and off.
type Machine struct {
	current State
}

// NewMachine creates a new machine.
func NewMachine() *Machine { _ = "STUB: not implemented"; return nil }

// setCurrent sets the current state of the machine.
func (m *Machine) setCurrent(s State) {
	_ = "STUB: not implemented"

	// On pushes the on button.
	return
}

func (m *Machine) On() {
	_ = "STUB: not implemented"

	// Off pushes the off button.
	return
}

func (m *Machine) Off() {
	_ = "STUB: not implemented"

	// State describes the internal state of the machine.
	return
}

type State interface {
	On(m *Machine)
	Off(m *Machine)
}

// ON describes the on button state.
type ON struct {
}

// NewON creates a new ON state.
func NewON() State {
	_ = "STUB: not implemented"

	// On does nothing.
	return *new(State)
}

func (o *ON) On(m *Machine) { _ = "STUB: not implemented"; return }

// Off switches the state from on to off.
func (o *ON) Off(m *Machine) { _ = "STUB: not implemented"; return }

// OFF describes the off button state.
type OFF struct {
}

// NewOFF creates a new OFF state.
func NewOFF() State {
	_ = "STUB: not implemented"

	// On switches the state from off to on.
	return *new(State)
}

func (o *OFF) On(m *Machine) { _ = "STUB: not implemented"; return }

// Off does nothing.
func (o *OFF) Off(m *Machine) { _ = "STUB: not implemented"; return }
