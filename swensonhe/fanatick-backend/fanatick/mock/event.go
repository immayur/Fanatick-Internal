package mock

import (
	"github.com/swensonhe/fanatick-backend/fanatick"
)

// EventGetter is a mock event getter.
type EventGetter struct {
	Invoked bool
	Fn      func(id string) (*fanatick.Event, error)
}

// NewEventGetter returns a new mock event getter.
func NewEventGetter() *EventGetter {
	return &EventGetter{
		Fn: func(id string) (*fanatick.Event, error) {
			return &fanatick.Event{}, nil
		},
	}
}

// Get invokes the function.
func (g *EventGetter) Get(id string) (*fanatick.Event, error) {
	g.Invoked = true
	return g.Fn(id)
}
