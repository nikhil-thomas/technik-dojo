package events

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nikhil-thomas/technik-dojo/cloud-native-programming-with-golang/chp02/myevents/internal/store"
	"github.com/urfave/negroni"
)

// ServeAPI sets routes and starts the server
func ServeAPI(endpoint string, str store.EventReadWriter) error {
	evtAPI := NewEventAPI(str)
	r := mux.NewRouter()

	n := negroni.Classic()
	n.UseHandler(r)

	eventsRouter := r.PathPrefix("/events").Subrouter()
	eventsRouter.Methods(http.MethodGet).Path("/{searchCriteria}/{search}").HandlerFunc(evtAPI.FindEvent)
	eventsRouter.Methods(http.MethodGet).Path("").HandlerFunc(evtAPI.AllEvents)
	eventsRouter.Methods(http.MethodPost).Path("").HandlerFunc(evtAPI.NewEvent)

	log.Println("listening on:", endpoint)
	return http.ListenAndServe(endpoint, n)
}
