package api

import (
	"fmt"

	"github.com/swensonhe/fanatick-backend/fanatick"
)

// EventService performs operations on events.
type SeatService struct {
	SeatStore fanatick.SeatStore
	Logger    fanatick.Logger
}

// Get returns an event.
func (svc *SeatService) GetSeatByID(id string) (*fanatick.Seat, error) {
	event, err := svc.SeatStore.GetSeatByID(id)
	if err != nil {
		if err == fanatick.ErrorNotFound {
			return nil, ErrorNotFound(`Seat not found.`)
		}

		go svc.Logger.Error(err)
		return nil, ErrorInternal()
	}

	return event, nil
}

// Query returns events.
func (svc *SeatService) QuerySeatsbyEventID(id string, params map[fanatick.SeatQueryParam]interface{}) ([]*fanatick.Seat, error) {
	events, err := svc.SeatStore.QuerySeatsbyEventID(id, params)
	if err != nil {
		go svc.Logger.Error(err)
		return nil, ErrorInternal()
	}

	return events, nil
}

func (svc *SeatService) CreateSeat(id string, seat *fanatick.Seat) error {
	tx := svc.SeatStore.BeginSeatTx()

	err := tx.CreateSeat(id, seat)

	if err != nil {
		if err == fanatick.ErrorInternal {
			return ErrorNotFound(`Unable to Create Event.`)
		}

		go svc.Logger.Error(err)
		return ErrorInternal()
	}

	err = tx.CommitSeatTx()
	fmt.Println(err)
	return nil
}
