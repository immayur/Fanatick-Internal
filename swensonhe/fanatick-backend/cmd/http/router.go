package main

import (
	"os"

	"github.com/sirupsen/logrus"
	httpSwagger "github.com/swaggo/http-swagger"
	"github.com/swensonhe/fanatick-backend/fanatick/api"
	fanaticFirebase "github.com/swensonhe/fanatick-backend/fanatick/firebase"
	"github.com/swensonhe/fanatick-backend/fanatick/postgres"

	"net/http"

	"github.com/go-chi/chi"
)

// NewRouter creates a new *chi.Mux
func NewRouter(db *postgres.DB, fsClient *fanaticFirebase.Client) *chi.Mux {
	router := chi.NewRouter()

	// public routes
	router.Get("/swagger/*", httpSwagger.WrapHandler)

	routes := GetRoutes(db, fsClient)
	for _, route := range routes {
		var handler http.Handler

		// Last thing that will be executed is the actual handler function
		handler = route.Handler

		if os.Getenv("SkipAllAuth") != "true" {
			if !route.skipAuth {
				handler = AuthMiddleware(route.Handler)
			}
		}

		router.Method(route.Method, route.Pattern, handler)
	}
	return router
}

type Route struct {
	Method   string
	Pattern  string
	Handler  http.HandlerFunc
	skipAuth bool
}

// MyServer is a wrapper for chi.Router pointer
type MyServer struct {
	r *chi.Mux
}

// ServeHTTP servers HTTP
func (s *MyServer) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	if origin := req.Header.Get("Origin"); origin != "" {
		rw.Header().Set("Access-Control-Allow-Origin", origin)
		rw.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, HEAD")
		rw.Header().Set("Access-Control-Allow-Credentials", "true")
		rw.Header().Set("Access-Control-Max-Age", "86400")
		// rw.Header().Set("Access-Control-Expose-Headers", "Content-Disposition")
		rw.Header().Set("Access-Control-Allow-Headers",
			"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-XSRF-Token, X-HTTP-Method-Override, X-Requested-With, Mobile-Cookie")
	}

	if req.Method == "OPTIONS" {
		return
	}

	// Let chi work
	s.r.ServeHTTP(rw, req)
}

type Routes []Route

func GetRoutes(db *postgres.DB, fsClient *fanaticFirebase.Client) Routes {

	var allRoutes Routes
	// event routes
	eventAPI := &api.EventService{
		EventStore: &postgres.EventDB{DB: db},
		Logger:     logrus.StandardLogger(),
	}

	eventRoutes := Routes{
		Route{
			Method:   "GET",
			Pattern:  BaseV1Url + "/events",
			Handler:  QueryEventsHandler(eventAPI),
			skipAuth: false,
		},

		Route{
			Method:   "GET",
			Pattern:  BaseV1Url + "/events/{id}",
			Handler:  GetEventHandler(eventAPI),
			skipAuth: false,
		},

		Route{
			Method:   "POST",
			Pattern:  BaseV1Url + "/events",
			Handler:  PostEventHandler(eventAPI),
			skipAuth: false,
		},

		Route{
			Method:   "DELETE",
			Pattern:  BaseV1Url + "/events/{id}",
			Handler:  DeleteEventHandler(eventAPI),
			skipAuth: false,
		},

		Route{
			Method:   "PUT",
			Pattern:  BaseV1Url + "/events/{id}",
			Handler:  UpdateEventHandler(eventAPI),
			skipAuth: false,
		},
	}

	userAPI := &api.UserService{
		UserStore:      &postgres.UserDB{DB: db},
		Logger:         logrus.StandardLogger(),
		FirebaseClient: fsClient,
	}

	userRoutes := Routes{
		Route{
			Method:   "GET",
			Pattern:  BaseV1Url + "/users/{firebase_uuid}",
			Handler:  GetUserByFireBaseUUIDHandler(userAPI),
			skipAuth: false,
		},

		Route{
			Method:   "POST",
			Pattern:  BaseV1Url + "/authentications",
			skipAuth: true,
			Handler:  PostUserHandler(userAPI),
		},
	}

	userProfileAPI := &api.UserProfileService{
		UserProfileStore: &postgres.UserProfileDB{DB: db},
		Logger:           logrus.StandardLogger(),
	}

	userProfileRoutes := Routes{
		Route{
			Method:   "GET",
			Pattern:  BaseV1Url + "/users/{user_id}/profile",
			Handler:  GetUserProfileByUserIdHandler(userProfileAPI),
			skipAuth: false,
		},

		Route{
			Method:   "POST",
			Pattern:  BaseV1Url + "/users/{user_id}/profile",
			Handler:  PostUserProfileHandler(userProfileAPI),
			skipAuth: false,
		},

		Route{
			Method:   "PATCH",
			Pattern:  BaseV1Url + "/users/{user_id}/profile",
			Handler:  PatchUserProfileHandler(userProfileAPI),
			skipAuth: false,
		},
	}

	seatAPI := &api.SeatService{
		SeatStore: &postgres.SeatDB{DB: db},
		Logger:    logrus.StandardLogger(),
	}

	seatRoutes := Routes{
		Route{
			Method:   "GET",
			Pattern:  BaseV1Url + "/events/{id}/seats",
			Handler:  QuerySeatsHandler(seatAPI),
			skipAuth: true,
		},

		Route{
			Method:   "GET",
			Pattern:  BaseV1Url + "/seats/{id}",
			Handler:  GetSeatHandler(seatAPI),
			skipAuth: true,
		},

		Route{
			Method:   "POST",
			Pattern:  BaseV1Url + "/events/{id}/seats",
			Handler:  PostSeatHandler(seatAPI),
			skipAuth: true,
		},
	}

	ticketAPI := &api.TicketService{
		TicketStore: &postgres.TicketDB{DB: db},
		Logger:    logrus.StandardLogger(),
	}

	ticketRoutes := Routes{
		Route{
			Method: "GET",
			Pattern: BaseV1Url + "/addTicket",
			Handler: PostTicketHandler(ticketAPI),
			skipAuth: true,
		},
	}

	allRoutes = append(allRoutes, eventRoutes...)
	allRoutes = append(allRoutes, userRoutes...)
	allRoutes = append(allRoutes, userProfileRoutes...)
	allRoutes = append(allRoutes, seatRoutes...)
	allRoutes = append(allRoutes, ticketRoutes...)
	return allRoutes
}
