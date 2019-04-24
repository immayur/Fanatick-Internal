package main

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"net/http"

	"github.com/swensonhe/fanatick-backend/fanatick"
)

// GetSeat godoc
// @Summary Show a seat
// @Description get seat by ID
// @ID int
// @Accept  json
// @Produce  json
// @Param id path int true "Seat ID"
// @Success {object} 200
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /seats/{id} [get]
func GetSeatHandler(seatGetter fanatick.SeatGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		seat, err := seatGetter.GetSeatByID(id)
		if err != nil {
			NewErrorWriter(w).Write(err)
			return
		}

		NewJSONWriter(w).Write(seat, http.StatusOK)
	}
}

// GetSeat godoc
// @Summary Show seats by event ID
// @Description get seats by event ID
// @ID int
// @Accept  json
// @Produce  json
// @Param id path int true "Event ID"
// @Success 200
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /events/{id}/seats [get]
func QuerySeatsHandler(seatQueryer fanatick.SeatQueryer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		params := map[fanatick.SeatQueryParam]interface{}{}

		if limit := r.URL.Query().Get("limit"); limit != "" {
			params[fanatick.SeatQueryParamLimit] = limit
		}

		if before := r.URL.Query().Get("before"); before != "" {
			params[fanatick.SeatQueryParamBefore] = before
		}

		events, err := seatQueryer.QuerySeatsbyEventID(id, params)
		if err != nil {
			NewErrorWriter(w).Write(err)
			return
		}

		NewJSONWriter(w).Write(events, http.StatusOK)
	}
}

func PostSeatHandler(seatCreator fanatick.SeatCreator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		seat := fanatick.Seat{}
		err := json.NewDecoder(r.Body).Decode(&seat)
		if err != nil {
			NewErrorWriter(w).Write(err)
			return
		}
		err = seatCreator.CreateSeat(id, &seat)
		if err != nil {
			NewErrorWriter(w).Write(err)
			return
		}

		NewJSONWriter(w).Write(seat, http.StatusOK)
	}
}

/*func PostTicketHandler(ticketAdder fanatick.TicketAdder) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//section := chi.URLParam(r, key: "section")
		//row := chi.URLParam(r, key: "row")
		//seat := chi.URLParam(r, key: "seat")
		fmt.Println("PostTickethandler")
		ticket := fanatick.Ticket{}
		err := json.NewDecoder(r.Body).Decode(&ticket)
		if err!=nil{
			fmt.Println("ticket error")
			NewErrorWriter(w).Write(err)
			return
		}
		fmt.Println(ticket)
		err = ticketAdder.AddTicket("12", &ticket)
		if err != nil {
			NewErrorWriter(w).Write(err)
			return
		}
		NewJSONWriter(w).Write(ticket, http.StatusOK)
	}
}
*/