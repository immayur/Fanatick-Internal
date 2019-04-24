package api

import (
	"github.com/swensonhe/fanatick-backend/fanatick"
)

// EventService performs operations on events.
type EventService struct {
	EventStore fanatick.EventStore
	EventTx    fanatick.EventTx
	Logger     fanatick.Logger
}

// Get returns an event.
func (svc *EventService) GetEvent(id string) (*fanatick.Event, error) {
	event, err := svc.EventStore.GetEvent(id)
	if err != nil {
		if err == fanatick.ErrorNotFound {
			return nil, ErrorNotFound(`Event not found.`)
		}

		go svc.Logger.Error(err)
		return nil, ErrorInternal(err.Error())
	}

	return event, nil
}

// Query returns events.
func (svc *EventService) QueryEvent(params map[fanatick.EventQueryParam]interface{}) ([]*fanatick.Event, error) {
	events, err := svc.EventStore.QueryEvent(params)
	if err != nil {
		go svc.Logger.Error(err)
		return nil, ErrorInternal(err.Error())
	}

	return events, nil
}

func (svc *EventService) CreateEvent(event *fanatick.Event) error {
	tx := svc.EventStore.BeginEventTx()

	err := tx.CreateEvent(event)
	if err != nil {
		if err == fanatick.ErrorInternal {
			return ErrorNotFound(`Unable to Create Event.`)
		}

		go svc.Logger.Error(err)
		return ErrorInternal(err.Error())
	}

	err = tx.CommitEventTx()
	if err != nil {
		return ErrorInternal(err.Error())
	}
	return nil
}

func (svc *EventService) DeleteEvent(id string) error {
	tx := svc.EventStore.BeginEventTx()

	err := tx.DeleteEvent(id)

	if err != nil {
		if err == fanatick.ErrorInternal {
			return ErrorNotFound(`Unable to Delete Event.`)
		}

		go svc.Logger.Error(err)
		return ErrorInternal()
	}

	err = tx.CommitEventTx()
	if err != nil {
		return ErrorInternal(err.Error())
	}

	return nil
}

func (svc *EventService) UpdateEvent(event *fanatick.Event) error {
	tx := svc.EventStore.BeginEventTx()

	err := tx.UpdateEvent(event)

	if err != nil {
		if err == fanatick.ErrorInternal {
			return ErrorNotFound(`Unable to Update Event.`)
		}

		go svc.Logger.Error(err)
		return ErrorInternal(err.Error())
	}

	err = tx.CommitEventTx()
	if err != nil {
		return ErrorInternal(err.Error())
	}

	return nil
}
