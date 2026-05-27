package creational

/*
	Example of builder pattern:

	builder := NewConcreteBuilder()
	director := NewDirector(builder)
	director.Construct()
	product := builder.GetResult()
*/

// Director is the object which orchestrates the building of a product.
type Director struct {
	builder Builder
}

// NewDirector creates a new Director with a specified Builder.
func NewDirector(builder Builder) Director {
	_ = "STUB: not implemented"
	return *

	// Construct builds the product from a series of steps.
	new(Director)
}

func (d *Director) Construct() {
	_ = "STUB: not implemented"

	// Builder is an interface for building.
	return
}

type Builder interface {
	Build()
}

// ConcreteBuilder is a builder for building a Product
type ConcreteBuilder struct {
	built bool
}

// NewConcreteBuilder returns a new Builder.
func NewConcreteBuilder() ConcreteBuilder { _ = "STUB: not implemented"; return *new(ConcreteBuilder) }

// Build builds the product.
func (b *ConcreteBuilder) Build() {
	_ = "STUB: not implemented"

	// GetResult returns the Product which has been build during the Build step.
	return
}

func (b *ConcreteBuilder) GetResult() Product {
	_ = "STUB: not implemented"
	return *

	// Product describes the product to be built.
	new(Product)
}

type Product struct {
	Built bool
}
