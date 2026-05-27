package behavioral

import (
	"io"
	"os"
)

var outputWriter io.Writer = os.Stdout // modified during testing

// Event describes an event to observe and notify on.
type Event struct {
	id string
}

// EventObserver describes an interface for observing events.
type EventObserver interface {
	OnNotify(event Event)
}

// observer is an implementation of the EventObserver interface.
type observer struct {
	name string
}

// NewEventObserver returns a new instance of an EventObserver.
func NewEventObserver(name string) EventObserver {
	_ = "STUB: not implemented"
	return *

	// OnNotify logs the event being notified on.
	new(EventObserver)
}

func (o *observer) OnNotify(event Event) { _ = "STUB: not implemented"; return }

// EventNotifier describes an interface for registering and de-registering observers to
// be notified when an event occurs.
type EventNotifier interface {
	Register(obs EventObserver)
	Deregister(obs EventObserver)
	Notify(event Event)
}

// eventNotifer is an implementation of the EventNotifier interface.
type eventNotifer struct {
	observers []EventObserver
}

// NewEventNotifier returns a new instance of an EventNotifier.
func NewEventNotifier() EventNotifier {
	_ = "STUB: not implemented"
	return *

	// Register registers a new observer for notifying on.
	new(EventNotifier)
}

func (e *eventNotifer) Register(obs EventObserver) { _ = "STUB: not implemented"; return }

// Deregister de-registers an observer for notifying on.
func (e *eventNotifer) Deregister(obs EventObserver) { _ = "STUB: not implemented"; return }

// Notify notifies all observers on an event.
func (e *eventNotifer) Notify(event Event) { _ = "STUB: not implemented"; return }
