package behavioral

// Command is a command which when executed on a person will call a given method.
type Command struct {
	person *Person
	method func()
}

// NewCommand creates a new Command.
func NewCommand(person *Person, method func()) Command {
	_ = "STUB: not implemented"
	return *new(Command)
}

// Execute executes the command.
func (c *Command) Execute() {
	_ = "STUB: not implemented"

	// Person with a given name and command to execute.
	return
}

type Person struct {
	name string
	cmd  Command
}

// NewPerson creates a new Person.
func NewPerson(name string, cmd Command) Person {
	_ = "STUB: not implemented"
	return *

	// Talk talks and the executes the follow on command.
	new(Person)
}

func (p *Person) Talk() { _ = "STUB: not implemented"; return }

// PassOn passes on by executing the follow on command.
func (p *Person) PassOn() { _ = "STUB: not implemented"; return }

// Gossip gossips and then executes follow on command.
func (p *Person) Gossip() { _ = "STUB: not implemented"; return }

// Listen listens without executing follow on command.
func (p *Person) Listen() { _ = "STUB: not implemented"; return }
