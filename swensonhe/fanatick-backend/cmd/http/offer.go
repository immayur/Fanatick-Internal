package main

import (
	"github.com/swensonhe/fanatick-backend/fanatick"
	"net/http"
)

func PostOfferHandler(offerCreator fanatick.OfferCreator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}