package main

import (
	"net/http"

	"github.com/go-chi/chi"

	"github.com/swensonhe/fanatick-backend/fanatick"
	"github.com/swensonhe/fanatick-backend/fanatick/api"
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
// @Router /events/{id} [get]
func GetEventHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		//event, err := eventGetter.Get(id)
		event, err := api.GetEventHandlerRequest(id)
		if err != nil {
			NewErrorWriter(w).Write(err)
			return
		}

		NewJSONWriter(w).Write(event, http.StatusOK)
	}
}

//stub function for getting all events

func GetAllEvents() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		events, err := api.GetAllEventsRequest()
		if err != nil {
			NewErrorWriter(w).Write(err)
			return
		}

		NewJSONWriter(w).Write(events, http.StatusOK)

	}
}

//stub function for seat listing
func GetSeatListing() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		events, err := api.GetSeatListingRequest()
		if err != nil {
			NewErrorWriter(w).Write(err)
			return
		}

		NewJSONWriter(w).Write(events, http.StatusOK)
	}
}

//stub function to get a seat details
func GetSeatDetails() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		events, err := api.GetSeatDetailsRequest()
		if err != nil {
			NewErrorWriter(w).Write(err)
			return
		}

		NewJSONWriter(w).Write(events, http.StatusOK)
	}
}

//stub function to make an offer
func MakeOffer() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		//this request contains parameters
		event, err := api.MakeOfferRequest()
		if err != nil {
			NewErrorWriter(w).Write(err)
			return
		}
		NewJSONWriter(w).Write(event, http.StatusCreated)
	}
}

//stub for update offer
func UpdateOffer() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		NewJSONWriter(w).Write(nil, http.StatusOK)
	}
}

//stub to delete an offer
func DeleteOffer() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(204)
	}
}

//stub for adding event
func AddEvent() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		event, err := api.AddEventRequest()
		if err != nil {
			NewErrorWriter(w).Write(err)
			return
		}
		NewJSONWriter(w).Write(event, http.StatusOK)
	}
}

//stub for updating event
func UpdateEvent() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		event, err := api.UpdateEventRequest()
		if err != nil {
			NewErrorWriter(w).Write(err)
			return
		}
		NewJSONWriter(w).Write(event, http.StatusOK)
	}
}

//stub for deleting an event
func DeleteEvent() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		event, err := api.DeleteEventRequest()
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
// @Router /events [get]
//parameter forbelow function (eventQueryer fanatick.EventQueryer)
func QueryEventsHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := map[fanatick.EventQueryParam]interface{}{}

		if limit := r.URL.Query().Get("limit"); limit != "" {
			params[fanatick.EventQueryParamLimit] = limit
		}

		if before := r.URL.Query().Get("before"); before != "" {
			params[fanatick.EventQueryParamBefore] = before
		}

		//events, err := eventQueryer.Query(params)
		events, err := api.GetAllEventsRequest()

		if err != nil {
			NewErrorWriter(w).Write(err)
			return
		}

		NewJSONWriter(w).Write(events, http.StatusOK)
	}
}
