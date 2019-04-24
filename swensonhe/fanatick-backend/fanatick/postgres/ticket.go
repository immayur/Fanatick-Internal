package postgres

import (
	"fmt"
	"github.com/segmentio/ksuid"
	"github.com/swensonhe/fanatick-backend/fanatick"
)

// EventDB is a database for events.
type TicketDB struct {
	*DB
}

// EventTx is an event transaction.
type TicketTx struct {
	*Tx
}



// BeginTx begins a transaction.
func (db *DB) BeginTicketTx() fanatick.TicketTx {
	return &TicketTx{Tx: db.Begin()}
}

func (tx *TicketTx) CommitTicketTx() error {

	err := tx.Tx.Commit()
	if err != nil {
		err := tx.Tx.Rollback()
		if err != nil {
			return fmt.Errorf("falied to rollback EventTx after commit EventTx failed")
		}
		return fmt.Errorf("falied to commit EventTx")
	}

	return nil
}


// // Create creates an event.
func (tx *TicketTx) AddTicket(ticket *fanatick.Ticket) error {
	query := `
		INSERT INTO tickets(Id,row,section,seat)
		VALUES ($1, $2, $3, $4)
	`

	ticket.ID = ksuid.New().String()
	//fmt.Print(seat)
	_, err := tx.Exec(query, ticket.ID, ticket.Row, ticket.Section, ticket.Seat)
	if err != nil {
		return err
	}

	return nil
}