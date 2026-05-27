package behavioral

// WorkerInterface defines an interface for a worker.
type WorkerInterface interface {
	GetUp()
	EatBreakfast()
	GoToWork()
	Work()
	ReturnHome()
	Relax()
	Sleep()
}

// Worker defines the worker.
type Worker struct {
	WorkerInterface
}

// NewWorker returns a new Worker.
func NewWorker(w WorkerInterface) *Worker {
	_ = "STUB: not implemented"

	// DailyRoutine is the template method for printing the workers daily routine.
	return nil
}

func (w *Worker) DailyRoutine() { _ = "STUB: not implemented"; return }

// PostMan is a worker.
type PostMan struct {
}

// GetUp prints what the postman does to get up.
func (w *PostMan) GetUp() { _ = "STUB: not implemented"; return }

// EatBreakfast prints what the postman does to eat breakfast.
func (w *PostMan) EatBreakfast() { _ = "STUB: not implemented"; return }

// GoToWork prints what the postman does to get to work.
func (w *PostMan) GoToWork() { _ = "STUB: not implemented"; return }

// Work prints what the postman does to work.
func (w *PostMan) Work() { _ = "STUB: not implemented"; return }

// ReturnHome prints what the postman does to get home.
func (w *PostMan) ReturnHome() { _ = "STUB: not implemented"; return }

// Relax prints what the postman does to relax.
func (w *PostMan) Relax() { _ = "STUB: not implemented"; return }

// Sleep prints what the postman does to sleep.
func (w *PostMan) Sleep() { _ = "STUB: not implemented"; return }
