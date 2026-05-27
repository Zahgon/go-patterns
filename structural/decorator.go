package structural

// TaskFunc represents a function to perform a task.
type TaskFunc func(int) int

// LogDecorate decorates a task functions execution with some logging.
func LogDecorate(fn TaskFunc) TaskFunc { _ = "STUB: not implemented"; return *new(TaskFunc) }
