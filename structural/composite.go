package structural

// Component describes the behavior that needs to be exercised uniformly
// across all primitive and composite objects.
type Component interface {
	Traverse()
}

// Leaf describes a primitive leaf object in the hierarchy.
type Leaf struct {
	value int
}

// NewLeaf creates a new leaf.
func NewLeaf(value int) *Leaf { _ = "STUB: not implemented"; return nil }

// Traverse prints the value of the leaf.
func (l *Leaf) Traverse() { _ = "STUB: not implemented"; return }

// Composite describes a composite of components.
type Composite struct {
	children []Component
}

// NewComposite creates a new composite.
func NewComposite() *Composite { _ = "STUB: not implemented"; return nil }

// Add adds a new component to the composite.
func (c *Composite) Add(component Component) { _ = "STUB: not implemented"; return }

// Traverse traverses the composites children.
func (c *Composite) Traverse() { _ = "STUB: not implemented"; return }
