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

//common respone for getevent stub and getallevents
type CommonResponseForEvent struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Venue     string `json:"venue"`
	City      string `json:"city"`
	StartAt   string `json:"start_at"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

//seat listing
type CommonResponseForSeatListing struct {
	SeatNumber string `json:"seat_number"`
}

// seat details
type AssociatedEvent struct {
	ID string `json:"id"`
}
type CommonResponseForSeatDetail struct {
	ID         string          `json:"id"`
	Event      AssociatedEvent `json:"event"`
	SeatNumber string          `json:"number"`
}

// Maker offer
type CommonResponseForOffer struct {
	Amount string                       `json:"amount"`
	Seats  CommonResponseForSeatListing `json:"seats"`
}

//commonrespone for add event
type CommonResponseForAddEvent struct {
	Status  string          `json:"status"`
	Code    string          `json:"code"`
	Message string          `json:"message"`
	Data    AssociatedEvent `json:"data"`
}

// EventQueryParam is an event query parameter.
type EventQueryParam string

// EventGetter is the interface that wraps an event get request.
type EventGetter interface {
	Get(id string) (*Event, error)
}

// EventQueryer is the interface that wraps an event query request.
type EventQueryer interface {
	Query(params map[EventQueryParam]interface{}) ([]*Event, error)
}

// EventCreator is the interface that wraps an event creation request.
type EventCreator interface {
	Create(event *Event) error
}

// EventUpdater is the interface that wraps an event update request.
type EventUpdater interface {
	Update(event *Event) error
}

// EventDeleter is the interface that wrapts an event delete request.
type EventDeleter interface {
	Delete(id string) error
}

// EventTxBeginner is the interface that wraps an event transaction starter.
type EventTxBeginner interface {
	BeginTx() EventTx
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
}

// The event query params.
const (
	// EventQueryParamLimit indicates the maximum number of events to return.
	EventQueryParamLimit = EventQueryParam("limit")

	// EventQueryParamBefore indicates the last event of the previously queried results.
	EventQueryParamBefore = EventQueryParam("before")
)
