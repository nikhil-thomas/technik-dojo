package main

// InMemoryPlayerStore implements Player store
type InMemoryPlayerStore struct {
	store map[string]int
}

// GetPlayerScore returns score of a player
func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return i.store[name]
}

// RecordWin returns score of a player
func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.store[name]++
}

// NewInMemoryPlaystore initializes an empty playerstore
func NewInMemoryPlaystore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{map[string]int{}}
}
