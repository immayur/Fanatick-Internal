package main

import (
	"encoding/json"
	"fmt"
	"github.com/swensonhe/fanatick-backend/fanatick"
	"net/http"
)


func PostTicketHandler(ticketAdder fanatick.TicketAdder) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("PostTickethandler")
		ticket := fanatick.Ticket{}
		fmt.Println(r.Body)
		err := json.NewDecoder(r.Body).Decode(&ticket)
		if err!=nil{
			fmt.Println("ticket error")
			NewErrorWriter(w).Write(err)
			return
		}
		fmt.Println(ticket)
		err = ticketAdder.AddTicket(&ticket)
		if err != nil {
			NewErrorWriter(w).Write(err)
			return
		}
		NewJSONWriter(w).Write(ticket, http.StatusOK)
	}
}
