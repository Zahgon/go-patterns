package structural

// CarModel is the first car subsystem which describes the car model.
type CarModel struct {
}

// NewCarModel creates a new car model.
func NewCarModel() *CarModel {
	_ = "STUB: not implemented"

	// SetModel sets the car model and logs.
	return nil
}

func (c *CarModel) SetModel() { _ = "STUB: not implemented"; return }

// CarEngine is the second car subsystem which describes the car engine.
type CarEngine struct {
}

// NewCarEngine creates a new car engine.
func NewCarEngine() *CarEngine { _ = "STUB: not implemented"; return nil }

// SetEngine sets the car engine and logs.
func (c *CarEngine) SetEngine() { _ = "STUB: not implemented"; return }

// CarBody is the third car subsystem which describes the car body.
type CarBody struct {
}

// NewCarBody creates a new car body.
func NewCarBody() *CarBody {
	_ = "STUB: not implemented"

	// SetBody sets the car body and logs.
	return nil
}

func (c *CarBody) SetBody() { _ = "STUB: not implemented"; return }

// CarAccessories is the fourth car subsystem which describes the car accessories.
type CarAccessories struct {
}

// NewCarAccessories creates new car accessories.
func NewCarAccessories() *CarAccessories { _ = "STUB: not implemented"; return nil }

// SetAccessories sets the car accessories and logs.
func (c *CarAccessories) SetAccessories() { _ = "STUB: not implemented"; return }

// CarFacade describes the car facade which provides a simplified interface to create a car.
type CarFacade struct {
	accessories *CarAccessories
	body        *CarBody
	engine      *CarEngine
	model       *CarModel
}

// NewCarFacade creates a new CarFacade.
func NewCarFacade() *CarFacade { _ = "STUB: not implemented"; return nil }

// CreateCompleteCar creates a new complete car.
func (c *CarFacade) CreateCompleteCar() { _ = "STUB: not implemented"; return }
