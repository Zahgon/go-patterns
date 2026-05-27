package structural

// TimeImp is the interface for different implementations of telling the time.
type TimeImp interface {
	Tell()
}

// BasicTimeImp defines a TimeImp which tells the time in 24 hour format.
type BasicTimeImp struct {
	hour   int
	minute int
}

// NewBasicTimeImp creates a new TimeImp.
func NewBasicTimeImp(hour, minute int) TimeImp { _ = "STUB: not implemented"; return *new(TimeImp) }

// Tell tells the time in 24 hour format.
func (t *BasicTimeImp) Tell() { _ = "STUB: not implemented"; return }

// CivilianTimeImp defines a TimeImp which tells the time in 12 hour format with meridem.
type CivilianTimeImp struct {
	BasicTimeImp
	meridiem string
}

// NewCivilianTimeImp creates a new TimeImp.
func NewCivilianTimeImp(hour, minute int, pm bool) TimeImp {
	_ = "STUB: not implemented"
	return *new(TimeImp)
}

// Tell tells the time in 12 hour format.
func (t *CivilianTimeImp) Tell() { _ = "STUB: not implemented"; return }

// ZuluTimeImp defines a TimeImp which tells the time in Zulu zone format.
type ZuluTimeImp struct {
	BasicTimeImp
	zone string
}

// NewZuluTimeImp creates a new TimeImp.
func NewZuluTimeImp(hour, minute, zoneID int) TimeImp {
	_ = "STUB: not implemented"
	return *new(TimeImp)
}

// Tell tells the time in 24 hour format in Zulu time zone.
func (t *ZuluTimeImp) Tell() { _ = "STUB: not implemented"; return }

// Time is the base struct for Time containing the implementation.
type Time struct {
	imp TimeImp
}

// NewTime creates a new time which uses a basic time for its implementation.
func NewTime(hour, minute int) *Time { _ = "STUB: not implemented"; return nil }

// Tell tells the time with the given implementation.
func (t *Time) Tell() {
	_ = "STUB: not implemented"

	// CivilianTime is a time with a civilian time implementation.
	return
}

type CivilianTime struct {
	Time
}

// NewCivilianTime creates a new civilian time.
func NewCivilianTime(hour, minute int, pm bool) *Time { _ = "STUB: not implemented"; return nil }

// ZuluTime is a time with a Zulu time implementation.
type ZuluTime struct {
	Time
}

// NewZuluTime creates a new Zulu time.
func NewZuluTime(hour, minute, zoneID int) *Time { _ = "STUB: not implemented"; return nil }
