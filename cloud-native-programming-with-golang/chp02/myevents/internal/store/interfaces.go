package store

// EventReadWriter interface spcifies Event behaviors
type EventReadWriter interface {
	AddEvent(Event) ([]byte, error)
	FindEvent([]byte) (Event, error)
	FindEventByName(string) (Event, error)
	FindAllEvents() ([]Event, error)
}
