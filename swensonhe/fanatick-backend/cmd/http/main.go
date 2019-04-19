package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	_ "github.com/lib/pq"
	"github.com/subosito/gotenv"
	httpSwagger "github.com/swaggo/http-swagger"
	"github.com/swensonhe/fanatick-backend/fanatick/postgres"
)

const BaseV1Url = "/api/v1"

func main() {

	// set swagger doc
	//setSwaggerHeaderDoc()
	gotenv.Load(".envrc-template", "credentials")
	// PostgresQL (DB)
	db, err := postgres.Open(DBConnectionString())
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()

	r.Get("/swagger/*", httpSwagger.WrapHandler)

	/*eventAPI := &api.EventService{
		EventStore: &postgres.EventDB{DB: db},
		Logger:     logrus.StandardLogger(),
	}*/

	//r.Get(BaseV1Url+"/events", QueryEventsHandler(eventAPI))
	//route for all eventlisting
	r.Get(BaseV1Url+"/events", QueryEventsHandler())

	//route for stub getEVent
	r.Get(BaseV1Url+"/getEvent/{id}", GetEventHandler())

	//route for user sign in
	r.Post(BaseV1Url+"/authentications/{token}", LoginEventHandler())

	//route for seat listing
	r.Get(BaseV1Url+"/events/{id}/seats", GetSeatListing())

	//route for getting seat details
	r.Get(BaseV1Url+"/seats/{id}", GetSeatDetails())

	//route for makeoffer statuscode should be 201ok
	r.Post(BaseV1Url+"/offers/{offerAmount, seatId, eventId}", MakeOffer())

	//route for update offer
	r.Put(BaseV1Url+"/offers/{offerAmount,seatId,eventId,offerId}", UpdateOffer())

	//route to cancel offer
	r.Delete(BaseV1Url+"/offers/{id}", DeleteOffer())

	//route to add an event
	r.Post(BaseV1Url+"/addEvent/{eventName, location, image, dateAndtime, ticketPrice, quantity,deliveryMethod,quickSell, notLeaveWithSingleTicket}", AddEvent())

	//route to update event
	r.Post(BaseV1Url+"/updateEvent/{eventId, eventName, location, image, dateAndtime, ticketPrice, quantity, deliveryMethod, quickSell, notLeaveWothSingleTicket}", UpdateEvent())
	
	//route to delete an event
	r.Post(BaseV1Url+"/deleteEvent/{eventId}", DeleteEvent())

	if err = ListenAndServe(MustGetenv("PORT"), r); err != nil {
		panic(err)
	}
}

// ListenAndServe serves the application.
func ListenAndServe(port string, handler http.Handler) error {
	fmt.Println("Listening on:", port)
	return http.ListenAndServe(fmt.Sprintf(":%s", port), handler)
}

// DBConnectionString returns the database connection string.
func DBConnectionString() string {
	connectionString := fmt.Sprintf(`host=%s`, MustGetenv("DB_HOST"))

	if port := os.Getenv("DB_PORT"); port != "" {
		connectionString += fmt.Sprintf(` port=%s`, port)
	}

	if user := os.Getenv("DB_USER"); user != "" {
		connectionString += fmt.Sprintf(` user=%s`, user)
	}

	if password := os.Getenv("DB_PASSWORD"); password != "" {
		connectionString += fmt.Sprintf(` password=%s`, password)
	}

	if name := os.Getenv("DB_NAME"); name != "" {
		connectionString += fmt.Sprintf(` dbname=%s`, name)
	}

	if mode := os.Getenv("DB_SSL_MODE"); mode != "" {
		connectionString += fmt.Sprintf(` sslmode=%s`, mode)
	}

	return connectionString
}

// MustGetenv gets an environment variable or panics.
func MustGetenv(key string) string {
	v := os.Getenv(key)
	if v == "" {
		panic(fmt.Sprintf("%s missing", key))
	}
	return v
}

/*func setSwaggerHeaderDoc() {
	docs.SwaggerInfo.Title = "Swagger Fanatic API"
	docs.SwaggerInfo.Description = "This is a fanatic server ."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:" + MustGetenv("PORT")
	docs.SwaggerInfo.BasePath = BaseV1Url
}*/
