package main

import (
	"net/http"

	"github.com/go-chi/chi"

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
// @Router /authentication/{token} [post]
func LoginEventHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := chi.URLParam(r, "token")
		//event, err := eventGetter.Get(id)
		event, err := api.AuthenticateTokenRequest(token)
		if err != nil {
			NewErrorWriter(w).Write(err)
			return
		}

		NewJSONWriter(w).Write(event, http.StatusOK)
	}
}
