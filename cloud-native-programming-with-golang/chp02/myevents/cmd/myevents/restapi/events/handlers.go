package events

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/nikhil-thomas/technik-dojo/cloud-native-programming-with-golang/chp02/myevents/internal/store"
)

type eventAPI struct {
	eventStore store.EventReadWriter
}

// NewEventAPI creates a new eventAPI
func NewEventAPI(evtStr store.EventReadWriter) *eventAPI {
	return &eventAPI{
		eventStore: evtStr,
	}
}

func (ea *eventAPI) FindEvent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	criteria, ok := vars["searchCriteria"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, `{"error": "no search criteria found"}`)
		return
	}

	searchKey, ok := vars["search"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, `{"error": "no search key found"}`)
		return
	}

	var event store.Event
	var err error

	switch strings.ToLower(criteria) {
	case "name":
		event, err = ea.eventStore.FindEventByName(searchKey)
	case "id":
		id, err := hex.DecodeString(searchKey)
		if err != nil {
			break
		}
		event, err = ea.eventStore.FindEvent(id)
	}
	if err != nil {
		fmt.Fprintf(w, `{"error": "%s"}`, err)
		return
	}
	w.Header().Set("Content-Type", "application/json;charset=utf8")
	err = json.NewEncoder(w).Encode(&event)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"error": "%s"}`, err)
		return
	}
}

func (ea *eventAPI) AllEvents(w http.ResponseWriter, r *http.Request) {
	events, err := ea.eventStore.FindAllEvents()
	if err != nil {
		fmt.Fprintf(w, `{"error": "%s"}`, err)
		return
	}
	w.Header().Set("Content-Type", "application/json;charset=utf8")
	err = json.NewEncoder(w).Encode(&events)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"error": "%s"}`, err)
		return
	}
}

func (ea *eventAPI) NewEvent(w http.ResponseWriter, r *http.Request) {
	event := store.Event{}
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"error": "%s"}`, err)
		return
	}

	id, err := ea.eventStore.AddEvent(event)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"error": "%s"}`, err)
		return
	}
	fmt.Fprintf(w, `{"id": "%s"}`, string(id))
}
