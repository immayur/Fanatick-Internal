package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"

	"github.com/swensonhe/fanatick-backend/fanatick"
)

// GetEvent godoc
// @Summary Show a event
// @Description get event by ID
// @ID int
// @Accept  json
// @Produce  json
// @Param id path int true "Event ID"
// @Success 200
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /api/v1/events/{id} [get]
func GetEventHandler(eventGetter fanatick.EventGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		event, err := eventGetter.GetEvent(id)
		if err != nil {
			NewErrorWriter(w).Write(err)
			return
		}

		NewJSONWriter(w).Write(event, http.StatusOK)
	}
}

// GetEvents godoc
// @Summary Show a events
// @Description get events
// @Accept  json
// @Produce  json
// @Success 200
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /api/v1/events [get]
func QueryEventsHandler(eventQueryer fanatick.EventQueryer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := map[fanatick.EventQueryParam]interface{}{}

		if limit := r.URL.Query().Get("limit"); limit != "" {
			params[fanatick.EventQueryParamLimit] = limit
		}

		if before := r.URL.Query().Get("before"); before != "" {
			params[fanatick.EventQueryParamBefore] = before
		}

		events, err := eventQueryer.QueryEvent(params)
		if err != nil {
			NewErrorWriter(w).Write(err)
			return
		}

		NewJSONWriter(w).Write(events, http.StatusOK)
	}
}

func PostEventHandler(eventCreator fanatick.EventCreator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		event := fanatick.Event{}
		err := json.NewDecoder(r.Body).Decode(&event)
		if err != nil {
			NewErrorWriter(w).Write(err)
			return
		}
		err = eventCreator.CreateEvent(&event)
		if err != nil {
			NewErrorWriter(w).Write(err)
			return
		}

		NewJSONWriter(w).Write(event, http.StatusOK)
	}
}

func UpdateEventHandler(eventUpdater fanatick.EventUpdater) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		event := fanatick.Event{}
		err := json.NewDecoder(r.Body).Decode(&event)
		if err != nil {
			NewErrorWriter(w).Write(err)
			return
		}
		err = eventUpdater.UpdateEvent(&event)
		if err != nil {
			NewErrorWriter(w).Write(err)
			return
		}

		NewJSONWriter(w).Write(event, http.StatusOK)
	}
}

func DeleteEventHandler(eventDeleter fanatick.EventDeleter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		event := fanatick.Event{}
		err := json.NewDecoder(r.Body).Decode(&event)
		if err != nil {
			NewErrorWriter(w).Write(err)
			return
		}
		err = eventDeleter.DeleteEvent(event.ID)
		if err != nil {
			NewErrorWriter(w).Write(err)
			return
		}

		NewJSONWriter(w).Write(event, http.StatusOK)
	}
}
