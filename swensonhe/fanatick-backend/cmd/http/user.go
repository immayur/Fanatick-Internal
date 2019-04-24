package main

import (
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/swensonhe/fanatick-backend/fanatick/jwt"

	"github.com/swensonhe/fanatick-backend/fanatick"
)

// GetUser godoc
// @Summary Show a user
// @Description get user by ID
// @ID int
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Success 200
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /api/v1/users/{firebase_uuid} [get]
func GetUserByFireBaseUUIDHandler(userGetter fanatick.UserGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		firebaseUUID := chi.URLParam(r, "firebase_uuid")
		user, err := userGetter.GetUserByFireBaseUUID(firebaseUUID)
		if err != nil {
			NewErrorWriter(w).Write(err)
			return
		}

		NewJSONWriter(w).Write(user, http.StatusOK)
	}
}

// PostUserHandler godoc
// @Summary create a user
// @Description create a user
// @ID int
// @Accept  json
// @Produce  json
// @Success 200
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /api/v1/users/ [post]
func PostUserHandler(userCreator fanatick.UserCreator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := fanatick.User{}
		err := decodeJSON(r, &user)
		if err != nil {
			NewErrorWriter(w).Write(err)
			return
		}
		// return user if exist --> login
		err = userCreator.CreateUser(&user)
		if err != nil {
			NewErrorWriter(w).Write(err)
			return
		}

		// generate token for login user
		ta := &jwt.TokenAuth{
			Secret:        os.Getenv("CLIENT_SECRET"),
			TokenDuration: 10,
		}

		token, err := ta.Generate(&user)
		if err != nil {
			NewErrorWriter(w).Write(err)
			return
		}

		userRes := &fanatick.UserCreateResponse{User: user, Token: token}
		NewJSONWriter(w).Write(userRes, http.StatusOK)
	}
}
