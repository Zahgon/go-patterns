package behavioral

// Strategy defines the interface for the strategy to execute.
type Strategy interface {
	Execute()
}

// strategyA defines an implementation of a Strategy to execute.
type strategyA struct {
}

// NewStrategyA creates a new instance of strategy A.
func NewStrategyA() Strategy {
	_ = "STUB: not implemented"
	return *

	// Execute executes strategy A.
	new(Strategy)
}

func (s *strategyA) Execute() { _ = "STUB: not implemented"; return }

// strategyB defines an implementation of a Strategy to execute.
type strategyB struct {
}

// NewStrategyB creates a new instance of strategy B.
func NewStrategyB() Strategy {
	_ = "STUB: not implemented"
	return *

	// Execute executes strategy B.
	new(Strategy)
}

func (s *strategyB) Execute() { _ = "STUB: not implemented"; return }

// Context defines a context for executing a strategy.
type Context struct {
	strategy Strategy
}

// NewContext creates a new instance of a context.
func NewContext() *Context {
	_ = "STUB: not implemented"

	// SetStrategy sets the strategy to execute for this context.
	return nil
}

func (c *Context) SetStrategy(strategy Strategy) { _ = "STUB: not implemented"; return }

// Execute executes the strategy.
func (c *Context) Execute() { _ = "STUB: not implemented"; return }
