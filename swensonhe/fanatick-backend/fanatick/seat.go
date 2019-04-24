package fanatick

import (
	"time"
)

// Seat is a Fanatick event.
type Seat struct {
	ID        string     `json:"id"`
	No        int        `json:"no"`
	Row       string     `json:"row"`
	Section   int        `json:"section"`
	SeatURL   string     `json:"image_url"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

// EventQueryParam is an event query parameter.
type SeatQueryParam string

// EventGetter is the interface that wraps an event get request.
type SeatGetter interface {
	GetSeatByID(id string) (*Seat, error)
}

// EventQueryer is the interface that wraps an event query request.
type SeatQueryer interface {
	QuerySeatsbyEventID(id string, params map[SeatQueryParam]interface{}) ([]*Seat, error)
}

// EventCreator is the interface that wraps an event creation request.
type SeatCreator interface {
	CreateSeat(eventId string, seat *Seat) error
}

// EventUpdater is the interface that wraps an event update request.
type SeatUpdater interface {
	UpdateSeat(seat *Seat) error
}

// EventDeleter is the interface that wrapts an event delete request.
type SeatDeleter interface {
	DeleteSeat(id string) error
}

// EventTxBeginner is the interface that wraps an event transaction starter.
type SeatTxBeginner interface {
	BeginSeatTx() SeatTx
}

type SeatTxCommitter interface {
	CommitSeatTx() error
}

// EventStore defines the operations of an event store.
type SeatStore interface {
	SeatGetter
	SeatQueryer
	SeatTxBeginner
}

// EventTx defines the operations that may be performed on an event update transaction.
type SeatTx interface {
	SeatCreator
	SeatUpdater
	SeatDeleter
	SeatTxCommitter
}

// The event query params.
const (
	// EventQueryParamLimit indicates the maximum number of events to return.
	SeatQueryParamLimit = SeatQueryParam("limit")

	// EventQueryParamBefore indicates the last event of the previously queried results.
	SeatQueryParamBefore = SeatQueryParam("before")
)
