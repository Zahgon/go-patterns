package creational

// Shape is an interface for interacting with a shape.
type Shape interface {
	Draw()
}

type circle struct {
}

func (c *circle) Draw() { _ = "STUB: not implemented"; return }

type square struct {
}

func (s *square) Draw() { _ = "STUB: not implemented"; return }

type ellipse struct {
}

func (e *ellipse) Draw() { _ = "STUB: not implemented"; return }

type rectangle struct {
}

func (r *rectangle) Draw() { _ = "STUB: not implemented"; return }

// ShapeFactory is an interface for a factory which can be used
// to create curved and straight shapes.
type ShapeFactory interface {
	CreateCurvedShape() Shape
	CreateStraightShape() Shape
}

type simpleShapeFactory struct {
}

// NewSimpleShapeFactory creates a new simpleShapeFactory.
func NewSimpleShapeFactory() ShapeFactory { _ = "STUB: not implemented"; return *new(ShapeFactory) }

func (s *simpleShapeFactory) CreateCurvedShape() Shape {
	_ = "STUB: not implemented"
	return *new(Shape)
}

func (s *simpleShapeFactory) CreateStraightShape() Shape {
	_ = "STUB: not implemented"
	return *new(Shape)
}

type robustShapeFactory struct {
}

// NewRobustShapeFactory creates a new robustShapeFactory.
func NewRobustShapeFactory() ShapeFactory { _ = "STUB: not implemented"; return *new(ShapeFactory) }

func (s *robustShapeFactory) CreateCurvedShape() Shape {
	_ = "STUB: not implemented"
	return *new(Shape)
}

func (s *robustShapeFactory) CreateStraightShape() Shape {
	_ = "STUB: not implemented"
	return *new(Shape)
}
