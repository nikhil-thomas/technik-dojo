package events

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nikhil-thomas/technik-dojo/cloud-native-programming-with-golang/chp02/myevents/internal/store"
	"github.com/urfave/negroni"
)

// ServeAPI sets routes and starts the server
func ServeAPI(endpoint, tlsendpoint string, str store.EventReadWriter) (<-chan error, <-chan error) {
	evtAPI := NewEventAPI(str)
	r := mux.NewRouter()

	n := negroni.Classic()
	n.UseHandler(r)

	eventsRouter := r.PathPrefix("/events").Subrouter()
	eventsRouter.Methods(http.MethodGet).Path("/{searchCriteria}/{search}").HandlerFunc(evtAPI.FindEvent)
	eventsRouter.Methods(http.MethodGet).Path("").HandlerFunc(evtAPI.AllEvents)
	eventsRouter.Methods(http.MethodPost).Path("").HandlerFunc(evtAPI.NewEvent)

	httpErrChan := make(chan error)
	httpTLSErrChan := make(chan error)

	go func() {
		defer close(httpErrChan)
		log.Println("http listening on:", endpoint)
		httpErrChan <- http.ListenAndServe(endpoint, n)
	}()

	go func() {
		defer close(httpTLSErrChan)
		log.Println("https listening on:", tlsendpoint)
		httpErrChan <- http.ListenAndServeTLS(tlsendpoint, "./keys/cert.pem", "./keys/key.pem", r)
	}()

	return httpErrChan, httpTLSErrChan
}
