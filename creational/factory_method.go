package creational

import (
	"io"
	"os"
)

var outputWriter io.Writer = os.Stdout // modified during testing

// StoogeType is used as a enum for stooge types.
type StoogeType int

// Names of stooges as enums.
const (
	Larry StoogeType = iota
	Moe
	Curly
)

// Stooge provides an interface for interacting with stooges.
type Stooge interface {
	SlapStick()
}

type larry struct {
}

func (s *larry) SlapStick() { _ = "STUB: not implemented"; return }

type moe struct {
}

func (s *moe) SlapStick() { _ = "STUB: not implemented"; return }

type curly struct {
}

func (s *curly) SlapStick() { _ = "STUB: not implemented"; return }

// NewStooge creates new stooges given the stooge type.
// Nil is returned if the stooge type is not recognised.
func NewStooge(stooge StoogeType) Stooge { _ = "STUB: not implemented"; return *new(Stooge) }
