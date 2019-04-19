package api

import (
	"github.com/swensonhe/fanatick-backend/fanatick"
)

// EventService performs operations on events.
type EventService struct {
	EventStore fanatick.EventStore
	Logger     fanatick.Logger
}

// Get returns an event.
func (svc *EventService) Get(id string) (*fanatick.Event, error) {
	event, err := svc.EventStore.Get(id)
	if err != nil {
		if err == fanatick.ErrorNotFound {
			return nil, ErrorNotFound(`Event not found.`)
		}

		go svc.Logger.Error(err)
		return nil, ErrorInternal()
	}

	return event, nil
}

//returns event details based on id
func GetEventHandlerRequest(id string) (*fanatick.CommonResponseForEvent, error) {
	events := fanatick.CommonResponseForEvent{}
	events.ID = "id"
	events.Title = "abc"
	events.Venue = "bnm"
	events.City = "bhb"
	events.StartAt = ""
	events.CreatedAt = ""
	events.UpdatedAt = ""
	return &events, nil
}

//returns all events
func GetAllEventsRequest() (*fanatick.CommonResponseForEvent, error) {
	events := fanatick.CommonResponseForEvent{}
	events.ID = "id"
	events.Title = "abc"
	events.Venue = "bnm"
	events.City = "bhb"
	events.StartAt = ""
	events.CreatedAt = ""
	events.UpdatedAt = ""
	return &events, nil
}

//returns Seat Listing
func GetSeatListingRequest() (*[]fanatick.CommonResponseForSeatListing, error) {
	events := fanatick.CommonResponseForSeatListing{}
	events.SeatNumber = "23"
	var seatListing []fanatick.CommonResponseForSeatListing
	seatListing = append(seatListing, events)
	return &seatListing, nil
}

func GetSeatDetailsRequest() (*fanatick.CommonResponseForSeatDetail, error) {
	events := fanatick.CommonResponseForSeatDetail{}
	events.ID = "b"
	eventID := fanatick.AssociatedEvent{
		ID: "2",
	}
	events.Event = eventID
	events.SeatNumber = "12"
	return &events, nil
}

//returns offer request
func MakeOfferRequest() (*fanatick.CommonResponseForOffer, error) {
	events := fanatick.CommonResponseForOffer{}
	events.Amount = "50$"
	seats := fanatick.CommonResponseForSeatListing{
		SeatNumber: "12",
	}
	events.Seats = seats
	return &events, nil
}

//return status after adding event
func AddEventRequest() (*fanatick.CommonResponseForAddEvent, error) {
	events := fanatick.CommonResponseForAddEvent{}
	events.Status = "200"
	events.Code = "success"
	events.Message = "Event Successfully Added"
	dataEvent := fanatick.AssociatedEvent{
		ID: "addfagfgh",
	}
	events.Data = dataEvent
	return &events, nil
}

//return status after updating an event
func UpdateEventRequest() (*fanatick.CommonResponseForAddEvent, error) {
	events := fanatick.CommonResponseForAddEvent{}
	events.Status = "200"
	events.Code = "success"
	events.Message = "Event Successfully Updated"
	dataEvent := fanatick.AssociatedEvent{
		ID: "addfagfgh",
	}
	events.Data = dataEvent
	return &events, nil
}

//return status after deleting an event
func DeleteEventRequest() (*fanatick.CommonResponseForAddEvent, error) {
	events := fanatick.CommonResponseForAddEvent{}
	events.Status = "200"
	events.Code = "success"
	events.Message = "Event Successfully Deleted"
	dataEvent := fanatick.AssociatedEvent{
		ID: "addfagfgh",
	}
	events.Data = dataEvent
	return &events, nil
}

// Query returns events.
func (svc *EventService) Query(params map[fanatick.EventQueryParam]interface{}) ([]*fanatick.Event, error) {
	events, err := svc.EventStore.Query(params)
	if err != nil {
		go svc.Logger.Error(err)
		return nil, ErrorInternal()
	}

	return events, nil
}
