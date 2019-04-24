package api

import (
	//"fmt"

	"github.com/swensonhe/fanatick-backend/fanatick"
)

// EventService performs operations on events.
type TicketService struct {
	TicketStore fanatick.TicketStore
	Logger    fanatick.Logger
}


func (svc *TicketService) AddTicket( ticket *fanatick.Ticket) error {
	tx := svc.TicketStore.BeginTicketTx()

	err := tx.AddTicket(ticket)

	if err != nil {
		if err == fanatick.ErrorInternal {
			return ErrorNotFound(`Unable to Create Event.`)
		}

		go svc.Logger.Error(err)
		return ErrorInternal()
	}

	err = tx.CommitTicketTx()
	//fmt.Println(err)
	return nil
}