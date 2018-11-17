package main

import (
	"log"
	"net/http"
)

// InMemoryPlayerStore implements Player store
type InMemoryPlayerStore struct{}

// GetPlayerScore returns score of a player
func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return 123
}

// RecordWin returns score of a player
func (i *InMemoryPlayerStore) RecordWin(name string) {

}

func main() {
	server := &PlayerServer{&InMemoryPlayerStore{}}

	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalln(err)
	}
}
