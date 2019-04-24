package fanatick

import (
	"time"
)

// Event is a Fanatick event.
type Event struct {
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	StartAt   time.Time  `json:"start_at"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

// EventQueryParam is an event query parameter.
type EventQueryParam string

// EventGetter is the interface that wraps an event get request.
type EventGetter interface {
	GetEvent(id string) (*Event, error)
}

// EventQueryer is the interface that wraps an event query request.
type EventQueryer interface {
	QueryEvent(params map[EventQueryParam]interface{}) ([]*Event, error)
}

// EventCreator is the interface that wraps an event creation request.
type EventCreator interface {
	CreateEvent(event *Event) error
}

// EventUpdater is the interface that wraps an event update request.
type EventUpdater interface {
	UpdateEvent(event *Event) error
}

// EventDeleter is the interface that wraps an event delete request.
type EventDeleter interface {
	DeleteEvent(id string) error
}

// EventTxBeginner is the interface that wraps an event transaction starter.
type EventTxBeginner interface {
	BeginEventTx() EventTx
}

type EventTxCommitter interface {
	CommitEventTx() error
}

// EventStore defines the operations of an event store.
type EventStore interface {
	EventGetter
	EventQueryer
	EventTxBeginner
}

// EventTx defines the operations that may be performed on an event update transaction.
type EventTx interface {
	EventCreator
	EventUpdater
	EventDeleter
	EventTxCommitter
}

// The event query params.
const (
	// EventQueryParamLimit indicates the maximum number of events to return.
	EventQueryParamLimit = EventQueryParam("limit")

	// EventQueryParamBefore indicates the last event of the previously queried results.
	EventQueryParamBefore = EventQueryParam("before")
)
