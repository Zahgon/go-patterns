package behavioral

// Expression represents an expression to evaluate.
type Expression interface {
	Interpret(variables map[string]Expression) int
}

// Integer represents an integer number.
type Integer struct {
	integer int
}

// Interpret returns the integer representation of the number.
func (n *Integer) Interpret(variables map[string]Expression) int {
	_ = "STUB: not implemented"

	// Plus represents the addition operation.
	return 0
}

type Plus struct {
	leftOperand  Expression
	rightOperand Expression
}

// Interpret interprets by adding the left and right variables.
func (p *Plus) Interpret(variables map[string]Expression) int { _ = "STUB: not implemented"; return 0 }

// Minus represents the subtraction operation.
type Minus struct {
	leftOperand  Expression
	rightOperand Expression
}

// Interpret interprets by subtracting the right from left variables.
func (m *Minus) Interpret(variables map[string]Expression) int { _ = "STUB: not implemented"; return 0 }

// Variable represents a variable.
type Variable struct {
	name string
}

// Interpret looks up the variable value and returns it, if not found returns zero.
func (v *Variable) Interpret(variables map[string]Expression) int {
	_ = "STUB: not implemented"
	return 0
}

// Evaluator evaluates the expression.
type Evaluator struct {
	syntaxTree Expression
}

// NewEvaluator creates a new Evaluator.
func NewEvaluator(expression string) *Evaluator { _ = "STUB: not implemented"; return nil }

// Interpret interprets the expression syntax tree.
func (e *Evaluator) Interpret(context map[string]Expression) int {
	_ = "STUB: not implemented"
	return 0
}

// Node represents a node in the stack.
type Node struct {
	value interface{}
	next  *Node
}

// Stack represents a stack with push and pop operations.
type Stack struct {
	top  *Node
	size int
}

// Push pushes a new value into the stack.
func (s *Stack) Push(value interface{}) { _ = "STUB: not implemented"; return }

// Pop pops a value out the stack.
func (s *Stack) Pop() interface{} { _ = "STUB: not implemented"; return nil }
