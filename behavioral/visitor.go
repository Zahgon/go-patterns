package behavioral

// Element defines an interface for accepting visitors.
type Element interface {
	Accept(v Visitor)
}

// This defines a struct which is an element.
type This struct {
}

// This returns 'This' as a string.
func (t *This) This() string {
	_ = "STUB: not implemented"

	// Accept accepts a visitor.
	return ""
}

func (t *This) Accept(v Visitor) {
	_ = "STUB: not implemented"

	// That defines a struct which is an element.
	return
}

type That struct {
}

// That returns 'That' as a string.
func (t *That) That() string {
	_ = "STUB: not implemented"

	// Accept accepts a visitor.
	return ""
}

func (t *That) Accept(v Visitor) {
	_ = "STUB: not implemented"

	// TheOther defines a struct which is an element.
	return
}

type TheOther struct {
}

// TheOther returns 'TheOther' as a string.
func (t *TheOther) TheOther() string {
	_ = "STUB: not implemented"

	// Accept accepts a visitor.
	return ""
}

func (t *TheOther) Accept(v Visitor) {
	_ = "STUB: not implemented"

	// Visitor defines an interface for visiting this, that and the other.
	return
}

type Visitor interface {
	VisitThis(e *This)
	VisitThat(e *That)
	VisitTheOther(e *TheOther)
}

// UpVisitor defines an up visitor.
type UpVisitor struct {
}

// VisitThis visits this.
func (v *UpVisitor) VisitThis(e *This) { _ = "STUB: not implemented"; return }

// VisitThat visits that.
func (v *UpVisitor) VisitThat(e *That) { _ = "STUB: not implemented"; return }

// VisitTheOther visits the other.
func (v *UpVisitor) VisitTheOther(e *TheOther) { _ = "STUB: not implemented"; return }

// DownVisitor defines a down visitor.
type DownVisitor struct {
}

// VisitThis visits this.
func (v *DownVisitor) VisitThis(e *This) { _ = "STUB: not implemented"; return }

// VisitThat visits that.
func (v *DownVisitor) VisitThat(e *That) { _ = "STUB: not implemented"; return }

// VisitTheOther visits the other.
func (v *DownVisitor) VisitTheOther(e *TheOther) { _ = "STUB: not implemented"; return }
