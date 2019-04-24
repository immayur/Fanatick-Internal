package fanatick

// Ticket is a Fanatick event ticket.
type Ticket struct {
	ID      string `json:"ticket_id"`
	Section string `json:"section"`
	Row     string `json:"row"`
	Seat    string `json:"seat"`
}

// TicketQueryParam is an event query parameter.
type TicketQueryParam string

/*// EventGetter is the interface that wraps an event get request.
type TicketGetter interface {
	GetTicketByID(id string) (*Seat, error)
}

// EventQueryer is the interface that wraps an event query request.
type TicketQueryer interface {
	QueryTicketsbyEventID(id string, params map[SeatQueryParam]interface{}) ([]*Seat, error)
}
*/

// TicketAdder is the interface that wraps an AddTicket request.
type TicketAdder interface {
	AddTicket(ticket *Ticket) error
}


// EventTxBeginner is the interface that wraps an event transaction starter.
type TicketTxBeginner interface {
	BeginTicketTx() TicketTx
}

type TicketTxCommitter interface {
	CommitTicketTx() error
}


// EventStore defines the operations of an event store.
type TicketStore interface {
	//TicketGetter
	//TicketQueryer
	TicketTxBeginner
}

// EventTx defines the operations that may be performed on an event update transaction.
type TicketTx interface {
	TicketAdder
	//UpdateTicket
	//DeleteTicket
	TicketTxCommitter
}