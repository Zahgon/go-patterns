package behavioral

// WildStallion describes an interface for a Wild Stallion band member.
type WildStallion interface {
	SetMediator(mediator Mediator)
}

// Bill describes Bill S. Preston, Esquire.
type Bill struct {
	mediator Mediator
}

// SetMediator sets the mediator.
func (b *Bill) SetMediator(mediator Mediator) { _ = "STUB: not implemented"; return }

// Respond responds.
func (b *Bill) Respond() { _ = "STUB: not implemented"; return }

// Ted describes Ted "Theodore" Logan.
type Ted struct {
	mediator Mediator
}

// SetMediator sets the mediator.
func (t *Ted) SetMediator(mediator Mediator) { _ = "STUB: not implemented"; return }

// Talk talks through mediator.
func (t *Ted) Talk() { _ = "STUB: not implemented"; return }

// Respond responds.
func (t *Ted) Respond() { _ = "STUB: not implemented"; return }

// Mediator describes the interface for communicating between Wild Stallion band members.
type Mediator interface {
	Communicate(who string)
}

// ConcreateMediator describes a mediator between Bill and Ted.
type ConcreateMediator struct {
	Bill
	Ted
}

// NewMediator creates a new ConcreateMediator.
func NewMediator() *ConcreateMediator { _ = "STUB: not implemented"; return nil }

// Communicate communicates between Bill and Ted.
func (m *ConcreateMediator) Communicate(who string) { _ = "STUB: not implemented"; return }
